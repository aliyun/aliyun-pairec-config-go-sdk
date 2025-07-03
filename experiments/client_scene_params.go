package experiments

import (
	"fmt"
	"strconv"
	"time"

	pairecservice20221213 "github.com/alibabacloud-go/pairecservice-20221213/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
)

// LoadSceneParamsData specifies a function to load param data from A/B Test Server
func (e *ExperimentClient) LoadSceneParamsData() {
	sceneParamData := make(map[string]model.SceneParams, 0)

	listScenesResponse, err := e.APIClient.SceneApi.ListAllScenes()
	if err != nil {
		e.logError(fmt.Errorf("list scenes error, err=%v", err))
		return
	}

	for _, scene := range listScenesResponse.Scenes {
		sceneParams := model.NewSceneParams()
		listParamsRequest := &pairecservice20221213.ListParamsRequest{}
		listParamsRequest.Environment = tea.String(common.EnvironmentDesc2OpenApiString[e.Environment])
		listParamsRequest.SceneId = tea.String(strconv.FormatInt(scene.SceneId, 10))
		listParamsRequest.Encrypted = tea.Bool(true)
		listParamsRequest.InstanceId = tea.String(e.InstanceId)

		paramResponse, err := e.APIClientV2.ListParams(listParamsRequest)

		if err != nil {
			e.logError(fmt.Errorf("list params error, err=%v", err))
			return
		}

		for _, param := range paramResponse.Body.Params {
			sceneParams.AddParam(*param.Name, *param.Value)
		}
		sceneParamData[scene.SceneName] = sceneParams
	}
	if len(sceneParamData) > 0 {
		e.sceneParamData = sceneParamData
	}
}

// loopLoadExperimentData async loop invoke LoadExperimentData function
func (e *ExperimentClient) loopLoadSceneParamsData() {

	for {
		time.Sleep(time.Minute)
		e.LoadSceneParamsData()
	}
}

func (e *ExperimentClient) GetSceneParams(sceneName string) model.SceneParams {
	sceneParams, ok := e.sceneParamData[sceneName]
	if !ok {
		return model.NewEmptySceneParams()
	}

	return sceneParams
}
