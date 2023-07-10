package api

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type CrowdApiService service

/*
CrowdApiService Get Crowd users By crowd ID
Get Crowd users By crowd ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param sceneId Scene Id to get scene info
@return InlineResponse2001
*/
func (a *CrowdApiService) GetCrowdUsersById(ctx context.Context, crowdId int64) (ListCrowdUsersResponse, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue ListCrowdUsersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/crowds/{crowd_id}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"crowd_id"+"}", fmt.Sprintf("%v", crowdId), -1)

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
