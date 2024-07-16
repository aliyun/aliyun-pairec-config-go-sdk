package api

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pairecservice"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
	jsoniter "github.com/json-iterator/go"
)

type TrafficControlTrafficsService service

func (tct *TrafficControlTrafficsService) SetTrafficControlTrafficFData(t model.TrafficControlTaskTrafficData) (string, error) {
	t.InstanceId = tct.instanceId
	req := pairecservice.CreateUpdateTrafficControlTaskTrafficRequest()
	body, _ := jsoniter.MarshalToString(t)

	req.TrafficControlTaskId = t.TrafficControlTaskId
	req.Body = body
	req.SetDomain(tct.client.GetDomain())

	response, err := tct.client.UpdateTrafficControlTaskTraffic(req)

	if err != nil {
		return "", err
	}
	return response.RequestId, nil

}
