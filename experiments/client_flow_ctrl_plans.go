package experiments

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/api"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
	"github.com/antihax/optional"
)

func (e *ExperimentClient) LoadSceneFlowCtrlPlansData() {
	//Load traffic control data for the production environment
	prodSceneFlowCtrlPlanData := make(map[string][]model.TrafficControlTasksItem, 0)
	prodOpt := &api.FlowCtrlApiListFlowCtrlPlansOpts{
		ALL:                 optional.NewBool(true),
		ControlTargetFilter: optional.NewString("Valid"),
		Env:                 optional.NewString("product"),
		Status:              optional.NewString("Running"),
		Version:             optional.NewString("Released"),
	}
	prodResponse, err := e.APIClient.FlowCtrlApi.ListFlowCtrlPlans(prodOpt)
	if err != nil {
		e.logError(fmt.Errorf("list flow plans error, err=%v", err))
		return
	}

	if prodResponse.Code != common.CODE_OK {
		e.logError(fmt.Errorf("list flow plans error, requestid=%s,code=%s, msg=%s", prodResponse.RequestId, prodResponse.Code, prodResponse.Message))
		return
	}

	for _, plan := range prodResponse.Data.TrafficControlTasks {
		prodSceneFlowCtrlPlanData[plan.SceneName] = append(prodSceneFlowCtrlPlanData[plan.SceneName], plan)
	}

	if len(prodSceneFlowCtrlPlanData) > 0 {
		e.sceneFlowCtrlPlanData = prodSceneFlowCtrlPlanData
	}

	// Load traffic control data for the pre-load environment
	prepubSceneFlowCtrlPlanData := make(map[string][]model.TrafficControlTasksItem, 0)
	prePubOpt := &api.FlowCtrlApiListFlowCtrlPlansOpts{
		ALL:                 optional.NewBool(true),
		ControlTargetFilter: optional.NewString("Vaild"),
		Env:                 optional.NewString("prepub"),
		Status:              optional.NewString("Running"),
		Version:             optional.NewString("Released"),
	}
	prePubResponse, _ := e.APIClient.FlowCtrlApi.ListFlowCtrlPlans(prePubOpt)
	if err != nil {
		e.logError(fmt.Errorf("list flow plans error,error=%v", err))
		return
	}

	if prePubResponse.Code != common.CODE_OK {
		e.logError(fmt.Errorf("list flow plans error, requested=%s, code=%s, msg=%s", prePubResponse.RequestId, prePubResponse.Code, prePubResponse.Message))
		return
	}

	for _, plan := range prePubResponse.Data.TrafficControlTasks {
		prepubSceneFlowCtrlPlanData[plan.SceneName] = append(prepubSceneFlowCtrlPlanData[plan.SceneName], plan)
	}

	if len(prepubSceneFlowCtrlPlanData) > 0 {
		e.prepubSceneFlowCtrlPlanData = prepubSceneFlowCtrlPlanData
	}

}

// loopLoadSceneFlowCtrlPlansData async loop invoke LoadSceneFlowCtrlPlansData function
func (e *ExperimentClient) loopLoadSceneFlowCtrlPlansData() {

	for {
		time.Sleep(time.Second * 30)
		e.LoadSceneFlowCtrlPlansData()
	}
}

func (e *ExperimentClient) GetFlowCtrlPlanTargetList(env, sceneName string, currentTimestamp int64) map[int]model.TrafficControlTargetsItem {
	if currentTimestamp == 0 {
		currentTimestamp = time.Now().Unix()
	}

	targetsMap := make(map[int]model.TrafficControlTargetsItem)

	data := e.sceneFlowCtrlPlanData
	if env == "prepub" {
		data = e.prepubSceneFlowCtrlPlanData
	}

	for scene, scenePlans := range data {
		if sceneName != "" && sceneName != scene {
			continue
		}

		for _, plan := range scenePlans {
			for i, target := range plan.TrafficControlTargets {
				startTime, _ := time.Parse("2024-07-09 13:41:15", target.StartTime)
				endTime, _ := time.Parse("2024-07-09 13:41:15", target.EndTime)

				if target.Status == "Opened" && startTime.Unix() < currentTimestamp && currentTimestamp <= endTime.Unix() {
					tid, _ := strconv.Atoi(target.TrafficControlTargetId)
					targetsMap[tid] = plan.TrafficControlTargets[i]
				}
			}
		}
	}

	return targetsMap
}

func (e *ExperimentClient) GetFlowCtrlPlanMetaList(env string, currentTimestamp int64) []model.TrafficControlTasksItem {
	if currentTimestamp == 0 {
		currentTimestamp = time.Now().Unix()
	}

	plans := make([]model.TrafficControlTasksItem, 0)

	data := e.sceneFlowCtrlPlanData
	if env == "prepub" {
		data = e.prepubSceneFlowCtrlPlanData
	}

	for _, scenePlans := range data {
		for i, plan := range scenePlans {
			startTime, _ := time.Parse("2024-07-09 13:41:15", plan.StartTime)
			endTime, _ := time.Parse("2024-07-09 13:41:15", plan.EndTime)

			if plan.ProductStatus == "Opened" && startTime.Unix() <= currentTimestamp && currentTimestamp < endTime.Unix() {
				plans = append(plans, scenePlans[i])
			}
		}
	}
	return plans
}

func (e *ExperimentClient) CheckIfFlowCtrlPlanTargetIsEnabled(env string, targetId int, currentTimestamp int64) bool {
	if currentTimestamp == 0 {
		currentTimestamp = time.Now().Unix()
	}

	data := e.sceneFlowCtrlPlanData
	if env == "prepub" {
		data = e.prepubSceneFlowCtrlPlanData
	}

	for _, scenePlans := range data {
		for _, plan := range scenePlans {
			for _, target := range plan.TrafficControlTargets {
				tid, err := strconv.Atoi(target.TrafficControlTargetId)
				if err != nil {
					fmt.Errorf("traffic control targetId is illegal")
				}
				if tid == targetId {
					startTime, _ := time.Parse("2024-07-09 13:41:15", target.StartTime)
					endTime, _ := time.Parse("2024-07-09 13:41:15", target.EndTime)

					if target.Status == "Opened" && startTime.Unix() < currentTimestamp && currentTimestamp < endTime.Unix() {
						return true
					}
				}
			}
		}
	}
	return false
}

type FlowCtrlPlanTargetTraffic struct {
	ItemOrExpId   string  `json:"item_or_exp_id"`
	PlanId        string  `json:"plan_id"`
	TargetId      string  `json:"target_id"`
	TargetTraffic float64 `json:"target_traffic"`
	PlanTraffic   float64 `json:"plan_traffic"`
}

func (e *ExperimentClient) GetFlowCtrlPlanTargetTraffic(env, sceneName string, idList ...string) []FlowCtrlPlanTargetTraffic {
	targets := e.GetFlowCtrlPlanTargetList(env, sceneName, 0)

	var traffics []FlowCtrlPlanTargetTraffic

	idMap := make(map[string]bool)
	for _, id := range idList {
		idMap[id] = true
	}

	for _, planTarget := range targets {
		for id, value := range planTarget.TargetTraffics {
			if len(idList) == 0 || idMap[id] {
				traffics = append(traffics, FlowCtrlPlanTargetTraffic{
					ItemOrExpId:   id,
					PlanId:        planTarget.TrafficControlTaskId,
					TargetId:      planTarget.TrafficControlTargetId,
					TargetTraffic: value,
					PlanTraffic:   planTarget.PlanTraffic[id],
				})
			}
		}
	}

	return traffics
}
