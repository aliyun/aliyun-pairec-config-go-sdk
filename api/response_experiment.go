package api

import "github.com/aliyun/aliyun-pairec-config-go-sdk/model"

type ListExperimentsResponse struct {
	BaseResponse
	Data map[string][]*model.Experiment `json:"data,omitempty"`
	// Scene *model.Scene
}
