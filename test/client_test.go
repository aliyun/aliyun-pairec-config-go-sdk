package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
)

func TestMatchExperiment(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Product_Desc)

	experimentContext := model.ExperimentContext{
		RequestId: "pvid",
		Uid:       "1034416498",
		FilterParams: map[string]interface{}{
			"country": "new12",
		},
	}

	experimentResult := client.MatchExperiment("HomePage", &experimentContext)

	fmt.Println(experimentResult.Info())
	fmt.Println(experimentResult.GetExpId())

	fmt.Println(experimentResult.GetExperimentParams().GetString("recall_param", "not exist"))
	fmt.Println(experimentResult.GetExperimentParams().GetString("rank_param", "not exist"))

}

func TestGetSceneParam(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Product_Desc)

	param := client.GetSceneParams("zhenghong_scene").GetString("test", "not exist")
	fmt.Println(param)
	params := client.GetSceneParams("pairec_config_scene").ListParams()
	t.Log(params)
}

func TestGetFeatureConsistencyJob(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Product_Desc)

	jobs := client.GetSceneParams("home_feed").GetFeatureConsistencyJobs()
	for _, job := range jobs {
		fmt.Println(job)
	}
}

func TestFeatureConsistencyBackflow(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	backflowData := model.FeatureConsistencyBackflowData{
		FeatureConsistencyCheckJobConfigId: "1",
		LogUserId:                          "100000081",
		LogItemId:                          "[\"262850386\",\"249988426\"]",
		UserFeatures:                       "",
		LogRequestId:                       "1130c79b-4375-4288-8b00-e575d645554f",
		SceneName:                          "home_feed",
		ServiceName:                        "embedding_recall",
	}
	resp, err := client.BackflowFeatureConsistencyCheckJobData(&backflowData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp, err)
}

func TestFeatureConsistencyReply(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
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

func TestCheckExperimentRoomDebugUsers(t *testing.T) {
	client := CreateExperimentClient(common.Environment_Prepub_Desc)
	client.LoadExperimentData()
	sceneMap := client.SceneMap
	for _, scene := range sceneMap {
		for _, experimentRoom := range scene.ExperimentRooms {
			fmt.Println("experiment_room", experimentRoom.ExpRoomId)
			fmt.Println("DebugUsers", experimentRoom.DebugUsers)
			fmt.Println("DebugCrowdId", experimentRoom.DebugCrowdId)
			fmt.Println("DebugCrowdIdUsers", experimentRoom.DebugCrowdIdUsers)
		}
	}
}
