package experiments

import (
	"fmt"
	pairecv2 "github.com/alibabacloud-go/pairecservice-20221213/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/api"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	serviceName string
)

func init() {
	value := os.Getenv("SERVICE_NAME")
	valueArr := strings.Split(value, "@")
	if len(valueArr) == 2 {
		serviceName = valueArr[0]
	}
}
func (e *ExperimentClient) loadTrafficControlTasks() {
	//Load traffic control data for the production environment
	productTrafficControlTasks := make([]*model.TrafficControlTask, 0)
	prodQueryParams := &ListTrafficControlTasksQueryParams{
		ALL:                 true,
		ControlTargetFilter: "Valid",
		Env:                 "Prod",
		Status:              "Running",
		Version:             "Released",
	}
	prodResponse, err := e.listTrafficControlTasks(prodQueryParams)
	if err != nil {
		e.logError(fmt.Errorf("list traffic control tasks error, err=%v", err))
		return
	}

	for _, task := range prodResponse.TrafficControlTasks {
		productTrafficControlTasks = append(productTrafficControlTasks, task)
	}

	prepubTrafficControlTasks := make([]*model.TrafficControlTask, 0)
	preQueryParams := &ListTrafficControlTasksQueryParams{
		ALL:                 true,
		ControlTargetFilter: "Valid",
		Env:                 "Pre",
		Status:              "Running",
		Version:             "Released",
	}
	prePubResponse, _ := e.listTrafficControlTasks(preQueryParams)
	if err != nil {
		e.logError(fmt.Errorf("list traffic control tasks error,error=%v", err))
		return
	}

	for _, task := range prePubResponse.TrafficControlTasks {
		prepubTrafficControlTasks = append(prepubTrafficControlTasks, task)
	}

	e.productTrafficControlTasks = productTrafficControlTasks
	e.prepubTrafficControlTasks = prepubTrafficControlTasks

}

func (e *ExperimentClient) loopLoadTrafficControlTasks() {

	for {
		time.Sleep(time.Second * 30)
		e.loadTrafficControlTasks()
	}
}

func (e *ExperimentClient) ListTrafficControlTasks(env string) []*model.TrafficControlTask {
	currentTimestamp := time.Now().Unix()

	validTrafficControlTasks := make([]*model.TrafficControlTask, 0)

	var allTrafficControlTasks = make([]*model.TrafficControlTask, 0)

	if env == common.Environment_Prepub_Desc {
		allTrafficControlTasks = e.prepubTrafficControlTasks
	} else if env == common.Environment_Product_Desc {
		allTrafficControlTasks = e.productTrafficControlTasks
	} else {
		return validTrafficControlTasks
	}

	// filter valid traffic control task
	for _, task := range allTrafficControlTasks {
		if task.ExecutionTime != "Permanent" {
			startTime, _ := time.Parse(time.RFC3339, task.StartTime)
			endTime, _ := time.Parse(time.RFC3339, task.EndTime)

			if env == common.Environment_Product_Desc {
				if task.ProductStatus == common.TrafficCtrlTask_Running_Status && startTime.Unix() <= currentTimestamp && currentTimestamp < endTime.Unix() {
					validTrafficControlTasks = append(validTrafficControlTasks, task)
				} else {
					continue
				}
			} else if env == common.Environment_Prepub_Desc {
				if task.PrepubStatus == common.TrafficCtrlTask_Running_Status && startTime.Unix() <= currentTimestamp && currentTimestamp < endTime.Unix() {
					validTrafficControlTasks = append(validTrafficControlTasks, task)
				} else {
					continue
				}

			}
		} else {
			if env == common.Environment_Product_Desc {
				if task.ProductStatus == common.TrafficCtrlTask_Running_Status {
					validTrafficControlTasks = append(validTrafficControlTasks, task)
				} else {
					continue
				}
			} else if env == common.Environment_Prepub_Desc {
				if task.PrepubStatus == common.TrafficCtrlTask_Running_Status {
					validTrafficControlTasks = append(validTrafficControlTasks, task)
				} else {
					continue
				}
			}
		}
	}

	// filter valid traffic control target
	for _, task := range validTrafficControlTasks {
		validTargets := make([]*model.TrafficControlTarget, 0)
		for _, target := range task.TrafficControlTargets {
			if e.isValidTrafficControlTarget(target) {
				validTargets = append(validTargets, target)
			}
		}
		task.TrafficControlTargets = validTargets
	}

	return validTrafficControlTasks
}

func (e *ExperimentClient) isValidTrafficControlTarget(target *model.TrafficControlTarget) bool {
	currentTimestamp := time.Now().Unix()

	startTime, _ := time.Parse(time.RFC3339, target.StartTime)
	endTime, _ := time.Parse(time.RFC3339, target.EndTime)

	// 不在时间范围内
	if startTime.Unix() > currentTimestamp || currentTimestamp >= endTime.Unix() {
		return false
	}

	// 状态过滤
	if target.Status == common.TrafficControlTargetStatusOfClosed {
		return false
	}

	if target.TrafficControlTargetId == "" {
		return false
	}

	return true
}

type TrafficControlTargetTraffic struct {
	ItemOrExpId            string    `json:"item_or_exp_id"`
	TrafficControlTaskId   string    `json:"traffic_control_task_id"`
	TrafficControlTargetId string    `json:"traffic_control_target_id"`
	TargetTraffic          float64   `json:"target_traffic"`
	TaskTraffic            float64   `json:"task_traffic"`
	RecordTime             time.Time `json:"record_time"`
}

func (e *ExperimentClient) GetTrafficControlActualTraffic(env string, expIdOrItemIdList ...string) map[string][]*TrafficControlTargetTraffic {
	tasks := e.ListTrafficControlTasks(env)

	resultTrafficsMap := make(map[string][]*TrafficControlTargetTraffic, 0) // key: target_id

	if len(expIdOrItemIdList) == 0 {
		return resultTrafficsMap
	}

	tmpActualTrafficMap := make(map[string]*TrafficControlTargetTraffic, 0)

	for _, task := range tasks {
		taskTrafficMap := task.ActualTraffic.TaskTraffics
		for _, target := range task.TrafficControlTargets {
			targetTraffic, ok := getTargetTraffic(target.TrafficControlTargetId, task.ActualTraffic.TargetTraffics)
			if ok {
				for expOrItemId, targetTrafficDetail := range targetTraffic.Data {

					key := fmt.Sprintf("%s@%s", target.TrafficControlTargetId, expOrItemId)
					recordTime := time.Unix(targetTrafficDetail.RecordTime, 0)

					taskTraffic := taskTrafficMap[expOrItemId]
					tmpTrafficInfo := &TrafficControlTargetTraffic{
						ItemOrExpId:            expOrItemId,
						TrafficControlTaskId:   task.TrafficControlTaskId,
						TrafficControlTargetId: target.TrafficControlTargetId,
						TargetTraffic:          targetTrafficDetail.Traffic,
						RecordTime:             recordTime,
						TaskTraffic:            taskTraffic.Traffic,
					}
					tmpActualTrafficMap[key] = tmpTrafficInfo
				}
			}
		}
	}

	var traffics []*TrafficControlTargetTraffic

	// filter valid traffic control target
	for _, actualTraffic := range tmpActualTrafficMap {
		ok := isItemInArray(actualTraffic.ItemOrExpId, expIdOrItemIdList)
		if ok {
			traffics = append(traffics, actualTraffic)
		}
	}

	for _, traffic := range traffics {
		_, ok := resultTrafficsMap[traffic.TrafficControlTargetId]
		if ok {
			resultTrafficsMap[traffic.TrafficControlTargetId] = append(resultTrafficsMap[traffic.TrafficControlTargetId], traffic)
		} else {
			resultTrafficsMap[traffic.TrafficControlTargetId] = []*TrafficControlTargetTraffic{traffic}
		}
	}

	return resultTrafficsMap
}

func getTargetTraffic(targetId string, targetTraffics []*model.TargetTraffic) (*model.TargetTraffic, bool) {
	for _, targetTraffic := range targetTraffics {
		if targetTraffic.TrafficControlTargetId == targetId {
			return targetTraffic, true
		}
	}
	return nil, false
}

type ListTrafficControlTasksQueryParams struct {
	Name                 string
	TrafficControlTaskId string
	SceneId              int32
	Env                  string
	Status               string
	Version              string
	ControlTargetFilter  string
	SortBy               string
	Order                string
	PageNumber           string
	PageSize             string
	ALL                  bool
}

// 调用 openAPI 获取 task 以及每个 task 的 traffic
func (e *ExperimentClient) listTrafficControlTasks(params *ListTrafficControlTasksQueryParams) (api.ListTrafficControlTasksResponse, error) {
	listTrafficControlTasksRequest := &pairecv2.ListTrafficControlTasksRequest{}
	listTrafficControlTasksRequest.InstanceId = tea.String(e.InstanceId)
	listTrafficControlTasksRequest.Environment = tea.String(params.Env)
	listTrafficControlTasksRequest.Status = tea.String(params.Status)
	listTrafficControlTasksRequest.Version = tea.String(params.Version)
	listTrafficControlTasksRequest.ControlTargetFilter = tea.String(params.ControlTargetFilter)
	listTrafficControlTasksRequest.All = tea.Bool(params.ALL)

	localVarReturnValue := api.ListTrafficControlTasksResponse{}
	response, err := e.APIClientV2.ListTrafficControlTasks(listTrafficControlTasksRequest)

	if err != nil {
		return localVarReturnValue, err
	}

	for _, trafficControlTask := range response.Body.TrafficControlTasks {

		// filter by service name
		if serviceName != "" {
			var find bool

			// 兼容旧数据
			if trafficControlTask.ServiceId != nil && *trafficControlTask.ServiceId != "" {
				getServiceRequest := &pairecv2.GetServiceRequest{
					InstanceId: tea.String(e.InstanceId),
				}
				serviceResponse, err := e.APIClientV2.GetService(trafficControlTask.ServiceId, getServiceRequest)
				if err != nil {
					return localVarReturnValue, err
				}
				var taskServiceName string
				if params.Env == common.OpenAPIEnvironmentPrepub {
					taskServiceName = fmt.Sprintf("%s_%s", *serviceResponse.Body.Name, common.Environment_Prepub_Desc)
				} else {
					taskServiceName = *serviceResponse.Body.Name
				}
				if taskServiceName == serviceName {
					find = true
				}
			}

			for _, serviceId := range trafficControlTask.ServiceIdList {
				getServiceRequest := &pairecv2.GetServiceRequest{}
				getServiceRequest.InstanceId = tea.String(e.InstanceId)
				serviceResponse, err := e.APIClientV2.GetService(tea.String(strconv.Itoa(int(*serviceId))), getServiceRequest)
				if err != nil {
					return localVarReturnValue, err
				}
				var taskServiceName string
				if params.Env == common.OpenAPIEnvironmentPrepub {
					taskServiceName = fmt.Sprintf("%s_%s", *serviceResponse.Body.Name, common.Environment_Prepub_Desc)
				} else {
					taskServiceName = *serviceResponse.Body.Name
				}
				if taskServiceName == serviceName {
					find = true
					break
				}
			}
			if !find {
				continue
			}
		}
		task := model.TrafficControlTaskConvert(trafficControlTask)

		for _, trafficControlTarget := range trafficControlTask.TrafficControlTargets {

			target := model.TrafficControlTargetConvert(trafficControlTarget)

			task.TrafficControlTargets = append(task.TrafficControlTargets, target)
		}
		// 获取每个 task 的实际流量
		getTrafficRequest := &pairecv2.GetTrafficControlTaskTrafficRequest{}
		getTrafficRequest.InstanceId = tea.String(e.InstanceId)
		getTrafficRequest.Environment = listTrafficControlTasksRequest.Environment
		trafficResponse, err := e.APIClientV2.GetTrafficControlTaskTraffic(tea.String(task.TrafficControlTaskId), getTrafficRequest)
		if err != nil {
			return localVarReturnValue, err
		}
		actualTraffic := model.ActualTrafficConvert(trafficResponse.Body.TrafficControlTaskTrafficInfo)

		task.ActualTraffic = actualTraffic

		localVarReturnValue.TrafficControlTasks = append(localVarReturnValue.TrafficControlTasks, task)

	}

	return localVarReturnValue, nil
}

func isItemInArray(element string, array []string) bool {
	for _, value := range array {
		if element == value {
			return true
		}
	}
	return false
}
