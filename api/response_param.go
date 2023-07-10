package api

import "github.com/aliyun/aliyun-pairec-config-go-sdk/model"

type ListParamsResponse struct {
	BaseResponse
	Data map[string][]*model.Param `json:"data,omitempty"`
}
