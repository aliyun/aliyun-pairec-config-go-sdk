package experiments

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/aliyun/aliyun-pairec-config-go-sdk/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/model"
)

func TestCreateExperimentClient(t *testing.T) {
	host := "http://localhost:8080"
	client, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(client)
}
func TestLoadExperimentData(t *testing.T) {
	host := "http://localhost:8080"
	_, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)))
	if err != nil {
		t.Error(err)
	}

	// client.LoadExperimentData()
}

func TestMatchExperiment(t *testing.T) {
	host := "http://localhost:8080"
	client, err := NewExperimentClient(host, common.Environment_Daily_Desc, WithLogger(LoggerFunc(log.Printf)), WithErrorLogger(LoggerFunc(log.Fatalf)))
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		Uid    string
		Result string
	}{}

	cases = append(cases, struct {
		Uid    string
		Result string
	}{Uid: strconv.Itoa(2125), Result: "base"})

	cases = append(cases, struct {
		Uid    string
		Result string
	}{Uid: strconv.Itoa(1000), Result: "test"})

	cases = append(cases, struct {
		Uid    string
		Result string
	}{Uid: strconv.Itoa(1004), Result: "base"})

	for _, tc := range cases {
		experimentContext := model.ExperimentContext{
			RequestId: "pvid",
			Uid:       tc.Uid,
			FilterParams: map[string]interface{}{
				"sex": "male",
				//"age": 35,
				"uid": 1,
			},
		}
		experimentResult := client.MatchExperiment("test", &experimentContext)

		fmt.Println(experimentResult.Info())
		r := experimentResult.GetExperimentParams().GetString("a", "not exist")
		if r != tc.Result {
			t.Errorf("expect:%s, result:%s", tc.Result, r)
		}
	}

}

func TestMatchExperiment2(t *testing.T) {
	host := "http://1730760139076263.cn-beijing.pai-eas.aliyuncs.com/api/predict/pairec_experiment"
	client, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)), WithErrorLogger(LoggerFunc(log.Fatalf)), WithToken("xxxx"))
	if err != nil {
		t.Fatal(err)
	}

	experimentContext := model.ExperimentContext{
		RequestId: "pvid",
		//Uid:       "211789768",
		Uid: "102441835",
		FilterParams: map[string]interface{}{
			"sex": "male",
			//"age": 35,
			"uid": 1,
		},
	}

	experimentResult := client.MatchExperiment("homepage", &experimentContext)

	fmt.Println(experimentResult.Info())
	fmt.Println(experimentResult.GetExpId())

	fmt.Println(experimentResult.GetLayerParams("rank").GetInt("intval", 0))
	fmt.Println(experimentResult.GetExperimentParams().GetInt("intval", 0))
	fmt.Println(experimentResult.GetExperimentParams().GetString("name", "not exist"))
	fmt.Println(experimentResult.GetExperimentParams().GetString("recall", "not exist"))

}
func TestGetSceneParam(t *testing.T) {
	host := "http://localhost:8000"
	client, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)), WithErrorLogger(LoggerFunc(log.Fatalf)))
	if err != nil {
		t.Fatal(err)
	}

	param := client.GetSceneParams("home_feed").GetString("_feature_consistency_job_", "not exist")
	fmt.Println(param)
}
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
