package api

import (
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
)

type ListFlowCtrlPlansResponse struct {
	BaseResponse
	Data struct {
		TrafficControlTasks []model.TrafficControlTasksItem `json:"traffic_control_tasks"`
	} `json:"data,omitempty"`
}
