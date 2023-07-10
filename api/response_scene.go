package api

import "github.com/aliyun/aliyun-pairec-config-go-sdk/model"

type ListScenesResponse struct {
	BaseResponse
	Data map[string][]*model.Scene `json:"data,omitempty"`
}
