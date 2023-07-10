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

// LoadSceneParamsData specifies a function to load param data from A/B Test Server
func (e *ExperimentClient) LoadSceneParamsData() {
	sceneParamData := make(map[string]model.SceneParams, 0)

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
		sceneParams := model.NewSceneParams()
		listParamsResponse, err := e.APIClient.ParamApi.GetParam(context.Background(), scene.SceneId,
			&api.ParamApiGetParamOpts{Environment: optional.NewString(e.Environment)})

		if err != nil {
			e.logError(fmt.Errorf("list params error, err=%v", err))
			continue
		}
		if listParamsResponse.Code != common.CODE_OK {
			e.logError(fmt.Errorf("list params error, requestid=%s,code=%s, msg=%s", listParamsResponse.RequestId, listParamsResponse.Code, listParamsResponse.Message))
			continue
		}
		for _, param := range listParamsResponse.Data["params"] {
			sceneParams.AddParam(param.ParamName, param.ParamValue)
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
