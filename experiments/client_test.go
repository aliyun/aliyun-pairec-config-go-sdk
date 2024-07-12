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
	//address := "pairecservice.cn-hangzhou.aliyuncs.com"
	//preAddress :=
	client, err := NewExperimentClient(instanceId, region, accessId, accessKey, environment,
		WithLogger(LoggerFunc(log.Printf)),
		WithDomain("pairecservice-pre.cn-hangzhou.aliyuncs.com"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func TestMatchExperiment(t *testing.T) {
	client := createExperimentClient(common.Environment_Product_Desc)

	experimentContext := model.ExperimentContext{
		RequestId: "pvid",
		Uid:       "1034416388",
		FilterParams: map[string]interface{}{
			"country": "new12",
		},
	}

	experimentResult := client.MatchExperiment("home_feed", &experimentContext)

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

func TestGetFeatureConsistencyJob(t *testing.T) {
	client := createExperimentClient(common.Environment_Prepub_Desc)

	jobs := client.GetSceneParams("home_feed").GetFeatureConsistencyJobs()
	for _, job := range jobs {
		fmt.Println(job)
	}
}

func TestFeatureConsistencyBackflow(t *testing.T) {
	client := createExperimentClient(common.Environment_Prepub_Desc)
	backflowData := model.FeatureConsistencyBackflowData{
		FeatureConsistencyCheckJobConfigId: "1",
		LogUserId:                          "100000081",
		LogItemId:                          "[\"262850386\",\"249988426\"]",
		UserFeatures:                       "",
		LogRequestId:                       "1130c79b-4375-4288-8b00-e575d645554f",
		SceneName:                          "home_feed",
	}
	resp, err := client.BackflowFeatureConsistencyCheckJobData(&backflowData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp, err)
}

func TestFeatureConsistencyReply(t *testing.T) {
	client := createExperimentClient(common.Environment_Prepub_Desc)
	replyData := model.FeatureConsistencyReplyData{
		FeatureConsistencyCheckJobConfigId: "1",
		LogUserId:                          "100000081",
		LogItemId:                          "[\"262850386\",\"249988426\"]",
		LogRequestId:                       "1130c79b-4375-4288-8b00-e575d645554f",
		SceneName:                          "home_feed",
	}
	resp, err := client.SyncFeatureConsistencyCheckJobReplayLog(&replyData)
	fmt.Println(resp, err)
	if err != nil {
		log.Fatal(err)
	}
}

// /**

func TestGetTrafficControlTaskMetaData(t *testing.T) {
	client := createExperimentClient(common.Environment_Prepub_Desc)
	plans := client.GetTrafficControlTaskMetaData("product", 0)
	fmt.Println("-----------")
	for _, plan := range plans {
		fmt.Printf("%++v", plan)
	}
}

func TestGetTrafficControlTargetData(t *testing.T) {
	client := createExperimentClient(common.Environment_Prepub_Desc)
	targets := client.GetTrafficControlTargetData("prepub", "", 0)
	for planId, target := range targets {
		fmt.Printf("%d %+v", planId, target)
	}
}

func TestCheckIfTrafficControlTargetIsEnabled(t *testing.T) {
	client := createExperimentClient(common.Environment_Prepub_Desc)
	enabled := client.CheckIfTrafficControlTargetIsEnabled("prepub", 45, 0)
	fmt.Println(enabled)
}

func TestCheckExperimentRoomDebugUsers(t *testing.T) {
	client := createExperimentClient(common.Environment_Prepub_Desc)
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

func TestGetTrafficControlTargetTraffic(t *testing.T) {
	client := createExperimentClient(common.Environment_Prepub_Desc)
	fmt.Println(client.GetTrafficControlTargetTraffic("prepub", "test1", ""))

	idList := []string{"ER_ALL", "12345678", "unknown"}
	fmt.Printf("%+v\n", client.GetTrafficControlTargetTraffic("prepub", "test1", idList...))
}
