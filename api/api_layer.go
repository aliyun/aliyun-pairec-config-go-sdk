package api

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/aliyun/aliyun-pairec-config-go-sdk/model"
)

// Linger please
var (
	_ context.Context
)

type LayerApiService service

/*
LayerApiService Create a new layer
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body Layer object that needs to be added

@return Response
*/
func (a *LayerApiService) AddLayer(ctx context.Context, body model.Layer) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/layers"

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

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
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
LayerApiService Delete layer By layer id
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param layerId Layer Id to delete

@return Response
*/
func (a *LayerApiService) DeleteLayerById(ctx context.Context, layerId int64) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Delete")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/layers/{layer_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"layer_id"+"}", fmt.Sprintf("%v", layerId), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
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
LayerApiService Get layer By layer_id
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param layerId Layer Id to get layer info

@return Response
*/
func (a *LayerApiService) GetLayerById(ctx context.Context, layerId int64) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/layers/{layer_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"layer_id"+"}", fmt.Sprintf("%v", layerId), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
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
LayerApiService list all Layers By filter condition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param expRoomId list all layers of the experiment room

@return Response
*/
func (a *LayerApiService) ListLayers(ctx context.Context, expRoomId int64) (ListLayersResponse, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ListLayersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/layers/all"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("exp_room_id", parameterToString(expRoomId, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
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
LayerApiService update layer data
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body Layer object that needs to be update
  - @param layerId ID of layer to update

@return Response
*/
func (a *LayerApiService) UpdateLayer(ctx context.Context, body model.Layer, layerId int64) (Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Response
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/layers/{layer_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"layer_id"+"}", fmt.Sprintf("%v", layerId), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
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
