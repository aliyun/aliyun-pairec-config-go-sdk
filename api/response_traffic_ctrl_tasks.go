package api

import (
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
)

type ListTrafficControlTasksResponse struct {
	TrafficControlTasks []*model.TrafficControlTask `json:"traffic_control_tasks"`
}
