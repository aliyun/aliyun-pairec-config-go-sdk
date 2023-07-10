package api

import "github.com/aliyun/aliyun-pairec-config-go-sdk/model"

type ListExperimentGroupsResponse struct {
	BaseResponse
	Data map[string][]*model.ExperimentGroup `json:"data,omitempty"`
	// Scene *model.Scene
}
