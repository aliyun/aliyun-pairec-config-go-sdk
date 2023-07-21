package experiments

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
)

func createExperimentClient(environment string) *ExperimentClient {
	region := "cn-hangzhou"
	instanceId := os.Getenv("INSTANCE_ID")
	accessId := os.Getenv("ACCESS_ID")
	accessKey := os.Getenv("ACCESS_KEY")
	client, err := NewExperimentClient(instanceId, region, accessId, accessKey, environment, WithLogger(LoggerFunc(log.Printf)), WithDomain("pairecservice.cn-hangzhou.aliyuncs.com"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func TestMatchExperiment2(t *testing.T) {
	client := createExperimentClient(common.Environment_Daily_Desc)

	experimentContext := model.ExperimentContext{
		RequestId:    "pvid",
		Uid:          "102441835",
		FilterParams: map[string]interface{}{},
	}

	experimentResult := client.MatchExperiment("homepage", &experimentContext)

	fmt.Println(experimentResult.Info())
	fmt.Println(experimentResult.GetExpId())

	fmt.Println(experimentResult.GetExperimentParams().GetString("version", "not exist"))
	fmt.Println(experimentResult.GetExperimentParams().GetString("rank_version", "not exist"))

}

func TestGetSceneParam(t *testing.T) {
	client := createExperimentClient(common.Environment_Daily_Desc)

	param := client.GetSceneParams("homepage").GetString("version", "not exist")
	fmt.Println(param)
}

/**
func TestGetFeatureConsistencyJob(t *testing.T) {
	host := "http://localhost:8000"
	client, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)), WithErrorLogger(LoggerFunc(log.Fatalf)))
	if err != nil {
		t.Fatal(err)
	}

	jobs := client.GetSceneParams("home_feed").GetFeatureConsistencyJobs()
	for _, job := range jobs {
		fmt.Println(job)
	}
}

func TestGetFlowCtrlPlanMetaList(t *testing.T) {
	host := "http://localhost:8000"
	client, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)), WithErrorLogger(LoggerFunc(log.Fatalf)))
	if err != nil {
		t.Fatal(err)
	}

	plans := client.GetFlowCtrlPlanMetaList("prepub", 0)
	for _, plan := range plans {
		fmt.Printf("%++v", plan)
	}
}

func TestGetFlowCtrlPlanTargetList(t *testing.T) {
	host := "http://localhost:8000"
	client, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)), WithErrorLogger(LoggerFunc(log.Fatalf)))
	if err != nil {
		t.Fatal(err)
	}

	targets := client.GetFlowCtrlPlanTargetList("prepub", "", 0)
	for planId, target := range targets {
		fmt.Printf("%d %+v", planId, target)
	}
}

func TestCheckIfFlowCtrlPlanIsEnabled(t *testing.T) {
	host := "http://localhost:8000"
	client, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)), WithErrorLogger(LoggerFunc(log.Fatalf)))
	if err != nil {
		t.Fatal(err)
	}

	enabled := client.CheckIfFlowCtrlPlanTargetIsEnabled("prepub", 9, 0)
	fmt.Println(enabled)
}

func TestCheckExperimentRoomDebugUsers(t *testing.T) {
	host := "http://localhost:8080"
	client, err := NewExperimentClient(host, common.Environment_Daily_Desc, WithLogger(LoggerFunc(log.Printf)))
	if err != nil {
		t.Error(err)
	}

	client.LoadExperimentData()
	sceneMap := client.sceneMap
	for _, scene := range sceneMap {
		for _, experimentRoom := range scene.ExperimentRooms {
			fmt.Println("experiment_room", experimentRoom.ExpRoomId)
			fmt.Println("DebugUsers", experimentRoom.DebugUsers)
			fmt.Println("DebugCrowdId", experimentRoom.DebugCrowdId)
			fmt.Println("DebugCrowdIdUsers", experimentRoom.DebugCrowdIdUsers)
		}
	}
}

func TestGetFlowCtrlPlanTargetTraffic(t *testing.T) {
	host := "http://localhost:8000"
	client, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)), WithErrorLogger(LoggerFunc(log.Fatalf)))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(client.GetFlowCtrlPlanTargetList("prepub", "test", 0))

	idList := []string{"ER_ALL", "12345678", "unknown"}
	fmt.Printf("%+v\n", client.GetFlowCtrlPlanTargetTraffic("prepub", "test", idList...))
}

**/
