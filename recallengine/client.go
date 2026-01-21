package recallengine

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	pairecservice20221213 "github.com/alibabacloud-go/pairecservice-20221213/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

var (
	defaultRequestTimeout = 500 * time.Millisecond
	defaultTransport      = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   200 * time.Millisecond,
			KeepAlive: 5 * time.Minute,
		}).DialContext,
		MaxIdleConns:          1000,
		MaxIdleConnsPerHost:   1000,
		MaxConnsPerHost:       1000,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
	}
	defaultHttpClient = &http.Client{
		Timeout:   defaultRequestTimeout,
		Transport: defaultTransport,
	}
)

type Client struct {
	Endpoint string
	Username string
	Password string

	RetryTimes int

	RequestHeaders map[string]string

	// Logger specifies a logger used to report internal changes within the writer
	Logger Logger

	// ErrorLogger is the logger to report errors
	ErrorLogger Logger

	httpClient *http.Client

	auth string
}

func NewClient(endpoint, username, password string, opts ...ClientOption) *Client {
	client := Client{
		Endpoint: endpoint,
		Username: username,
		Password: password,

		httpClient: defaultHttpClient,
	}

	for _, opt := range opts {
		opt(&client)
	}

	return &client
}

func (c *Client) Recall(request *RecallRequest) (*RecallResponse, error) {
	if c.RetryTimes > 0 {
		var err error
		for i := 0; i < c.RetryTimes; i++ {
			var recallResponse *RecallResponse
			recallResponse, err = c.doRecall(request)
			if err == nil {
				return recallResponse, nil
			} else if c.Logger != nil {
				c.Logger.Printf("recallengine: recall failed, retrying..., err: %s", err.Error())
			}
		}
		return nil, err
	} else {
		return c.doRecall(request)
	}
}

func (c *Client) doRecall(request *RecallRequest) (*RecallResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("invalid request, err: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/recall", c.Endpoint)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Auth", c.buildAuth())

	for k, v := range c.RequestHeaders {
		req.Header.Set(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed, err: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body, err: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var bodyMap map[string]any
		if err := json.Unmarshal(respBody, &bodyMap); err == nil {
			if msg, ok := bodyMap["message"]; ok {
				return nil, fmt.Errorf("request failed, response status code: %d, message: %s", resp.StatusCode, msg)
			}
		}
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}

	record := UnSerializeRecord(respBody)

	return &RecallResponse{Result: record}, nil
}

func (c *Client) buildAuth() string {
	if c.auth == "" {
		c.auth = base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))
	}
	return c.auth
}

type GetRecallEngineEndpointOption struct {
	PAIRecServiceEndpoint string

	AccessKeyId string

	AccessKeySecret string

	VpcId string
}

func GetRecallEngineEndpoint(instanceId, regionId string, option *GetRecallEngineEndpointOption) (string, error) {
	var pairecServiceEndpoint string
	if option != nil && option.PAIRecServiceEndpoint != "" {
		pairecServiceEndpoint = option.PAIRecServiceEndpoint
	} else {
		pairecServiceEndpoint = fmt.Sprintf("pairecservice-vpc.%s.aliyuncs.com", regionId)
	}

	config := &openapi.Config{
		Endpoint: tea.String(pairecServiceEndpoint),
	}

	if option != nil && option.AccessKeyId != "" && option.AccessKeySecret != "" {
		config.AccessKeyId = tea.String(option.AccessKeyId)
		config.AccessKeySecret = tea.String(option.AccessKeySecret)
	} else {
		credential, err := credentials.NewCredential(nil)
		if err != nil {
			return "", fmt.Errorf("failed to create credential, err: %w", err)
		}
		config.Credential = credential
	}

	pairecClient, err := pairecservice20221213.NewClient(config)
	if err != nil {
		return "", err
	}

	request := &pairecservice20221213.GetRecallManagementConfigRequest{
		InstanceId: tea.String(instanceId),
	}

	response, err := pairecClient.GetRecallManagementConfig(request)
	if response.StatusCode != nil && *response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid response: %s", response.String())
	}

	if response.Body != nil && len(response.Body.NetworkConfigs) > 0 {
		if option.VpcId == "" {
			if len(response.Body.NetworkConfigs) == 1 {
				if response.Body.NetworkConfigs[0].PrivateLinkAddress != nil && response.Body.NetworkConfigs[0].Status != nil {
					if *response.Body.NetworkConfigs[0].Status != "Connected" {
						return "", errors.New("endpoint unavailable")
					}
					return *response.Body.NetworkConfigs[0].PrivateLinkAddress, nil
				}
			} else {
				return "", errors.New("multiple VPCs were found, please specify the VpcId")
			}
		} else {
			for _, networkConfig := range response.Body.NetworkConfigs {
				if networkConfig != nil && networkConfig.VpcId != nil && networkConfig.PrivateLinkAddress != nil && response.Body.NetworkConfigs[0].Status != nil {
					if *response.Body.NetworkConfigs[0].Status != "Connected" {
						continue
					}

					if option.VpcId == *networkConfig.VpcId {
						return *networkConfig.PrivateLinkAddress, nil
					}
				}
			}
		}
	}

	return "", errors.New("no available endpoints found")
}
