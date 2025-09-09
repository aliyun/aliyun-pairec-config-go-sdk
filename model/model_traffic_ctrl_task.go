package model

import (
	pairecv2 "github.com/alibabacloud-go/pairecservice-20221213/v3/client"
)

type TrafficControlTask struct {
	TrafficControlTaskId           string                  `json:"TrafficControlTaskId" xml:"TrafficControlTaskId"`
	Name                           string                  `json:"Name" xml:"Name"`
	Description                    string                  `json:"Description" xml:"Description"`
	SceneId                        string                  `json:"SceneId" xml:"SceneId"`
	SceneName                      string                  `json:"SceneName" xml:"SceneName"`
	EffectiveSceneIds              []int                   `json:"EffectiveSceneIds"`
	EffectiveSceneNames            []string                `json:"EffectiveSceneNames"`
	ProductStatus                  string                  `json:"ProductStatus" xml:"ProductStatus"`
	PrepubStatus                   string                  `json:"PrepubStatus" xml:"PrepubStatus"`
	ExecutionTime                  string                  `json:"ExecutionTime" xml:"ExecutionTime"`
	StartTime                      string                  `json:"StartTime" xml:"StartTime"`
	EndTime                        string                  `json:"EndTime" xml:"EndTime"`
	BehaviorTableMetaId            string                  `json:"BehaviorTableMetaId" xml:"BehaviorTableMetaId"`
	UserTableMetaId                string                  `json:"UserTableMetaId" xml:"UserTableMetaId"`
	ItemTableMetaId                string                  `json:"ItemTableMetaId" xml:"ItemTableMetaId"`
	UserConditionType              string                  `json:"UserConditionType" xml:"UserConditionType"`
	UserConditionArray             string                  `json:"UserConditionArray" xml:"UserConditionArray"`
	UserConditionExpress           string                  `json:"UserConditionExpress" xml:"UserConditionExpress"`
	StatisBehaviorConditionType    string                  `json:"StatisBehaviorConditionType" xml:"StatisBehaviorConditionType"`
	StatisBehaviorConditionArray   string                  `json:"StatisBehaviorConditionArray" xml:"StatisBehaviorConditionArray"`
	StatisBehaviorConditionExpress string                  `json:"StatisBehaviorConditionExpress" xml:"StatisBehaviorConditionExpress"`
	ControlType                    string                  `json:"ControlType" xml:"ControlType"`
	ControlGranularity             string                  `json:"ControlGranularity" xml:"ControlGranularity"`
	ControlLogic                   string                  `json:"ControlLogic" xml:"ControlLogic"`
	ItemConditionType              string                  `json:"ItemConditionType" xml:"ItemConditionType"`
	ItemConditionArray             string                  `json:"ItemConditionArray" xml:"ItemConditionArray"`
	ItemConditionExpress           string                  `json:"ItemConditionExpress" xml:"ItemConditionExpress"`
	ReleaseStage                   string                  `json:"ReleaseStage"`
	GmtCreateTime                  string                  `json:"GmtCreateTime" xml:"GmtCreateTime"`
	GmtModifiedTime                string                  `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
	EverPublished                  bool                    `json:"EverPublished"`
	ServiceId                      string                  `json:"ServiceId"`
	ServiceIds                     []int                   `json:"ServiceIds"`
	PreExperimentIds               string                  `json:"PreExperimentIds"`
	ProdExperimentIds              string                  `json:"ProdExperimentIds"`
	TrafficControlTargets          []*TrafficControlTarget `json:"TrafficControlTargets" xml:"TrafficControlTargets"`

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
	task.ExecutionTime = *trafficControlTask.ExecutionTime
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
	TrafficControlTaskId   string                         `json:"TrafficControlTaskId" xml:"TrafficControlTaskId"`
	Name                   string                         `json:"Name" xml:"Name"`
	Event                  string                         `json:"Event" xml:"Event"`
	GmtModifiedTime        string                         `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
	ToleranceValue         int64                          `json:"ToleranceValue" xml:"ToleranceValue"`
	Value                  float32                        `json:"Value" xml:"Value"`
	TrafficControlTargetId string                         `json:"TrafficControlTargetId" xml:"TrafficControlTargetId"`
	ItemConditionType      string                         `json:"ItemConditionType" xml:"ItemConditionType"`
	StartTime              string                         `json:"StartTime" xml:"StartTime"`
	GmtCreateTime          string                         `json:"GmtCreateTime" xml:"GmtCreateTime"`
	EndTime                string                         `json:"EndTime" xml:"EndTime"`
	StatisPeriod           string                         `json:"StatisPeriod" xml:"StatisPeriod"`
	NewProductRegulation   bool                           `json:"NewProductRegulation" xml:"NewProductRegulation"`
	ItemConditionArray     string                         `json:"ItemConditionArray" xml:"ItemConditionArray"`
	Status                 string                         `json:"Status" xml:"Status"`
	RecallName             string                         `json:"RecallName" xml:"RecallName"`
	ItemConditionExpress   string                         `json:"ItemConditionExpress" xml:"ItemConditionExpress"`
	SplitParts             TrafficControlTargetSplitParts `json:"SplitParts" xml:"SplitParts"`
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
