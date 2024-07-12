package model

import "github.com/aliyun/alibaba-cloud-sdk-go/services/pairecservice"

type TrafficControlTask struct {
	TrafficControlTaskId string `json:"TrafficControlTaskId" xml:"TrafficControlTaskId"`
	Name                 string `json:"Name" xml:"Name"`
	Description          string `json:"Description" xml:"Description"`
	SceneId              string `json:"SceneId" xml:"SceneId"`
	SceneName            string `json:"SceneName" xml:"SceneName"`
	ProductStatus        string `json:"ProductStatus" xml:"ProductStatus"`
	PrepubStatus         string `json:"PrepubStatus" xml:"PrepubStatus"`
	ExecutionTime        string `json:"ExecutionTime" xml:"ExecutionTime"`
	StartTime            string `json:"StartTime" xml:"StartTime"`
	EndTime              string `json:"EndTime" xml:"EndTime"`

	BehaviorTableMetaId string `json:"BehaviorTableMetaId" xml:"BehaviorTableMetaId"`
	UserTableMetaId     string `json:"UserTableMetaId" xml:"UserTableMetaId"`
	ItemTableMetaId     string `json:"ItemTableMetaId" xml:"ItemTableMetaId"`

	BehaviorTableMeta *pairecservice.TableMetasItem `json:"BehaviorTableMeta"`
	UserTableMeta     *pairecservice.TableMetasItem `json:"UserTableMeta"`
	ItemTableMeta     *pairecservice.TableMetasItem `json:"ItemTableMeta"`

	UserConditionType              string                 `json:"UserConditionType" xml:"UserConditionType"`
	UserConditionArray             string                 `json:"UserConditionArray" xml:"UserConditionArray"`
	UserConditionExpress           string                 `json:"UserConditionExpress" xml:"UserConditionExpress"`
	StatisBehaviorConditionType    string                 `json:"StatisBehaviorConditionType" xml:"StatisBehaviorConditionType"`
	StatisBehaviorConditionArray   string                 `json:"StatisBehaviorConditionArray" xml:"StatisBehaviorConditionArray"`
	StatisBahaviorConditionExpress string                 `json:"StatisBahaviorConditionExpress" xml:"StatisBahaviorConditionExpress"`
	ControlType                    string                 `json:"ControlType" xml:"ControlType"`
	ControlGranularity             string                 `json:"ControlGranularity" xml:"ControlGranularity"`
	ControlLogic                   string                 `json:"ControlLogic" xml:"ControlLogic"`
	ItemConditionType              string                 `json:"ItemConditionType" xml:"ItemConditionType"`
	ItemConditionArray             string                 `json:"ItemConditionArray" xml:"ItemConditionArray"`
	ItemConditionExpress           string                 `json:"ItemConditionExpress" xml:"ItemConditionExpress"`
	GmtCreateTime                  string                 `json:"GmtCreateTime" xml:"GmtCreateTime"`
	GmtModifiedTime                string                 `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
	EverPublished                  bool                   `json:"EverPublished" xml:"EverPublished"`
	TrafficControlTargets          []TrafficControlTarget `json:"TrafficControlTargets" xml:"TrafficControlTargets"`
}

type TrafficControlTarget struct {
	TrafficControlTaskId   string                                          `json:"TrafficControlTaskId" xml:"TrafficControlTaskId"`
	Name                   string                                          `json:"Name" xml:"Name"`
	Event                  string                                          `json:"Event" xml:"Event"`
	GmtModifiedTime        string                                          `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
	ToleranceValue         int64                                           `json:"ToleranceValue" xml:"ToleranceValue"`
	Value                  float64                                         `json:"Value" xml:"Value"`
	TrafficControlTargetId string                                          `json:"TrafficControlTargetId" xml:"TrafficControlTargetId"`
	ItemConditionType      string                                          `json:"ItemConditionType" xml:"ItemConditionType"`
	StartTime              string                                          `json:"StartTime" xml:"StartTime"`
	GmtCreateTime          string                                          `json:"GmtCreateTime" xml:"GmtCreateTime"`
	EndTime                string                                          `json:"EndTime" xml:"EndTime"`
	StatisPeriod           string                                          `json:"StatisPeriod" xml:"StatisPeriod"`
	NewProductRegulation   bool                                            `json:"NewProductRegulation" xml:"NewProductRegulation"`
	ItemConditionArray     string                                          `json:"ItemConditionArray" xml:"ItemConditionArray"`
	Status                 string                                          `json:"Status" xml:"Status"`
	RecallName             string                                          `json:"RecallName" xml:"RecallName"`
	ItemConditionExpress   string                                          `json:"ItemConditionExpress" xml:"ItemConditionExpress"`
	SplitParts             pairecservice.SplitPartsInGetTrafficControlTask `json:"SplitParts" xml:"SplitParts"`

	PlanTraffic    map[string]float64 `json:"plan_traffics"`
	TargetTraffics map[string]float64 `json:"target_traffics"`
}
