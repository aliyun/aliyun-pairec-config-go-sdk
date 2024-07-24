package api

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
)

type ListTrafficControlTasksResponse struct {
	*responses.BaseResponse
	RequestId           string                     `json:"requestId"`
	TotalCount          string                     `json:"totalCount"`
	TrafficControlTasks []model.TrafficControlTask `json:"traffic_control_tasks"`
}
