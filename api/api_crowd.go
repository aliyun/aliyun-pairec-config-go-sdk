package api

type CrowdApiService service

/*
CrowdApiService Get Crowd users By crowd ID
Get Crowd users By crowd ID
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param sceneId Scene Id to get scene info

@return InlineResponse2001
*/
func (a *CrowdApiService) GetCrowdUsersById(crowdId int64) (ListCrowdUsersResponse, error) {
	var (
		localVarReturnValue ListCrowdUsersResponse
	)

	return localVarReturnValue, nil
}
