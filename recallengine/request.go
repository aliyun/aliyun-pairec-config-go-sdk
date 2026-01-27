package recallengine

type Request struct {
	RequestId string `json:"request_id"`
}

type RecallConf struct {
	Trigger string `json:"trigger"`
	Count   int    `json:"count"`
}

type RecallRequest struct {
	Request
	InstanceId    string                `json:"instance_id"`
	Service       string                `json:"service"`
	Version       string                `json:"version"`
	Uid           string                `json:"uid"`
	Recalls       map[string]RecallConf `json:"recalls"`
	ExposureList  string                `json:"exposure_list"`
	ContextParams map[string]any        `json:"context_params"`
	Debug         bool                  `json:"debug"`
}

type WriteRequest struct {
	Request
	Content   []map[string]any `json:"content"`
	VersionId string
}
