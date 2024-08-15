package experiments

import (
	"fmt"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/api"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
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
	prodResponse, err := e.APIClient.TrafficControlApi.ListTrafficControlTasks(prodOpt)
	if err != nil {
		e.logError(fmt.Errorf("list flow tasks error, err=%v", err))
		return
	}

	for _, task := range prodResponse.TrafficControlTasks {
		prodSceneTrafficControlTasksData[task.SceneName] = append(prodSceneTrafficControlTasksData[task.SceneName], task)
	}

	if len(prodSceneTrafficControlTasksData) > 0 {
		e.productSceneTrafficControlTaskData = prodSceneTrafficControlTasksData
	}

	// Load traffic control data for the pre-load environment
	prepubSceneTrafficControlTasksData := make(map[string][]model.TrafficControlTask, 0)
	prePubOpt := &api.TrafficControlApiListTrafficControlTasksOpts{
		ALL:                 optional.NewBool(true),
		ControlTargetFilter: optional.NewString("Valid"),
		Env:                 optional.NewString("prepub"),
		Status:              optional.NewString("Running"),
		Version:             optional.NewString("Released"),
	}
	prePubResponse, _ := e.APIClient.TrafficControlApi.ListTrafficControlTasks(prePubOpt)
	if err != nil {
		e.logError(fmt.Errorf("list flow tasks error,error=%v", err))
		return
	}

	for _, task := range prePubResponse.TrafficControlTasks {
		prepubSceneTrafficControlTasksData[task.SceneName] = append(prepubSceneTrafficControlTasksData[task.SceneName], task)
	}

	if len(prepubSceneTrafficControlTasksData) > 0 {
		e.prepubSceneTrafficControlTaskData = prepubSceneTrafficControlTasksData
	}

}

// loopLoadSceneFlowCtrlPlansData async loop invoke LoadSceneFlowCtrlPlansData function
func (e *ExperimentClient) loopLoadSceneFlowCtrlPlansData() {

	for {
		time.Sleep(time.Second * 30)
		e.LoadSceneTrafficControlTasksData()
	}
}

func (e *ExperimentClient) SetTrafficControlTraffic(trafficData model.TrafficControlTaskTrafficData) (string, error) {
	response, err := e.APIClient.TrafficControlApi.SetTrafficControlTrafficFData(trafficData)
	return response, err
}

func (e *ExperimentClient) GetTrafficControlTargetData(env, sceneName string, currentTimestamp int64) map[string]model.TrafficControlTarget {
	if currentTimestamp == 0 {
		currentTimestamp = time.Now().Unix()
	}

	trafficControlTargets := make(map[string]model.TrafficControlTarget)

	var data = make(map[string][]model.TrafficControlTask, 0)

	if env == common.Environment_Prepub_Desc {
		data = e.prepubSceneTrafficControlTaskData
	} else if env == common.Environment_Product_Desc {
		data = e.productSceneTrafficControlTaskData
	} else {
		return nil
	}

	for scene, sceneTraffics := range data {
		if sceneName != "" && sceneName != scene {
			continue
		}

		for _, task := range sceneTraffics {
			for i, target := range task.TrafficControlTargets {
				if task.ExecutionTime != "Permanent" {
					startTime, _ := time.Parse(time.RFC3339, target.StartTime)
					endTime, _ := time.Parse(time.RFC3339, target.EndTime)

					if startTime.Unix() >= endTime.Unix() {
						e.logError(fmt.Errorf("The subtarget time for %s's traffic control task is incorrect. \n", task.Name))
					}

					if target.Status == common.TrafficControlTargets_Status_Open && startTime.Unix() < currentTimestamp && currentTimestamp <= endTime.Unix() {
						trafficControlTargets[target.TrafficControlTargetId] = task.TrafficControlTargets[i]
					}
				} else {
					if target.Status == common.TrafficControlTargets_Status_Open {
						trafficControlTargets[target.TrafficControlTargetId] = task.TrafficControlTargets[i]
					}
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

	var data = make(map[string][]model.TrafficControlTask, 0)

	if env == common.Environment_Prepub_Desc {
		data = e.prepubSceneTrafficControlTaskData
	} else if env == common.Environment_Product_Desc {
		data = e.productSceneTrafficControlTaskData
	} else {
		return nil
	}

	for _, sceneTraffics := range data {
		for i, task := range sceneTraffics {
			if task.ExecutionTime != "Permanent" {
				startTime, _ := time.Parse(time.RFC3339, task.StartTime)
				endTime, _ := time.Parse(time.RFC3339, task.EndTime)

				if startTime.Unix() >= endTime.Unix() {
					e.logError(fmt.Errorf("The traffic control task time of %s is incorrect. \n", task.Name))
				}

				if env == common.Environment_Product_Desc {
					if task.ProductStatus == common.TrafficCtrlTask_Running_Status && startTime.Unix() <= currentTimestamp && currentTimestamp < endTime.Unix() {
						traffics = append(traffics, sceneTraffics[i])
					}
				} else if env == common.Environment_Prepub_Desc {
					if task.PrepubStatus == common.TrafficCtrlTask_Running_Status && startTime.Unix() <= currentTimestamp && currentTimestamp < endTime.Unix() {
						traffics = append(traffics, sceneTraffics[i])
					}

				}
			} else {
				if env == common.Environment_Product_Desc {
					if task.ProductStatus == common.TrafficCtrlTask_Running_Status {
						traffics = append(traffics, sceneTraffics[i])
					}
				} else if env == common.Environment_Prepub_Desc {
					if task.PrepubStatus == common.TrafficCtrlTask_Running_Status {
						traffics = append(traffics, sceneTraffics[i])
					}
				}
			}

		}
	}
	return traffics
}

func (e *ExperimentClient) CheckIfTrafficControlTargetIsEnabled(env string, targetId int, currentTimestamp int64) bool {
	if currentTimestamp == 0 {
		currentTimestamp = time.Now().Unix()
	}

	var data = make(map[string][]model.TrafficControlTask, 0)

	if env == common.Environment_Prepub_Desc {
		data = e.prepubSceneTrafficControlTaskData
	} else if env == common.Environment_Product_Desc {
		data = e.productSceneTrafficControlTaskData
	} else {
		return false
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

					if startTime.Unix() >= endTime.Unix() {
						e.logError(fmt.Errorf("The traffic control task time is incorrect. Procedure\n"))
					}

					if target.Status == common.TrafficControlTargets_Status_Open && startTime.Unix() < currentTimestamp && currentTimestamp < endTime.Unix() {
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
	TrafficControlTaskId   string  `json:"traffic_control_task_id"`
	TrafficControlTargetId string  `json:"traffic_control_target_id"`
	TargetTraffic          float64 `json:"target_traffic"`
	TaskTraffic            float64 `json:"task_traffic"`
}

func (e *ExperimentClient) GetTrafficControlTargetTraffic(env, sceneName string, idList ...string) []TrafficControlTargetTraffic {
	targets := e.GetTrafficControlTargetData(env, sceneName, 0)

	var traffics []TrafficControlTargetTraffic

	idMap := make(map[string]bool)
	for _, id := range idList {
		if id == "" {
			continue
		}
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
					TaskTraffic:            trafficTarget.TaskTraffics[id],
				})
			}
		}
	}

	return traffics
}
