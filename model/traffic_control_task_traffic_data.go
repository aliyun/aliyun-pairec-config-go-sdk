package model

import "math/big"

type TrafficControlTaskTrafficData struct {
	TrafficControlTaskId string        `json:"TrafficControlTaskId"`
	InstanceId           string        `json:"InstanceId"`
	Environment          string        `json:"Environment"`
	Traffics             []TrafficData `json:"Traffics"`
}

type TrafficData struct {
	TrafficControlTargetId         string  `json:"TrafficControlTargetId"`
	RecordTime                     string  `json:"RecordTime"`
	ItemOrExperimentId             string  `json:"ItemOrExperimentId,omitempty"`
	TrafficControlTargetTraffic    big.Int `json:"TrafficControlTargetTraffic,omitempty"`
	TrafficControlTargetAimTraffic float64 `json:"TrafficControlTargetAimTraffic"`
	TrafficControlTaskTraffic      big.Int `json:"TrafficControlTaskTraffic"`
}
