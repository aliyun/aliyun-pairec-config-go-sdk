package model

import (
	pairecv2 "github.com/alibabacloud-go/pairecservice-20221213/v3/client"
)

type TrafficControlTask struct {
	TrafficControlTaskId           string                  `json:"TrafficControlTaskId"`
	Name                           string                  `json:"Name"`
	Description                    string                  `json:"Description"`
	SceneId                        string                  `json:"SceneId"`
	SceneName                      string                  `json:"SceneName"`
	EffectiveSceneIds              []int                   `json:"EffectiveSceneIds"`
	EffectiveSceneNames            []string                `json:"EffectiveSceneNames"`
	ProductStatus                  string                  `json:"ProductStatus"`
	PrepubStatus                   string                  `json:"PrepubStatus"`
	ExecutionTime                  string                  `json:"ExecutionTime"`
	StartTime                      string                  `json:"StartTime"`
	EndTime                        string                  `json:"EndTime"`
	BehaviorTableMetaId            string                  `json:"BehaviorTableMetaId"`
	UserTableMetaId                string                  `json:"UserTableMetaId"`
	ItemTableMetaId                string                  `json:"ItemTableMetaId"`
	UserConditionType              string                  `json:"UserConditionType"`
	UserConditionArray             string                  `json:"UserConditionArray"`
	UserConditionExpress           string                  `json:"UserConditionExpress"`
	StatisBehaviorConditionType    string                  `json:"StatisBehaviorConditionType"`
	StatisBehaviorConditionArray   string                  `json:"StatisBehaviorConditionArray"`
	StatisBehaviorConditionExpress string                  `json:"StatisBehaviorConditionExpress"`
	ControlType                    string                  `json:"ControlType"`
	ControlGranularity             string                  `json:"ControlGranularity"`
	ControlLogic                   string                  `json:"ControlLogic"`
	ItemConditionType              string                  `json:"ItemConditionType"`
	ItemConditionArray             string                  `json:"ItemConditionArray"`
	ItemConditionExpress           string                  `json:"ItemConditionExpress"`
	ReleaseStage                   string                  `json:"ReleaseStage"`
	GmtCreateTime                  string                  `json:"GmtCreateTime"`
	GmtModifiedTime                string                  `json:"GmtModifiedTime"`
	EverPublished                  bool                    `json:"EverPublished"`
	ServiceId                      string                  `json:"ServiceId"`
	ServiceIds                     []int                   `json:"ServiceIds"`
	PreExperimentIds               string                  `json:"PreExperimentIds"`
	ProdExperimentIds              string                  `json:"ProdExperimentIds"`
	TrafficControlTargets          []*TrafficControlTarget `json:"TrafficControlTargets"`

	ActualTraffic TrafficControlActualTraffic `json:"ActualTraffic"`
}

func TrafficControlTaskConvert(trafficControlTask *pairecv2.ListTrafficControlTasksResponseBodyTrafficControlTasks) *TrafficControlTask {
	task := &TrafficControlTask{}

	task.TrafficControlTaskId = *trafficControlTask.TrafficControlTaskId
	task.Name = *trafficControlTask.Name
	task.Description = *trafficControlTask.Description
	task.SceneId = *trafficControlTask.SceneId
	task.SceneName = *trafficControlTask.SceneName

	for _, effectiveSceneId := range trafficControlTask.EffectiveSceneIds {
		task.EffectiveSceneIds = append(task.EffectiveSceneIds, int(*effectiveSceneId))
	}
	for _, effectiveSceneName := range trafficControlTask.EffectiveSceneNameList {
		task.EffectiveSceneNames = append(task.EffectiveSceneNames, *effectiveSceneName)
	}

	task.ProductStatus = *trafficControlTask.ProductStatus
	task.PrepubStatus = *trafficControlTask.PrepubStatus
	if trafficControlTask.ExecutionTime != nil {
		task.ExecutionTime = *trafficControlTask.ExecutionTime
	} else {
		task.ExecutionTime = ""
	}

	task.StartTime = *trafficControlTask.StartTime
	task.EndTime = *trafficControlTask.EndTime
	task.BehaviorTableMetaId = *trafficControlTask.BehaviorTableMetaId
	task.UserTableMetaId = *trafficControlTask.UserTableMetaId
	task.ItemTableMetaId = *trafficControlTask.ItemTableMetaId
	task.UserConditionType = *trafficControlTask.UserConditionType
	task.UserConditionArray = *trafficControlTask.UserConditionArray
	task.UserConditionExpress = *trafficControlTask.UserConditionExpress
	task.StatisBehaviorConditionType = *trafficControlTask.StatisBehaviorConditionType
	task.StatisBehaviorConditionArray = *trafficControlTask.StatisBehaviorConditionArray

	if trafficControlTask.StatisBahaviorConditionExpress != nil {
		task.StatisBehaviorConditionExpress = *trafficControlTask.StatisBahaviorConditionExpress
	} else {
		task.StatisBehaviorConditionExpress = ""
	}
	task.ControlType = *trafficControlTask.ControlType
	task.ControlGranularity = *trafficControlTask.ControlGranularity
	task.ControlLogic = *trafficControlTask.ControlLogic
	task.ItemConditionType = *trafficControlTask.ItemConditionType
	task.ItemConditionArray = *trafficControlTask.ItemConditionArray
	task.ItemConditionExpress = *trafficControlTask.ItemConditionExpress
	task.EverPublished = *trafficControlTask.EverPublished
	task.ServiceId = *trafficControlTask.ServiceId
	for _, serviceId := range trafficControlTask.ServiceIdList {
		task.ServiceIds = append(task.ServiceIds, int(*serviceId))
	}
	task.PreExperimentIds = *trafficControlTask.PreExperimentIds
	task.ProdExperimentIds = *trafficControlTask.ProdExperimentIds
	task.GmtCreateTime = *trafficControlTask.GmtCreateTime
	task.GmtModifiedTime = *trafficControlTask.GmtModifiedTime

	return task
}

type TrafficControlTarget struct {
	TrafficControlTaskId   string                         `json:"TrafficControlTaskId"`
	Name                   string                         `json:"Name"`
	Event                  string                         `json:"Event"`
	GmtModifiedTime        string                         `json:"GmtModifiedTime"`
	ToleranceValue         int64                          `json:"ToleranceValue"`
	Value                  float32                        `json:"Value"`
	TrafficControlTargetId string                         `json:"TrafficControlTargetId"`
	ItemConditionType      string                         `json:"ItemConditionType"`
	StartTime              string                         `json:"StartTime"`
	GmtCreateTime          string                         `json:"GmtCreateTime"`
	EndTime                string                         `json:"EndTime"`
	StatisPeriod           string                         `json:"StatisPeriod"`
	NewProductRegulation   bool                           `json:"NewProductRegulation"`
	ItemConditionArray     string                         `json:"ItemConditionArray"`
	Status                 string                         `json:"Status"`
	RecallName             string                         `json:"RecallName"`
	ItemConditionExpress   string                         `json:"ItemConditionExpress"`
	SplitParts             TrafficControlTargetSplitParts `json:"SplitParts"`
}

func TrafficControlTargetConvert(trafficControlTarget *pairecv2.ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargets) *TrafficControlTarget {
	target := &TrafficControlTarget{}
	target.TrafficControlTaskId = *trafficControlTarget.TrafficControlTaskId
	target.TrafficControlTargetId = *trafficControlTarget.TrafficControlTargetId
	target.Name = *trafficControlTarget.Name
	target.StartTime = *trafficControlTarget.StartTime
	target.EndTime = *trafficControlTarget.EndTime
	target.ItemConditionType = *trafficControlTarget.ItemConditionType
	target.ItemConditionArray = *trafficControlTarget.ItemConditionArray
	target.ItemConditionExpress = *trafficControlTarget.ItemConditionExpress
	target.Event = *trafficControlTarget.Event
	target.Value = *trafficControlTarget.Value
	target.StatisPeriod = *trafficControlTarget.StatisPeriod
	target.ToleranceValue = *trafficControlTarget.ToleranceValue
	target.NewProductRegulation = *trafficControlTarget.NewProductRegulation
	target.RecallName = *trafficControlTarget.RecallName
	target.Status = *trafficControlTarget.Status
	target.GmtCreateTime = *trafficControlTarget.GmtCreateTime
	target.GmtModifiedTime = *trafficControlTarget.GmtModifiedTime

	splitParts := TrafficControlTargetSplitPartsConvert(trafficControlTarget.SplitParts)
	target.SplitParts = splitParts
	return target
}

type TrafficControlTargetSplitParts struct {
	TimePoints []int   `json:"TimePoints"`
	SetValues  []int64 `json:"SetValues"`
}

func TrafficControlTargetSplitPartsConvert(trafficControlTargetSplitParts *pairecv2.ListTrafficControlTasksResponseBodyTrafficControlTasksTrafficControlTargetsSplitParts) TrafficControlTargetSplitParts {
	splitParts := TrafficControlTargetSplitParts{}
	for _, timePoint := range trafficControlTargetSplitParts.TimePoints {
		splitParts.TimePoints = append(splitParts.TimePoints, int(*timePoint))
	}
	for _, setValue := range trafficControlTargetSplitParts.SetValues {
		splitParts.SetValues = append(splitParts.SetValues, *setValue)
	}
	return splitParts
}

type TrafficControlActualTraffic struct {
	TaskTraffics   map[string]TaskTrafficDetail `json:"TaskTraffics"`
	TargetTraffics []*TargetTraffic             `json:"TargetTraffics"`
}
type TaskTrafficDetail struct {
	Traffic float64 `json:"Traffic"`
}

type TargetTraffic struct {
	TrafficControlTargetId string                         `json:"TrafficContorlTargetId"`
	Data                   map[string]TargetTrafficDetail `json:"Data"`
}

type TargetTrafficDetail struct {
	Traffic    float64 `json:"Traffic"`
	RecordTime int64   `json:"RecordTime"`
}

func ActualTrafficConvert(trafficInfo *pairecv2.GetTrafficControlTaskTrafficResponseBodyTrafficControlTaskTrafficInfo) TrafficControlActualTraffic {

	actualTraffic := TrafficControlActualTraffic{}

	taskTraffic := make(map[string]TaskTrafficDetail, 0)
	for k, v := range trafficInfo.TaskTraffics {
		taskDetail := TaskTrafficDetail{
			Traffic: *v.Traffic,
		}
		taskTraffic[k] = taskDetail
	}

	actualTraffic.TaskTraffics = taskTraffic

	targetTraffics := make([]*TargetTraffic, 0)
	for _, value := range trafficInfo.TargetTraffics {
		targetTraffic := &TargetTraffic{
			TrafficControlTargetId: *value.TrafficContorlTargetId,
		}

		data := make(map[string]TargetTrafficDetail, 0)
		for k, v := range value.Data {
			targetDetail := TargetTrafficDetail{
				Traffic:    *v.Traffic,
				RecordTime: *v.RecordTime,
			}
			data[k] = targetDetail
		}
		targetTraffic.Data = data

		targetTraffics = append(targetTraffics, targetTraffic)
	}

	actualTraffic.TargetTraffics = targetTraffics

	return actualTraffic
}
