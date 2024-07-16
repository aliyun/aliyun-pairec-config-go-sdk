package loadTest

import (
	"fmt"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
	"math/big"
	"testing"
)

func TestGetTrafficControlTaskMetaData(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	plans := client.GetTrafficControlTaskMetaData("Prod", 0)
	fmt.Println("-----------")
	for _, plan := range plans {
		fmt.Printf("%++v\n", plan)
		for _, target := range plan.TrafficControlTargets {
			fmt.Printf("%++v\n", target)
		}

	}
}

func TestGetTrafficControlTargetData(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	targets := client.GetTrafficControlTargetData("Pre", "", 0)
	for targetId, target := range targets {
		fmt.Printf("%d %+v\n", targetId, target)
	}
}

func TestCheckIfTrafficControlTargetIsEnabled(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	enabled := client.CheckIfTrafficControlTargetIsEnabled("Pre", 45, 0)
	fmt.Println(enabled)
}

func TestGetTrafficControlTargetTraffic(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	var trafficsArray []model.TrafficData
	traffics := model.TrafficData{
		TrafficControlTargetId:         "47",
		RecordTime:                     "2024-07-21 T13:05:06.111Z",
		ItemOrExperimentId:             "",
		TrafficControlTargetTraffic:    *big.NewInt(20000),
		TrafficControlTargetAimTraffic: 100.0,
		TrafficControlTaskTraffic:      *big.NewInt(10000),
	}
	trafficsArray = append(trafficsArray, traffics)
	trafficsData := model.TrafficControlTaskTrafficData{
		TrafficControlTaskId: "57",
		Traffics:             trafficsArray,
		Environment:          "Pre",
	}
	requestId, err := client.SetTrafficControlTraffic(trafficsData)

	if err != nil {
		fmt.Printf("err=%v\n", err)
	} else {
		fmt.Printf("requestId=%v\n", requestId)
	}

	fmt.Println(client.GetTrafficControlTargetTraffic("Pre", "HomePage", ""))
	//
	//idList := []string{"ER_ALL", "12345678", "unknown"}
	//fmt.Printf("%+v\n", client.GetTrafficControlTargetTraffic("prepub", "test1", idList...))

}
