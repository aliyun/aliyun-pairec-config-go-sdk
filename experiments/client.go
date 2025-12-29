package experiments

import (
	"crypto/md5"
	"errors"
	"fmt"
	"hash/fnv"
	"strconv"
	"time"

	"github.com/alibabacloud-go/tea/tea"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	pairecservice20221213 "github.com/alibabacloud-go/pairecservice-20221213/v3/client"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/api"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
	"github.com/aliyun/credentials-go/credentials"
)

//var PAIRecEndpoint = "pairecservice-vpc.%s.aliyuncs.com"

type ClientOption func(c *ExperimentClient)

func WithLogger(l Logger) ClientOption {
	return func(e *ExperimentClient) {
		e.Logger = l
	}
}

func WithErrorLogger(l Logger) ClientOption {
	return func(e *ExperimentClient) {
		e.ErrorLogger = l
	}
}

func WithDomain(domian string) ClientOption {
	return func(e *ExperimentClient) {
		e.APIClient.SetDomain(domian)
	}
}

type ExperimentClient struct {
	// Environment control the sdk shoud get which environment data .
	// Valid value is daily, prepub,product
	Environment string

	// APIClient invoke api to connect to pairecservice open api
	APIClient *api.APIClient

	// open api v2
	APIClientV2 *pairecservice20221213.Client

	InstanceId string

	//
	SceneMap map[string]*model.Scene

	// sceneParamData map of parameters of scene name
	sceneParamData map[string]model.SceneParams

	// sceneFlowCtrlPlanData map of flow ctrl plan of scene name
	productTrafficControlTasks []*model.TrafficControlTask

	// prepubSceneFlowCtrlPlanData map of flow ctrl plan of scene name (prepub env)
	prepubTrafficControlTasks []*model.TrafficControlTask

	// Logger specifies a logger used to report internal changes within the writer
	Logger Logger

	// ErrorLogger is the logger to report errors
	ErrorLogger Logger
}

func NewExperimentClient(instanceId, regionId, accessKeyId, accessKeySecret, environment string, opts ...ClientOption) (*ExperimentClient, error) {
	client := ExperimentClient{
		Environment: environment,
		SceneMap:    make(map[string]*model.Scene, 0),
		InstanceId:  instanceId,
	}

	var err error
	client.APIClient, err = api.NewAPIClient(instanceId, regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(&client)
	}

	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	//endpoint := fmt.Sprintf(nAIRecEndpoint, regionId)
	config.Endpoint = tea.String(client.APIClient.GetDomain())

	var (
		clientV2 *pairecservice20221213.Client
	)
	if accessKeyId == "" || accessKeySecret == "" {
		credential, err1 := credentials.NewCredential(nil)
		if err1 != nil {
			return nil, err1
		}

		config.Credential = credential
		clientV2, err = pairecservice20221213.NewClient(config)
	} else {
		clientV2, err = pairecservice20221213.NewClient(config)
	}

	if err != nil {
		return nil, err
	}
	client.APIClientV2 = clientV2

	if err := client.Validate(); err != nil {
		return nil, err
	}

	client.LoadExperimentData()
	client.LoadSceneParamsData()

	go client.loopLoadExperimentData()
	go client.loopLoadSceneParamsData()

	client.LoadTrafficControlTasks()
	go client.LoopLoadTrafficControlTasks()

	return &client, nil
}

// Validate check the  ExperimentClient value
func (e *ExperimentClient) Validate() error {
	if e.Environment == "" {
		return errors.New("environment is empty")
	}
	if err := common.CheckEnvironmentValue(e.Environment); err != nil {
		return err
	}
	return nil
}

// MatchExperiment specifies to find match experiment by the ExperimentContext
// If not find the scene return error or return ExperimentResult
func (e *ExperimentClient) MatchExperiment(sceneName string, experimentContext *model.ExperimentContext) *model.ExperimentResult {
	sceneData := e.SceneMap

	scene, exist := sceneData[sceneName]

	if !exist {
		if sceneName != model.GlobalSceneName {
			e.logError(fmt.Errorf("scene:%s, not found the scene info", sceneName))
		}
		return model.NewExperimentResult(sceneName, experimentContext)
	}

	experimentResult := model.NewExperimentResult(sceneName, experimentContext)

	if sceneName != model.GlobalSceneName {
		experimentResult.GlobalSceneExperimentResult = e.MatchExperiment(model.GlobalSceneName, experimentContext)
	}

	var defaultExperimentRoom *model.ExperimentRoom
	var matchExperimentRoom *model.ExperimentRoom
	// first find base experiment room
	for _, experimentRoom := range scene.ExperimentRooms {
		if experimentRoom.Type == common.ExpRoom_Type_Base {
			defaultExperimentRoom = experimentRoom
			break
		}
	}
	// if experiment room has debug users then matchExperimentRoom
	for _, experimentRoom := range scene.ExperimentRooms {
		if experimentRoom.MatchDebugUsers(experimentContext) {
			matchExperimentRoom = experimentRoom
			break
		}
	}
	// if matchExperimentRoom is null, so no debug users found
	// then find no base experiment room is match
	if matchExperimentRoom == nil {
		for _, experimentRoom := range scene.ExperimentRooms {
			if experimentRoom.Type != common.ExpRoom_Type_Base {
				if experimentRoom.Match(experimentContext) {
					matchExperimentRoom = experimentRoom
					break
				}
			}
		}
	}

	if matchExperimentRoom == nil {
		matchExperimentRoom = defaultExperimentRoom
	}

	if matchExperimentRoom != nil {
		experimentResult.ExperimentRoom = matchExperimentRoom
		experimentResult.Layers = matchExperimentRoom.Layers

		for _, layer := range matchExperimentRoom.Layers {
			experimentGroup := layer.FindMatchExperimentGroup(experimentContext)
			if experimentGroup != nil {
				experimentResult.AddMatchExperimentGroup(layer.LayerName, experimentGroup)
				var defaultExperiment *model.Experiment
				var matchExperiment *model.Experiment
				for _, experiment := range experimentGroup.Experiments {
					if experiment.Type == common.Experiment_Type_Default {
						defaultExperiment = experiment
					}

				}
				// find match experiment
				if matchExperiment == nil {
					// first match debug users
					for _, experiment := range experimentGroup.Experiments {
						if experiment.Type != common.Experiment_Type_Default && experiment.MatchDebugUsers(experimentContext) {
							matchExperiment = experiment
							e.logInfo("match experiment debug users uid:%s", experimentContext.Uid)
							break
						}
					}
					if matchExperiment == nil {
						var hashKey string
						if experimentGroup.DistributionType == common.ExpGroup_Distribution_Type_TimeDuration {
							currTime := time.Now()
							duration := (currTime.Unix() % 86400) / int64((experimentGroup.DistributionTimeDuration * 60))
							hashKey = fmt.Sprintf("%s_%d_EXPROOM%d_LAYER%d_EXPGROUP%d", currTime.Format("20060102"), duration, experimentGroup.ExpRoomId,
								experimentGroup.LayerId, experimentGroup.ExpGroupId)

						} else {
							hashKey = fmt.Sprintf("%s_EXPROOM%d_LAYER%d_EXPGROUP%d", experimentContext.Uid, experimentGroup.ExpRoomId,
								experimentGroup.LayerId, experimentGroup.ExpGroupId)
						}
						hashValue := e.hashValue(hashKey)
						//e.logInfo("match experiment hash key:%s, value:%d", hashKey, hashValue)
						hashValueStr := strconv.FormatUint(hashValue, 10)
						experimentContext.SetExperimentHashString(hashValueStr)
						for _, experiment := range experimentGroup.Experiments {
							if experiment.Type != common.Experiment_Type_Default && experiment.Match(experimentContext) {
								matchExperiment = experiment
								break
							}
						}
					}

				}

				if matchExperiment == nil {
					// if defaultExperiment not found ,set baseExperiment is  defaultExperiment
					if defaultExperiment == nil {
						for _, experiment := range experimentGroup.Experiments {
							if experiment.Type == common.Experiment_Type_Base {
								defaultExperiment = experiment.Clone()
								defaultExperiment.Type = common.Experiment_Type_Default
							}
						}
					}

					matchExperiment = defaultExperiment
				}

				if matchExperiment != nil {
					experimentResult.AddMatchExperiment(layer.LayerName, matchExperiment)
				}
			}
		}

	}

	experimentResult.Init()
	return experimentResult
}

func (e *ExperimentClient) hashValue(hashKey string) uint64 {
	md5 := md5.Sum([]byte(hashKey))
	hash := fnv.New64()
	hash.Write(md5[:])

	return hash.Sum64()
}
func (e *ExperimentClient) logInfo(msg string, args ...interface{}) {
	if e.Logger != nil {
		e.Logger.Printf(msg, args...)
	}
}

func (e *ExperimentClient) ListTrafficControlTasks(env string) []*model.TrafficControlTask {
	if env == common.Environment_Prepub_Desc {
		return e.prepubTrafficControlTasks
	} else {
		return e.productTrafficControlTasks
	}
}
