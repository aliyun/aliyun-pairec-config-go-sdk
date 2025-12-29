package test

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"testing"
)

func TestListTrafficControlTasks(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	// 预发运行中的调控任务
	//preTasks := client.ListTrafficControlTasks(common.Environment_Prepub_Desc)
	//
	//for _, task := range preTasks {
	//	d, _ := json.Marshal(task)
	//	fmt.Println(fmt.Sprintf("%s: %s\n", task.TrafficControlTaskId, string(d)))
	//}
	// 生产运行中的调控任务
	tasks := client.ListTrafficControlTasks(common.Environment_Product_Desc)
	for _, task := range tasks {
		d, _ := json.Marshal(task)
		fmt.Println(fmt.Sprintf("%s: %s\n", task.TrafficControlTaskId, string(d)))
	}
}

func TestGetTrafficControlActualTraffic(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)

	// 比例调控的流量
	percentControlTraffics := client.GetTrafficControlActualTraffic(common.Environment_Product_Desc, "ER_ALL")
	d1, _ := json.Marshal(percentControlTraffics)
	fmt.Println(string(d1))
	// 单品调控的流量
	itemIdList := []string{"7809641366", "7802947776"}
	singleControlTraffics := client.GetTrafficControlActualTraffic(common.Environment_Product_Desc, itemIdList...)
	d2, _ := json.Marshal(singleControlTraffics)
	fmt.Println(string(d2))

}
