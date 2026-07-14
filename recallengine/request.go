package recallengine

type Request struct {
	RequestId string `json:"request_id"`
}

type RecallOptions struct {
	TriggerLimit int `json:"trigger_limit,omitempty"` // per-trigger recall limit, 0 means no limit
	Timeout      int `json:"timeout,omitempty"`       // per-way recall timeout in milliseconds, 0 means no per-way timeout
}

type RecallConf struct {
	Trigger string         `json:"trigger"`
	Count   int            `json:"count"`
	Options *RecallOptions `json:"options,omitempty"`
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
	RetainFields  []string              `json:"retain_fields"`
}

type WriteRequest struct {
	Request
	Content   []map[string]any `json:"content"`
	VersionId string
}
