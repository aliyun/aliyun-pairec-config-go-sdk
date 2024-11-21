package test

import (
	"fmt"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"testing"
)

func TestGetTrafficControlTaskMetaData(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	plans := client.GetTrafficControlTaskMetaData("prepub", 0)

	for _, plan := range plans {
		fmt.Printf("---【Prepub task】---%s,%s\n", plan.TrafficControlTaskId, plan.Name)
		for _, target := range plan.TrafficControlTargets {
			fmt.Printf("---【Prepub target】----%s，%s\n", target.TrafficControlTargetId, target.Name)
		}
	}

	proPlans := client.GetTrafficControlTaskMetaData("product", 0)

	for _, plan := range proPlans {
		fmt.Printf("---【Product task】---%s,%s\n", plan.TrafficControlTaskId, plan.Name)
		for _, target := range plan.TrafficControlTargets {
			fmt.Printf("---【Product target】----%s，%s\n", target.TrafficControlTargetId, target.Name)
		}
	}
}

func TestGetTrafficControlTargetData(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	preTargets := client.GetTrafficControlTargetData("prepub", "", 0)
	for targetId, target := range preTargets {
		fmt.Printf("%s %+v\n", targetId, target)
	}

	proTargets := client.GetTrafficControlTargetData("product", "", 0)
	for targetId, target := range proTargets {
		fmt.Printf("%s %+v\n", targetId, target)
	}
}

func TestCheckIfTrafficControlTargetIsEnabled(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	enabled := client.CheckIfTrafficControlTargetIsEnabled("product", 1, 0)
	fmt.Println(enabled)
}

func TestGetTrafficControlTargetTraffic(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	//var trafficsArray []model.TrafficData
	//traffics := model.TrafficData{
	//	TrafficControlTargetId:         "47",
	//	RecordTime:                     "2024-07-21 T13:05:06.111Z",
	//	ItemOrExperimentId:             "",
	//	TrafficControlTargetTraffic:    *big.NewInt(20000),
	//	TrafficControlTargetAimTraffic: 100.0,
	//	TrafficControlTaskTraffic:      *big.NewInt(10000),
	//}
	//trafficsArray = append(trafficsArray, traffics)
	//trafficsData := model.TrafficControlTaskTrafficData{
	//	TrafficControlTaskId: "57",
	//	Traffics:             trafficsArray,
	//	Environment:          "Pre",
	//}
	//requestId, err := client.SetTrafficControlTraffic(trafficsData)

	//if err != nil {
	//	fmt.Printf("err=%v\n", err)
	//} else {
	//	fmt.Printf("requestId=%v\n", requestId)
	//}

	fmt.Println(client.GetTrafficControlTargetTraffic("prepub", "home_feed", "ER1_L1_L6_L8_L10#EG16#E26"))
	//idList := []string{"ER_ALL", "12345678", "unknown"}
	//fmt.Printf("%+v\n", client.GetTrafficControlTargetTraffic("prepub", "test1", idList...))

}
