package recallengine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Write(instanceId, table string, request *WriteRequest) (*WriteResponse, error) {
	if c.RetryTimes > 0 {
		var err error
		for i := 0; i < c.RetryTimes; i++ {
			var writeResponse *WriteResponse
			writeResponse, err = c.doWrite(instanceId, table, request)
			if err == nil {
				return writeResponse, nil
			} else if c.Logger != nil {
				c.Logger.Printf("recallengine: write failed, retrying..., err: %s", err.Error())
			}
		}
		return nil, err
	} else {
		return c.doWrite(instanceId, table, request)
	}
}

func (c *Client) doWrite(instanceId, table string, request *WriteRequest) (*WriteResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("invalid request, err: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/tables/%s/default/%s/write", c.Endpoint, instanceId, table)
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
				return nil, fmt.Errorf("write request failed, response status code: %d, message: %s", resp.StatusCode, msg)
			}
		}
		return nil, fmt.Errorf("write request failed, response status code: %d", resp.StatusCode)
	}

	response := WriteResponse{}
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body, err: %w", err)
	}

	return &response, nil
}
