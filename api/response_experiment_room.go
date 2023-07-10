package api

import "github.com/aliyun/aliyun-pairec-config-go-sdk/model"

type ListExperimentRoomsResponse struct {
	BaseResponse
	Data map[string][]*model.ExperimentRoom `json:"data,omitempty"`
}
