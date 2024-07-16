package model

import "math/big"

type TrafficControlTaskTrafficData struct {
	TrafficControlTaskId string        `json:"traffic_control_task_id"`
	InstanceId           string        `json:"instance_id"`
	Environment          string        `json:"environment"`
	Traffics             []TrafficData `json:"traffics"`
}

type TrafficData struct {
	TrafficControlTargetId         string  `json:"traffic_control_target_id"`
	RecordTime                     string  `json:"record_time"`
	ItemOrExperimentId             string  `json:"item_or_experiment_id"`
	TrafficControlTargetTraffic    big.Int `json:"traffic_control_target_traffic"`
	TrafficControlTargetAimTraffic float64 `json:"traffic_control_target_aim_traffic"`
	TrafficControlTaskTraffic      big.Int `json:"traffic_control_task_traffic"`
}
