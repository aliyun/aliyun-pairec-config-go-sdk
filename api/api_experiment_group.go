package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/model"
)

// Linger please
var (
	_ context.Context
)

type ExperimentGroupApiService service

/*
ExperimentGroupApiService Create a new experiment group
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body ExperimentGroup object that needs to be added

@return InlineResponse2008
*/
func (a *ExperimentGroupApiService) AddExperimentGroup(ctx context.Context, body model.ExperimentGroup) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/experiment_groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
		return localVarReturnValue, errors.New(fmt.Sprintf("Http Status code:%d", localVarHttpResponse.StatusCode))
	} else {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
	}
	return localVarReturnValue, nil
}

/*
ExperimentGroupApiService Delete ExperimentGroup By exp_group_id
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param expGroupId ExperimentGroup Id to delete

@return ModelApiResponse
*/
func (a *ExperimentGroupApiService) DeleteExperimentGroupById(ctx context.Context, expGroupId int64) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/experiment_groups/{exp_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"exp_group_id"+"}", fmt.Sprintf("%v", expGroupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
		return localVarReturnValue, errors.New(fmt.Sprintf("Http Status code:%d", localVarHttpResponse.StatusCode))
	} else {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
	}
	return localVarReturnValue, nil
}

/*
ExperimentGroupApiService Get ExperimentGroup By exp_group_id
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param expGroupId Experiment Id to get experiment data

@return Response
*/
func (a *ExperimentGroupApiService) GetExperimentGroupById(ctx context.Context, expGroupId int64) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/experiment_groups/{exp_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"exp_group_id"+"}", fmt.Sprintf("%v", expGroupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
		return localVarReturnValue, errors.New(fmt.Sprintf("Http Status code:%d", localVarHttpResponse.StatusCode))
	} else {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
	}
	return localVarReturnValue, nil
}

/*
 ExperimentGroupApiService list all ExperimentGroups By filter condition
  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  * @param layerId list all experiment groups of the layer
  * @param optional nil or *ExperimentGroupApiListExperimentGroupsOpts - Optional Parameters:
	  * @param "Status" (optional.Int32) -  list the status experiment groups of the layer
 @return InlineResponse2009
*/

type ExperimentGroupApiListExperimentGroupsOpts struct {
	Status optional.Uint32
}

func (a *ExperimentGroupApiService) ListExperimentGroups(ctx context.Context, layerId int64, localVarOptionals *ExperimentGroupApiListExperimentGroupsOpts) (ListExperimentGroupsResponse, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ListExperimentGroupsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/experiment_groups/all"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("layer_id", parameterToString(layerId, ""))
	if localVarOptionals != nil && localVarOptionals.Status.IsSet() {
		localVarQueryParams.Add("status", parameterToString(localVarOptionals.Status.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			err = fmt.Errorf("failed to decode resp, err=%w, body=%s", err, string(localVarBody))
			return localVarReturnValue, err
		}
		return localVarReturnValue, errors.New(fmt.Sprintf("Http Status code:%d", localVarHttpResponse.StatusCode))
	} else {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			err = fmt.Errorf("failed to decode resp, err=%w, body=%s", err, string(localVarBody))
			return localVarReturnValue, err
		}
	}
	return localVarReturnValue, nil
}

/*
ExperimentGroupApiService change the status of experiment group to offline
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param expGroupId ExperimentGroup Id to get experiment group data

@return Response
*/
func (a *ExperimentGroupApiService) OfflineExperimentGroup(ctx context.Context, expGroupId int64) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/experiment_groups/{exp_group_id}/offline"
	localVarPath = strings.Replace(localVarPath, "{"+"exp_group_id"+"}", fmt.Sprintf("%v", expGroupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
		return localVarReturnValue, errors.New(fmt.Sprintf("Http Status code:%d", localVarHttpResponse.StatusCode))
	} else {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
	}
	return localVarReturnValue, nil
}

/*
ExperimentGroupApiService change the status of experiment group to online
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param expGroupId ExperimentGroup Id to get experiment  group data

@return Response
*/
func (a *ExperimentGroupApiService) OnlineExperimentGroup(ctx context.Context, expGroupId int64) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/experiment_groups/{exp_group_id}/online"
	localVarPath = strings.Replace(localVarPath, "{"+"exp_group_id"+"}", fmt.Sprintf("%v", expGroupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
		return localVarReturnValue, errors.New(fmt.Sprintf("Http Status code:%d", localVarHttpResponse.StatusCode))
	} else {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
	}
	return localVarReturnValue, nil
}

/*
ExperimentGroupApiService update experiment group  data
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body ExperimentGroup object that needs to be update
  - @param expGroupId ID of experiment group to update

@return Response
*/
func (a *ExperimentGroupApiService) UpdateExperimentGroup(ctx context.Context, body model.ExperimentGroup, expGroupId int64) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/experiment_groups/{exp_group_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"exp_group_id"+"}", fmt.Sprintf("%v", expGroupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode != 200 {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
		return localVarReturnValue, errors.New(fmt.Sprintf("Http Status code:%d", localVarHttpResponse.StatusCode))
	} else {
		err = a.client.decodeResponse(&localVarReturnValue, localVarBody)
		if err != nil {
			return localVarReturnValue, err
		}
	}
	return localVarReturnValue, nil
}
