package recallengine

type RecallResponse struct {
	Result *Record
}

type Response struct {
	RequestId string                 `json:"request_id,omitempty"`
	Code      string                 `json:"code"`
	Message   string                 `json:"message,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

type WriteResponse struct {
	Response
}
