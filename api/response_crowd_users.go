package api

type ListCrowdUsersResponse struct {
	BaseResponse
	Data map[string][]string `json:"data,omitempty"`
}
