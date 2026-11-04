package model

type FeatureConsistencyJob struct {
	JobId                                int    `json:"job_id"`
	JobName                              string `json:"job_name,omitempty"`
	SceneId                              int    `json:"scene_id"`
	SampleRate                           int    `json:"sample_rate"`
	ModelType                            string `json:"model_type"`
	FeatureBackflowQueueType             string `json:"feature_backflow_queue_type"`
	FeatureBackflowQueueDatahubAccessId  string `json:"feature_backflow_queue_datahub_access_id,omitempty"`
	FeatureBackflowQueueDatahubAccessKey string `json:"feature_backflow_queue_datahub_access_key,omitempty"`
	FeatureBackflowQueueDatahubEndpoint  string `json:"feature_backflow_queue_datahub_endpoint,omitempty"`
	FeatureBackflowQueueDatahubProject   string `json:"feature_backflow_queue_datahub_project,omitempty"`
	FeatureBackflowQueueDatahubTopic     string `json:"feature_backflow_queue_datahub_topic,omitempty"`
	FeatureBackflowEASVpcAddress         string `json:"feature_backflow_eas_vpc_address"`
	FeatureBackflowEASToken              string `json:"feature_backflow_eas_token"`
	EasModelUrl                          string `json:"eas_model_url"`
	EasModelServiceName                  string `json:"eas_model_service_name"`
	NeedFeatureReply                     int    `json:"need_feature_reply"`
	FeatureReplyHost                     string `json:"feature_reply_host,omitempty"`
	FeatureReplyToken                    string `json:"feature_reply_token,omitempty"`
	FeatureReplyQueueType                string `json:"feature_reply_queue_type,omitempty"`
	FeatureReplyQueueDatahubAccessId     string `json:"feature_reply_queue_datahub_access_id,omitempty"`
	FeatureReplyQueueDatahubAccessKey    string `json:"feature_reply_queue_datahub_access_key,omitempty"`
	FeatureReplyQueueDatahubEndpoint     string `json:"feature_reply_queue_datahub_endpoint,omitempty"`
	FeatureReplyQueueDatahubProject      string `json:"feature_reply_queue_datahub_project,omitempty"`
	FeatureReplyQueueDatahubTopic        string `json:"feature_reply_queue_datahub_topic,omitempty"`
	FeatureReplyEASVpcAddress            string `json:"feature_reply_eas_vpc_address"`
	FeatureReplyEASToken                 string `json:"feature_reply_eas_token"`
	Status                               int    `json:"status"`
	StartTime                            int64  `json:"start_time"`
	EndTime                              int64  `json:"end_time"`
	// SkipFeatureReply is a runtime-only field delivered via t_param.param_value JSON.
	// When true, the engine carries the special debug level in the online scoring
	// request so the model service writes features to OSS CSV directly, skipping the
	// /api/feature_reply relay. A missing field decodes to false (backward compatible).
	SkipFeatureReply bool `gorm:"-" json:"skip_feature_reply"`
	// LatestHistoryJobId is the history id of the current job run, delivered via
	// t_param.param_value JSON. The engine uses it as the task_id meta_data value so
	// the model service lands OSS CSV under a per-run subdirectory
	// ({log_oss_path}/{task_id}/{uuid}.csv). It is unique per run, unlike JobId which
	// is stable across runs and would pile up historical CSVs in one directory.
	LatestHistoryJobId int `json:"latest_history_job_id"`
}
