package api

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pairecservice"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
)

type TrafficControlTrafficsService service

func (tct *TrafficControlTrafficsService) SetTrafficControlTrafficFData(t model.TrafficControlTaskTrafficData) string {
	t.InstanceId = tct.instanceId
	req := pairecservice.CreateUpdateTrafficControlTaskTrafficRequest()
	body := fmt.Sprintf("%v", t)
	req.TrafficControlTaskId = t.TrafficControlTaskId
	req.Body = body
	req.SetDomain(tct.client.GetDomain())

	response, err := tct.client.UpdateTrafficControlTaskTraffic(req)

	if err != nil {
		return ""
	}

	return response.RequestId

}
