package test

import (
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/experiments"
	"log"
	"os"
)

func CreateExperimentClient(environment string) *experiments.ExperimentClient {
	region := "cn-hangzhou"
	instanceId := os.Getenv("INSTANCE_ID")
	accessId := os.Getenv("ACCESS_ID")
	accessKey := os.Getenv("ACCESS_KEY")
	//address := "pairecservice.cn-hangzhou.aliyuncs.com"
	//preAddress :=
	client, err := experiments.NewExperimentClient(instanceId, region, accessId, accessKey, environment,
		experiments.WithLogger(experiments.LoggerFunc(log.Printf)),
		experiments.WithDomain("pairecservice-pre.cn-hangzhou.aliyuncs.com"))

	if err != nil {
		log.Fatal(err)
	}
	return client
}
