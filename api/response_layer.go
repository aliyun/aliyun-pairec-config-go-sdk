package api

import "github.com/aliyun/aliyun-pairec-config-go-sdk/model"

type ListLayersResponse struct {
	BaseResponse
	Data map[string][]*model.Layer `json:"data,omitempty"`
}
