package test

import (
	"log"
	"os"

	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/experiments"
)

func CreateExperimentClient(environment string) *experiments.ExperimentClient {
	region := "cn-hangzhou"

	accessId := os.Getenv("ACCESS_ID")
	accessKey := os.Getenv("ACCESS_KEY")
	instanceId := os.Getenv("INSTANCE_ID")
	//address := "pairecservice." + region + ".aliyuncs.com"
	preAddress := "pairecservice-pre." + region + ".aliyuncs.com"
	client, err := experiments.NewExperimentClient(instanceId, region, accessId, accessKey, environment,
		experiments.WithLogger(experiments.LoggerFunc(log.Printf)),
		experiments.WithDomain(preAddress))

	if err != nil {
		log.Fatal(err)
	}
	return client
}
