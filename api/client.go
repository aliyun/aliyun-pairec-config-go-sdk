package api

import (
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	pairecservice20221213 "github.com/alibabacloud-go/pairecservice-20221213/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"net"
	"net/http"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pairecservice"
	credentialsv2 "github.com/aliyun/credentials-go/credentials"
)

var (
	defaultTransport = &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			d := net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}
			return d.DialContext(ctx, "tcp4", addr)
		},
		MaxIdleConns:          200,
		MaxIdleConnsPerHost:   200,
		MaxConnsPerHost:       200,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
)

// APIClient manages communication with the Pairec Experiment Restful Api API v1.0.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	*pairecservice.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	region string

	domain string

	// API Services
	ExperimentApi *ExperimentApiService

	ExperimentGroupApi *ExperimentGroupApiService

	ExperimentRoomApi *ExperimentRoomApiService

	LayerApi *LayerApiService

	SceneApi *SceneApiService

	ParamApi *ParamApiService

	CrowdApi *CrowdApiService

	//TrafficControlApi *TrafficControlApiService

	FeatureConsistencyCheckApi *FeatureConsistencyCheckService

	clientV2 *pairecservice20221213.Client
}

type service struct {
	client     *APIClient
	instanceId string
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(instanceId, region, accessId, accessKey string) (*APIClient, error) {
	var (
		client   *pairecservice.Client
		err      error
		clientV2 *pairecservice20221213.Client
	)
	config := &openapi.Config{}
	config.Endpoint = tea.String(fmt.Sprintf("pairecservice-vpc.%s.aliyuncs.com", region))

	if accessId == "" || accessKey == "" {
		defaultProvider := credentials.NewDefaultCredentialsProvider()
		sdkConfig := sdk.NewConfig()
		sdkConfig.Scheme = "https"
		client, err = pairecservice.NewClientWithOptions(region, sdkConfig, defaultProvider)
		if err != nil {
			return nil, err
		}
		credential, err1 := credentialsv2.NewCredential(nil)
		if err1 != nil {
			return nil, err1
		}
		config.Credential = credential
		clientV2, err = pairecservice20221213.NewClient(config)
		if err != nil {
			return nil, err
		}
	} else {
		client, err = pairecservice.NewClientWithAccessKey(region, accessId, accessKey)
		if err != nil {
			return nil, err
		}

		config.AccessKeyId = tea.String(accessId)
		config.AccessKeySecret = tea.String(accessKey)
		clientV2, err = pairecservice20221213.NewClient(config)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}
	client.SetTransport(defaultTransport)

	c := &APIClient{
		Client:   client,
		region:   region,
		clientV2: clientV2,
	}
	c.common.client = c
	c.common.instanceId = instanceId

	// API Services
	c.ExperimentApi = (*ExperimentApiService)(&c.common)
	c.ExperimentGroupApi = (*ExperimentGroupApiService)(&c.common)
	c.ExperimentRoomApi = (*ExperimentRoomApiService)(&c.common)
	c.LayerApi = (*LayerApiService)(&c.common)
	c.SceneApi = (*SceneApiService)(&c.common)
	c.ParamApi = (*ParamApiService)(&c.common)
	c.CrowdApi = (*CrowdApiService)(&c.common)
	//c.TrafficControlApi = (*TrafficControlApiService)(&c.common)
	c.FeatureConsistencyCheckApi = (*FeatureConsistencyCheckService)(&c.common)
	return c, nil
}

func (c *APIClient) GetDomain() string {
	if c.domain == "" {
		c.domain = fmt.Sprintf("pairecservice-vpc.%s.aliyuncs.com", c.region)
	}

	return c.domain
}

func (c *APIClient) SetDomain(domain string) {
	c.domain = domain
}

/**
func (c *APIClient) Init(accessId, accessKey string) error {
	endpoint := c.GetDomain()
	protol := "http"
	config := &openapi.Config{
		AccessKeyId:     &accessId,
		AccessKeySecret: &accessKey,
		Endpoint:        &endpoint,
		Protocol:        &protol,
	}

	client, err := pairecserviceV2.NewClient(config)

	if err != nil {
		return err
	}

	c.v2Client = client
	return nil
}

**/
