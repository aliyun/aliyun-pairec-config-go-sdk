package experiments

import (
	"fmt"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/api"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
	"github.com/antihax/optional"
	"strconv"
	"time"
)

func (e *ExperimentClient) LoadSceneTrafficControlTasksData() {
	//Load traffic control data for the production environment
	prodSceneTrafficControlTasksData := make(map[string][]model.TrafficControlTask, 0)
	prodOpt := &api.TrafficControlApiListTrafficControlTasksOpts{
		ALL:                 optional.NewBool(true),
		ControlTargetFilter: optional.NewString("Valid"),
		Env:                 optional.NewString("product"),
		Status:              optional.NewString("Running"),
		Version:             optional.NewString("Released"),
	}
	prodResponse, err := e.APIClient.FlowCtrlApi.ListTrafficControlTasks(prodOpt)
	if err != nil {
		e.logError(fmt.Errorf("list flow plans error, err=%v", err))
		return
	}

	for _, plan := range prodResponse.TrafficControlTasks {
		prodSceneTrafficControlTasksData[plan.SceneName] = append(prodSceneTrafficControlTasksData[plan.SceneName], plan)
	}

	if len(prodSceneTrafficControlTasksData) > 0 {
		e.productSceneTrafficControlTaskData = prodSceneTrafficControlTasksData
	}

	// Load traffic control data for the pre-load environment
	prepubSceneFlowCtrlPlanData := make(map[string][]model.TrafficControlTask, 0)
	prePubOpt := &api.TrafficControlApiListTrafficControlTasksOpts{
		ALL:                 optional.NewBool(true),
		ControlTargetFilter: optional.NewString("Valid"),
		Env:                 optional.NewString("prepub"),
		Status:              optional.NewString("Running"),
		Version:             optional.NewString("Released"),
	}
	prePubResponse, _ := e.APIClient.FlowCtrlApi.ListTrafficControlTasks(prePubOpt)
	if err != nil {
		e.logError(fmt.Errorf("list flow plans error,error=%v", err))
		return
	}

	for _, plan := range prePubResponse.TrafficControlTasks {
		prepubSceneFlowCtrlPlanData[plan.SceneName] = append(prepubSceneFlowCtrlPlanData[plan.SceneName], plan)
	}

	if len(prepubSceneFlowCtrlPlanData) > 0 {
		e.prepubSceneTrafficControlTaskData = prepubSceneFlowCtrlPlanData
	}

}

// loopLoadSceneFlowCtrlPlansData async loop invoke LoadSceneFlowCtrlPlansData function
func (e *ExperimentClient) loopLoadSceneFlowCtrlPlansData() {

	for {
		time.Sleep(time.Second * 30)
		e.LoadSceneTrafficControlTasksData()
	}
}

func (e *ExperimentClient) GetTrafficControlTargetData(env, sceneName string, currentTimestamp int64) map[int]model.TrafficControlTarget {
	if currentTimestamp == 0 {
		currentTimestamp = time.Now().Unix()
	}

	trafficControlTargets := make(map[int]model.TrafficControlTarget)

	data := e.productSceneTrafficControlTaskData
	if env == "prepub" {
		data = e.prepubSceneTrafficControlTaskData
	}

	for scene, sceneTraffics := range data {
		if sceneName != "" && sceneName != scene {
			continue
		}

		for _, traffic := range sceneTraffics {
			for i, target := range traffic.TrafficControlTargets {
				startTime, _ := time.Parse(time.RFC3339, target.StartTime)
				endTime, _ := time.Parse(time.RFC3339, target.EndTime)

				if target.Status == "Opened" && startTime.Unix() < currentTimestamp && currentTimestamp <= endTime.Unix() {
					tid, _ := strconv.Atoi(target.TrafficControlTargetId)
					trafficControlTargets[tid] = traffic.TrafficControlTargets[i]
				}
			}
		}
	}

	return trafficControlTargets
}

func (e *ExperimentClient) GetTrafficControlTaskMetaData(env string, currentTimestamp int64) []model.TrafficControlTask {
	if currentTimestamp == 0 {
		currentTimestamp = time.Now().Unix()
	}

	traffics := make([]model.TrafficControlTask, 0)

	data := e.productSceneTrafficControlTaskData

	if env == "prepub" {
		data = e.prepubSceneTrafficControlTaskData
	}

	for _, sceneTraffics := range data {
		for i, traffic := range sceneTraffics {
			startTime, _ := time.Parse(time.RFC3339, traffic.StartTime)
			endTime, _ := time.Parse(time.RFC3339, traffic.EndTime)

			if traffic.ProductStatus == "Running" && startTime.Unix() <= currentTimestamp && currentTimestamp < endTime.Unix() {
				traffics = append(traffics, sceneTraffics[i])
			}
		}
	}
	return traffics
}

func (e *ExperimentClient) CheckIfTrafficControlTargetIsEnabled(env string, targetId int, currentTimestamp int64) bool {
	if currentTimestamp == 0 {
		currentTimestamp = time.Now().Unix()
	}

	data := e.productSceneTrafficControlTaskData
	if env == "prepub" {
		data = e.prepubSceneTrafficControlTaskData
	}

	for _, sceneTraffics := range data {
		for _, traffic := range sceneTraffics {
			for _, target := range traffic.TrafficControlTargets {
				tid, err := strconv.Atoi(target.TrafficControlTargetId)
				if err != nil {
					e.logError(fmt.Errorf("traffic control targetId is illegal"))
				}
				if tid == targetId {
					startTime, _ := time.Parse(time.RFC3339, target.StartTime)
					endTime, _ := time.Parse(time.RFC3339, target.EndTime)

					if target.Status == "Running" && startTime.Unix() < currentTimestamp && currentTimestamp < endTime.Unix() {
						return true
					}
				}
			}
		}
	}
	return false
}

type TrafficControlTargetTraffic struct {
	ItemOrExpId            string  `json:"item_or_exp_id"`
	TrafficControlTaskId   string  `json:"plan_id"`
	TrafficControlTargetId string  `json:"target_id"`
	TargetTraffic          float64 `json:"target_traffic"`
	PlanTraffic            float64 `json:"plan_traffic"`
}

func (e *ExperimentClient) GetTrafficControlTargetTraffic(env, sceneName string, idList ...string) []TrafficControlTargetTraffic {
	targets := e.GetTrafficControlTargetData(env, sceneName, 0)

	var traffics []TrafficControlTargetTraffic

	idMap := make(map[string]bool)
	for _, id := range idList {
		idMap[id] = true
	}

	for _, trafficTarget := range targets {
		for id, value := range trafficTarget.TargetTraffics {
			if len(idList) == 0 || idMap[id] {
				traffics = append(traffics, TrafficControlTargetTraffic{
					ItemOrExpId:            id,
					TrafficControlTaskId:   trafficTarget.TrafficControlTaskId,
					TrafficControlTargetId: trafficTarget.TrafficControlTargetId,
					TargetTraffic:          value,
					PlanTraffic:            trafficTarget.PlanTraffic[id],
				})
			}
		}
	}

	return traffics
}
