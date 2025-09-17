package api

import (
	"context"
	"encoding/json"
	"fmt"
	pairecservice20221213 "github.com/alibabacloud-go/pairecservice-20221213/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pairecservice"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
	"github.com/antihax/optional"
	jsoniter "github.com/json-iterator/go"
	"time"
)

// Linger please
var (
	_ context.Context
)

type TrafficControlApiService service

type TrafficControlApiListTrafficControlTasksOpts struct {
	Name                 optional.String
	TrafficControlTaskId optional.String
	SceneId              optional.Int32
	Env                  optional.String
	Status               optional.String
	Version              optional.String
	ControlTargetFilter  optional.String
	SortBy               optional.String
	Order                optional.String
	PageNumber           optional.Int32
	PageSize             optional.Int32
	ALL                  optional.Bool
}

func (fca *TrafficControlApiService) ListTrafficControlTasks(localVarOptionals *TrafficControlApiListTrafficControlTasksOpts, serviceName string) (ListTrafficControlTasksResponse, error) {
	listTrafficControlRequest := pairecservice.CreateListTrafficControlTasksRequest()
	listTrafficControlRequest.InstanceId = fca.instanceId
	listTrafficControlRequest.SetDomain(fca.client.GetDomain())

	if localVarOptionals.Env.Value() == common.Environment_Daily_Desc {
		listTrafficControlRequest.Environment = "Daily"
	} else if localVarOptionals.Env.Value() == common.Environment_Prepub_Desc {
		listTrafficControlRequest.Environment = "Pre"
	} else if localVarOptionals.Env.Value() == common.Environment_Product_Desc {
		listTrafficControlRequest.Environment = "Prod"
	}

	if localVarOptionals.Status.Value() == common.TrafficCtrlTask_NotRunning_Status {
		listTrafficControlRequest.Status = "NotRunning"
	} else if localVarOptionals.Status.Value() == common.TrafficCtrlTask_Ready_Status {
		listTrafficControlRequest.Status = "Ready"
	} else if localVarOptionals.Status.Value() == common.TrafficCtrlTask_Running_Status {
		listTrafficControlRequest.Status = "Running"
	} else if localVarOptionals.Status.Value() == common.TrafficCtrlTask_Finished_Status {
		listTrafficControlRequest.Status = "Finished"
	}

	if localVarOptionals.Version.Value() == common.Version_Latest {
		listTrafficControlRequest.Version = "Latest"
	} else if localVarOptionals.Version.Value() == common.Version_Released {
		listTrafficControlRequest.Version = "Released"
	}

	if localVarOptionals.ControlTargetFilter.Value() == common.ControlTargetFilter_All {
		listTrafficControlRequest.ControlTargetFilter = "All"
	} else if localVarOptionals.ControlTargetFilter.Value() == common.ControlTargetFilter_Vaild {
		listTrafficControlRequest.ControlTargetFilter = "Valid"
	} else if localVarOptionals.ControlTargetFilter.Value() == common.ControlTargetFilter_None {
		listTrafficControlRequest.ControlTargetFilter = "None"
	}

	listTrafficControlRequest.All = requests.NewBoolean(localVarOptionals.ALL.Value())

	var (
		localVarReturnValue     ListTrafficControlTasksResponse
		trafficControlTaskArray []model.TrafficControlTask
	)
	response, err := fca.client.ListTrafficControlTasks(listTrafficControlRequest)

	if err != nil || response == nil {
		return localVarReturnValue, err
	}

	if len(response.TrafficControlTasks) == 0 {
		return localVarReturnValue, err
	}

	for _, trafficControlTask := range response.TrafficControlTasks {

		if trafficControlTask.ServiceId != "" && serviceName != "" {
			getServiceRequest := &pairecservice20221213.GetServiceRequest{
				InstanceId: tea.String(fca.instanceId),
			}
			serviceResponse, err := fca.client.clientV2.GetService(tea.String(trafficControlTask.ServiceId), getServiceRequest)
			if err != nil {
				return localVarReturnValue, err
			}
			var taskServiceName string
			if localVarOptionals.Env.Value() == common.Environment_Prepub_Desc {
				taskServiceName = fmt.Sprintf("%s_%s", *serviceResponse.Body.Name, common.Environment_Prepub_Desc)
			} else {
				taskServiceName = *serviceResponse.Body.Name
			}
			if taskServiceName != serviceName {
				continue
			}
		}

		var task model.TrafficControlTask
		// List of storage traffic control tasks
		task.TrafficControlTaskId = trafficControlTask.TrafficControlTaskId
		task.Name = trafficControlTask.Name
		task.Description = trafficControlTask.Description
		task.SceneId = trafficControlTask.SceneId
		task.SceneName = trafficControlTask.SceneName
		task.ProductStatus = trafficControlTask.ProductStatus
		task.PrepubStatus = trafficControlTask.PrepubStatus

		task.ExecutionTime = trafficControlTask.ExecutionTime
		task.StartTime = trafficControlTask.StartTime
		task.EndTime = trafficControlTask.EndTime
		task.BehaviorTableMetaId = trafficControlTask.BehaviorTableMetaId
		task.UserTableMetaId = trafficControlTask.UserTableMetaId
		task.ItemTableMetaId = trafficControlTask.ItemTableMetaId
		task.UserConditionType = trafficControlTask.UserConditionType
		task.UserConditionArray = trafficControlTask.UserConditionArray
		task.UserConditionExpress = trafficControlTask.UserConditionExpress
		task.StatisBehaviorConditionType = trafficControlTask.StatisBehaviorConditionType
		task.StatisBehaviorConditionArray = trafficControlTask.StatisBehaviorConditionArray
		task.StatisBahaviorConditionExpress = trafficControlTask.StatisBahaviorConditionExpress
		task.ControlType = trafficControlTask.ControlType
		task.ControlGranularity = trafficControlTask.ControlGranularity
		task.ControlLogic = trafficControlTask.ControlLogic
		task.ItemConditionType = trafficControlTask.ItemConditionType
		task.ItemConditionArray = trafficControlTask.ItemConditionArray
		task.ItemConditionExpress = trafficControlTask.ItemConditionExpress

		task.GmtCreateTime = trafficControlTask.GmtCreateTime
		task.GmtModifiedTime = trafficControlTask.GmtModifiedTime

		var trafficControlTargets []model.TrafficControlTarget
		trafficControlTargetsMap := make(map[string]model.TrafficControlTarget)

		// List of storage traffic control targets

		behaviorTableMetaRequest := pairecservice.CreateGetTableMetaRequest()
		behaviorTableMetaRequest.TableMetaId = trafficControlTask.BehaviorTableMetaId
		behaviorTableMetaRequest.InstanceId = fca.instanceId
		behaviorTableMetaRequest.SetDomain(fca.client.GetDomain())

		behaviorTableMeta, err := fca.client.GetTableMeta(behaviorTableMetaRequest)

		if err != nil {
			return localVarReturnValue, nil
		}

		task.BehaviorTableMeta = &pairecservice.TableMetasItem{
			Name:            behaviorTableMeta.Name,
			ResourceId:      behaviorTableMeta.ResourceId,
			TableName:       behaviorTableMeta.TableName,
			Type:            behaviorTableMeta.Type,
			Description:     behaviorTableMeta.Description,
			Module:          behaviorTableMeta.Module,
			Url:             behaviorTableMeta.Url,
			GmtCreateTime:   behaviorTableMeta.GmtCreateTime,
			GmtModifiedTime: behaviorTableMeta.GmtModifiedTime,
			GmtImportedTime: behaviorTableMeta.GmtImportedTime,
			Config:          behaviorTableMeta.Config,
			Fields:          behaviorTableMeta.Fields,
		}

		userTableMetaRequest := pairecservice.CreateGetTableMetaRequest()
		userTableMetaRequest.TableMetaId = trafficControlTask.UserTableMetaId
		userTableMetaRequest.InstanceId = fca.instanceId
		userTableMetaRequest.SetDomain(fca.client.GetDomain())

		userTableMeta, err := fca.client.GetTableMeta(userTableMetaRequest)

		if err != nil {
			return localVarReturnValue, err
		}

		task.UserTableMeta = &pairecservice.TableMetasItem{
			Name:            userTableMeta.Name,
			ResourceId:      userTableMeta.ResourceId,
			TableName:       userTableMeta.TableName,
			Type:            userTableMeta.Type,
			Description:     userTableMeta.Description,
			Module:          userTableMeta.Module,
			Url:             userTableMeta.Url,
			GmtCreateTime:   userTableMeta.GmtCreateTime,
			GmtModifiedTime: userTableMeta.GmtModifiedTime,
			GmtImportedTime: userTableMeta.GmtImportedTime,
			Config:          userTableMeta.Config,
			Fields:          userTableMeta.Fields,
		}

		itemTableMetaRequest := pairecservice.CreateGetTableMetaRequest()
		itemTableMetaRequest.TableMetaId = trafficControlTask.ItemTableMetaId
		itemTableMetaRequest.InstanceId = fca.instanceId
		itemTableMetaRequest.SetDomain(fca.client.GetDomain())

		itemTableMeta, err := fca.client.GetTableMeta(itemTableMetaRequest)

		if err != nil {
			return localVarReturnValue, err
		}

		task.ItemTableMeta = &pairecservice.TableMetasItem{
			Name:            itemTableMeta.Name,
			ResourceId:      itemTableMeta.ResourceId,
			TableName:       itemTableMeta.TableName,
			Type:            itemTableMeta.Type,
			Description:     itemTableMeta.Description,
			Module:          itemTableMeta.Module,
			Url:             itemTableMeta.Url,
			GmtCreateTime:   itemTableMeta.GmtCreateTime,
			GmtModifiedTime: itemTableMeta.GmtModifiedTime,
			GmtImportedTime: itemTableMeta.GmtImportedTime,
			Config:          itemTableMeta.Config,
			Fields:          itemTableMeta.Fields,
		}

		for _, target := range trafficControlTask.TrafficControlTargets {
			var t model.TrafficControlTarget
			t.TrafficControlTaskId = trafficControlTask.TrafficControlTaskId
			t.TrafficControlTargetId = target.TrafficControlTargetId
			t.Name = target.Name
			t.StartTime = target.StartTime
			t.EndTime = target.EndTime
			t.ItemConditionType = target.ItemConditionType
			t.ItemConditionArray = target.ItemConditionArray
			t.ItemConditionExpress = target.ItemConditionExpress
			t.Event = target.Event
			t.Value = target.Value
			t.StatisPeriod = target.StatisPeriod
			t.ToleranceValue = target.ToleranceValue
			t.NewProductRegulation = target.NewProductRegulation
			t.RecallName = target.RecallName
			t.Status = target.Status
			t.SplitParts = target.SplitParts
			t.GmtCreateTime = target.GmtCreateTime
			t.GmtModifiedTime = target.GmtModifiedTime
			trafficControlTargetsMap[target.TrafficControlTargetId] = t

			//Obtain traffic details about a traffic control task
			trafficControlTaskTrafficRequest := pairecservice.CreateGetTrafficControlTaskTrafficRequest()
			trafficControlTaskTrafficRequest.TrafficControlTaskId = task.TrafficControlTaskId
			trafficControlTaskTrafficRequest.InstanceId = fca.instanceId
			trafficControlTaskTrafficRequest.Environment = listTrafficControlRequest.Environment
			trafficControlTaskTrafficRequest.SetDomain(fca.client.GetDomain())
			tResponse, err := fca.client.common.client.GetTrafficControlTaskTraffic(trafficControlTaskTrafficRequest)

			if err != nil {
				return localVarReturnValue, err
			}

			traffic := tResponse.TrafficControlTaskTrafficInfo

			for _, targetTraffic := range traffic.TargetTraffics {

				if len(tResponse.TrafficControlTaskTrafficInfo.TaskTraffics) == 0 || len(tResponse.TrafficControlTaskTrafficInfo.TargetTraffics) == 0 {
					continue
				}
				taskTraffics := make(map[string]float64)
				targetTraffics := make(map[string]float64)

				tempTarget, ok := trafficControlTargetsMap[targetTraffic.TrafficContorlTargetId]
				if !ok {
					continue
				}

				for targetId, trafficDetails := range traffic.TaskTraffics {
					taskTraffic := make(map[string]float64, 0)
					trafficDet, _ := json.Marshal(trafficDetails)
					_ = json.Unmarshal(trafficDet, &taskTraffic)
					taskTraffics[targetId] = taskTraffic["Traffic"]

				}
				tempTarget.TaskTraffics = taskTraffics
				//targetTraffic.Data 由 array -> map
				var recordTime int64
				for targetId, targetTrafficDetails := range targetTraffic.Data {
					targetTrafficNew := make(map[string]float64, 0)
					targetTrafficDet, _ := json.Marshal(targetTrafficDetails)
					_ = json.Unmarshal(targetTrafficDet, &targetTrafficNew)
					targetTraffics[targetId] = targetTrafficNew["Traffic"]
					recordTime = int64(targetTrafficNew["RecordTime"])
				}
				tempTarget.TargetTraffics = targetTraffics
				tempTarget.RecordTime = time.Unix(recordTime, 0)
				trafficControlTargetsMap[tempTarget.TrafficControlTargetId] = tempTarget
			}
		}

		for _, target := range trafficControlTargetsMap {
			trafficControlTargets = append(trafficControlTargets, target)
		}
		task.TrafficControlTargets = trafficControlTargets
		trafficControlTaskArray = append(trafficControlTaskArray, task)
	}

	localVarReturnValue.TrafficControlTasks = trafficControlTaskArray
	return localVarReturnValue, nil
}

type TaskTraffic struct {
	PlanTraffic float64 `json:"plan_traffic"`
}

type TargetTraffic struct {
	TargetTraffic float64 `json:"target_traffic"`
	RecordTime    int64   `json:"record_time"`
}

func (tct *TrafficControlApiService) SetTrafficControlTrafficFData(t model.TrafficControlTaskTrafficData) (string, error) {
	t.InstanceId = tct.instanceId
	req := pairecservice.CreateUpdateTrafficControlTaskTrafficRequest()
	body, _ := jsoniter.MarshalToString(t)

	req.TrafficControlTaskId = t.TrafficControlTaskId
	req.Body = body
	req.SetDomain(tct.client.GetDomain())

	response, err := tct.client.UpdateTrafficControlTaskTraffic(req)

	if err != nil {
		return "", err
	}
	return response.RequestId, nil

}
