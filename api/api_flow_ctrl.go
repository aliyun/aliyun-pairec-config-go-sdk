package api

import (
	"context"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pairecservice"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
	"github.com/antihax/optional"
	"strconv"
)

// Linger please
var (
	_ context.Context
)

type FlowCtrlApiService service

/*
FlowCtrlApiService 获取流控计划列表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *FlowCtrlApiListFlowCtrlPlansOpts - Optional Parameters:
     * @param "SceneId" (optional.Int32) -
     * @param "Status" (optional.String) -
@return ListFlowCtrlPlansResponse
*/

type FlowCtrlApiListFlowCtrlPlansOpts struct {
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

func (fca *FlowCtrlApiService) ListFlowCtrlPlans(localVarOptionals *FlowCtrlApiListFlowCtrlPlansOpts) (ListFlowCtrlPlansResponse, error) {
	listFlowCtrlRequest := pairecservice.CreateListTrafficControlTasksRequest()
	//listFlowCtrlRequest.
	listFlowCtrlRequest.InstanceId = fca.instanceId

	if localVarOptionals.Env.Value() == common.Environment_Daily_Desc {
		listFlowCtrlRequest.Environment = "Daily"
	} else if localVarOptionals.Env.Value() == common.Environment_Prepub_Desc {
		listFlowCtrlRequest.Environment = "Pre"
	} else if localVarOptionals.Env.Value() == common.Environment_Product_Desc {
		listFlowCtrlRequest.Environment = "Prod"
	}

	if localVarOptionals.Status.Value() == common.FlowCtrlPlan_NotRunning_Status {
		listFlowCtrlRequest.Status = "NotRunning"
	} else if localVarOptionals.Status.Value() == common.FlowCtrlPlan_Ready_Status {
		listFlowCtrlRequest.Status = "Ready"
	} else if localVarOptionals.Status.Value() == common.FlowCtrlPlan_Running_Status {
		listFlowCtrlRequest.Status = "Running"
	} else if localVarOptionals.Status.Value() == common.FlowCtrlPlan_Finished_Status {
		listFlowCtrlRequest.Status = "Finished"
	}

	if localVarOptionals.Version.Value() == common.Version_Latest {
		listFlowCtrlRequest.Version = "Latest"
	} else if localVarOptionals.Version.Value() == common.Version_Released {
		listFlowCtrlRequest.Version = "Released"
	}

	if localVarOptionals.ControlTargetFilter.Value() == common.ControlTargetFilter_All {
		listFlowCtrlRequest.ControlTargetFilter = "All"
	} else if localVarOptionals.ControlTargetFilter.Value() == common.ControlTargetFilter_Vaild {
		listFlowCtrlRequest.ControlTargetFilter = "Valid"
	} else if localVarOptionals.ControlTargetFilter.Value() == common.ControlTargetFilter_None {
		listFlowCtrlRequest.ControlTargetFilter = "None"
	}

	listFlowCtrlRequest.All = requests.NewBoolean(localVarOptionals.ALL.Value())

	var (
		localVarReturnValue ListFlowCtrlPlansResponse
		flowCtrlPlanArray   []model.TrafficControlTasksItem
	)
	response, err := fca.client.ListTrafficControlTasks(listFlowCtrlRequest)

	if err != nil || response == nil {
		return localVarReturnValue, err
	}

	if len(response.TrafficControlTasks) == 0 {
		return localVarReturnValue, err
	}

	for tIndex, trafficControlTask := range response.TrafficControlTasks {
		var task model.TrafficControlTasksItem
		//存储流量调控任务列表
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

		if task.TrafficControlTargets == nil {
			continue
		}
		var targets []model.TrafficControlTargetsItem
		//存储流量调控目标列表
		for index, target := range trafficControlTask.TrafficControlTargets {
			var t model.TrafficControlTargetsItem
			t.TrafficControlTaskId = trafficControlTask.TrafficControlTaskId
			t.TrafficControlTargetId = target.TrafficControlTargetId
			t.Name = target.Name
			t.StartTime = trafficControlTask.StartTime
			t.EndTime = trafficControlTask.EndTime
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
			targets[index] = t
		}

		task.TrafficControlTargets = targets

		getTableMetaRequest := pairecservice.CreateGetTableMetaRequest()
		getTableMetaRequest.TableMetaId = trafficControlTask.BehaviorTableMetaId
		getTableMetaRequest.InstanceId = fca.instanceId

		bMeta, err := fca.client.GetTableMeta(getTableMetaRequest)

		if err == nil {
			task.BehaviorTableMeta = &pairecservice.TableMetasItem{
				Name:            bMeta.Name,
				ResourceId:      bMeta.ResourceId,
				TableName:       bMeta.TableName,
				Type:            bMeta.Type,
				Description:     bMeta.Description,
				Module:          bMeta.Module,
				Url:             bMeta.Url,
				GmtCreateTime:   bMeta.GmtCreateTime,
				GmtModifiedTime: bMeta.GmtModifiedTime,
				GmtImportedTime: bMeta.GmtImportedTime,
				Config:          bMeta.Config,
				Fields:          bMeta.Fields,
			}
		}

		uGetTableMetaRequest := pairecservice.CreateGetTableMetaRequest()
		getTableMetaRequest.TableMetaId = trafficControlTask.UserTableMetaId
		getTableMetaRequest.InstanceId = fca.instanceId

		uMeta, err := fca.client.GetTableMeta(uGetTableMetaRequest)

		if err == nil {
			task.UserTableMeta = &pairecservice.TableMetasItem{
				Name:            uMeta.Name,
				ResourceId:      uMeta.ResourceId,
				TableName:       uMeta.TableName,
				Type:            uMeta.Type,
				Description:     uMeta.Description,
				Module:          uMeta.Module,
				Url:             uMeta.Url,
				GmtCreateTime:   uMeta.GmtCreateTime,
				GmtModifiedTime: uMeta.GmtModifiedTime,
				GmtImportedTime: uMeta.GmtImportedTime,
				Config:          uMeta.Config,
				Fields:          uMeta.Fields,
			}
		}

		iGetTableMetaRequest := pairecservice.CreateGetTableMetaRequest()
		getTableMetaRequest.TableMetaId = trafficControlTask.ItemTableMetaId
		getTableMetaRequest.InstanceId = fca.instanceId

		iMeta, err := fca.client.GetTableMeta(iGetTableMetaRequest)

		if err == nil {
			task.ItemTableMeta = &pairecservice.TableMetasItem{
				Name:            iMeta.Name,
				ResourceId:      iMeta.ResourceId,
				TableName:       iMeta.TableName,
				Type:            iMeta.Type,
				Description:     iMeta.Description,
				Module:          iMeta.Module,
				Url:             iMeta.Url,
				GmtCreateTime:   iMeta.GmtCreateTime,
				GmtModifiedTime: iMeta.GmtModifiedTime,
				GmtImportedTime: iMeta.GmtImportedTime,
				Config:          iMeta.Config,
				Fields:          iMeta.Fields,
			}
		}
		//Obtain traffic details about a traffic control task
		for _, task := range response.TrafficControlTasks {
			for _, targetTask := range task.TrafficControlTargets {

				trafficControlTaskTrafficRequest := pairecservice.CreateGetTrafficControlTaskTrafficRequest()
				trafficControlTaskTrafficRequest.TrafficControlTaskId = targetTask.TrafficControlTargetId
				trafficControlTaskTrafficRequest.InstanceId = fca.instanceId
				trafficControlTaskTrafficRequest.Environment = listFlowCtrlRequest.Environment
				tResponse, err := fca.client.common.client.GetTrafficControlTaskTraffic(trafficControlTaskTrafficRequest)
				if tResponse == nil || err != nil {
					continue
				}

				traffic := tResponse.TrafficControlTaskTrafficInfo
				for _, targetTraffic := range traffic.TargetTraffics {
					tid, err := strconv.Atoi(targetTraffic.TrafficContorlTargetId)
					if err == nil {
						toSetTraffic := targets[tid]
						for k, v := range traffic.TaskTraffics {
							toSetTraffic.PlanTraffic[k] = v.(float64)
						}
						for k, v := range targetTraffic.Data[0] {
							toSetTraffic.TargetTraffics[k] = v.(float64)
						}

					}
				}
			}

		}
		flowCtrlPlanArray[tIndex] = task
	}

	localVarReturnValue.Data.TrafficControlTasks = flowCtrlPlanArray
	return localVarReturnValue, nil
}
