package experiments

import (
	"context"
	"fmt"
	"time"

	"github.com/antihax/optional"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/api"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/model"
)

func (e *ExperimentClient) logError(err error) {
	if e.ErrorLogger != nil {
		e.ErrorLogger.Printf(err.Error())
		return
	}

	if e.Logger != nil {
		e.Logger.Printf(err.Error())
	}
}

// LoadExperimentData specifies a function to load data from A/B Test Server
func (e *ExperimentClient) LoadExperimentData() {
	sceneData := make(map[string]*model.Scene, 0)

	listScenesResponse, err := e.APIClient.SceneApi.ListAllScenes(context.Background())
	if err != nil {
		e.logError(fmt.Errorf("list scenes error, err=%v", err))
		return
	}

	if listScenesResponse.Code != common.CODE_OK {
		e.logError(fmt.Errorf("list scenes error, requestid=%s,code=%s, msg=%s", listScenesResponse.RequestId, listScenesResponse.Code, listScenesResponse.Message))
		return
	}

	for _, scene := range listScenesResponse.Data["scenes"] {
		listExpRoomsResponse, err := e.APIClient.ExperimentRoomApi.ListExperimentRooms(context.Background(), e.Environment,
			&api.ExperimentRoomApiListExperimentRoomsOpts{SceneId: optional.NewInt64(scene.SceneId), Status: optional.NewUint32(common.ExpRoom_Status_Online)})

		if err != nil {
			e.logError(fmt.Errorf("list experiment rooms error, err=%v", err))
			continue
		}
		if listExpRoomsResponse.Code != common.CODE_OK {
			e.logError(fmt.Errorf("list experiment rooms error, requestid=%s,code=%s, msg=%s", listExpRoomsResponse.RequestId, listExpRoomsResponse.Code, listExpRoomsResponse.Message))
			continue
		}
		for _, experimentRoom := range listExpRoomsResponse.Data["experiment_rooms"] {
			if experimentRoom.DebugCrowdId != 0 {
				listCrowdUsersResponse, err := e.APIClient.CrowdApi.GetCrowdUsersById(context.Background(), experimentRoom.DebugCrowdId)
				if err != nil {
					e.logError(fmt.Errorf("list crowd users error, err=%v", err))
					continue
				}
				experimentRoom.DebugCrowdIdUsers = listCrowdUsersResponse.Data["users"]
			}
			// ExperimentRoom init
			if err := experimentRoom.Init(); err != nil {
				e.logError(fmt.Errorf("experiment room init error, err=%v", err))
				continue
			}

			scene.AddExperimentRoom(experimentRoom)
			listLayersResponse, err := e.APIClient.LayerApi.ListLayers(context.Background(), experimentRoom.ExpRoomId)
			if err != nil {
				e.logError(fmt.Errorf("list layers error, err=%v", err))
				continue
			}
			if listLayersResponse.Code != common.CODE_OK {
				e.logError(fmt.Errorf("list layers error, requestid=%s,code=%s, msg=%s", listLayersResponse.RequestId, listLayersResponse.Code, listLayersResponse.Message))
				continue
			}

			for _, layer := range listLayersResponse.Data["layers"] {
				experimentRoom.AddLayer(layer)

				listExperimentGroupResponse, err := e.APIClient.ExperimentGroupApi.ListExperimentGroups(context.Background(), layer.LayerId,
					&api.ExperimentGroupApiListExperimentGroupsOpts{Status: optional.NewUint32(common.ExpGroup_Status_Online)})
				if err != nil {
					e.logError(fmt.Errorf("list experiment groups error, err=%v", err))
					continue
				}
				if listExperimentGroupResponse.Code != common.CODE_OK {
					e.logError(fmt.Errorf("list experiment groups error, requestid=%s,code=%s, msg=%s", listExperimentGroupResponse.RequestId, listExperimentGroupResponse.Code, listExperimentGroupResponse.Message))
					continue
				}

				for _, experimentGroup := range listExperimentGroupResponse.Data["experiment_groups"] {
					if experimentGroup.CrowdId != 0 {
						listCrowdUsersResponse, err := e.APIClient.CrowdApi.GetCrowdUsersById(context.Background(), experimentGroup.CrowdId)
						if err != nil {
							e.logError(fmt.Errorf("list crowd users error, err=%v", err))
							continue
						}
						experimentGroup.CrowdUsers = listCrowdUsersResponse.Data["users"]
					}

					if experimentGroup.DebugCrowdId != 0 {
						listCrowdUsersResponse, err := e.APIClient.CrowdApi.GetCrowdUsersById(context.Background(), experimentGroup.DebugCrowdId)
						if err != nil {
							e.logError(fmt.Errorf("list crowd users error, err=%v", err))
							continue
						}
						experimentGroup.DebugCrowdUsers = listCrowdUsersResponse.Data["users"]
					}

					// ExperimentGroup init
					if err := experimentGroup.Init(); err != nil {
						e.logError(fmt.Errorf("experiment group init error, err=%v", err))
						continue
					}

					layer.AddExperimentGroup(experimentGroup)

					listExperimentsResponse, err := e.APIClient.ExperimentApi.ListExperiments(context.Background(), experimentGroup.ExpGroupId,
						&api.ExperimentApiListExperimentsOpts{Status: optional.NewUint32(common.Experiment_Status_Online)})
					if err != nil {
						e.logError(fmt.Errorf("list experiments  error, err=%v", err))
						continue
					}
					if listExperimentsResponse.Code != common.CODE_OK {
						e.logError(fmt.Errorf("list experiments error, requestid=%s,code=%s, msg=%s", listExperimentsResponse.RequestId, listExperimentsResponse.Code, listExperimentsResponse.Message))
						continue
					}

					for _, experiment := range listExperimentsResponse.Data["experiments"] {
						if experiment.DebugCrowdId != 0 {
							listCrowdUsersResponse, err := e.APIClient.CrowdApi.GetCrowdUsersById(context.Background(), experiment.DebugCrowdId)
							if err != nil {
								e.logError(fmt.Errorf("list crowd users error, err=%v", err))
								continue
							}
							experiment.DebugCrowdUsers = listCrowdUsersResponse.Data["users"]
						}
						if err := experiment.Init(); err != nil {
							e.logError(fmt.Errorf("experiment init error, err=%v", err))
							continue
						}
						experimentGroup.AddExperiment(experiment)
					}
				}
			}
		}
		sceneData[scene.SceneName] = scene
	}
	if len(sceneData) > 0 {
		e.sceneMap = sceneData
	}
}

// loopLoadExperimentData async loop invoke LoadExperimentData function
func (e *ExperimentClient) loopLoadExperimentData() {

	for {
		time.Sleep(time.Minute)
		e.LoadExperimentData()
	}
}
