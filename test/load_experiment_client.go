package test

import (
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/experiments"
	"log"
	"os"
)

func CreateExperimentClient(environment string) *experiments.ExperimentClient {
	region := "cn-shanghai"

	accessId := os.Getenv("ACCESS_ID")
	accessKey := os.Getenv("ACCESS_KEY")
	instanceId := os.Getenv("INSTANCE_ID")
	address := "pairecservice." + region + ".aliyuncs.com"
	//preAddress := "pairecservice-pre." + region + ".aliyuncs.com"
	client, err := experiments.NewExperimentClient(instanceId, region, accessId, accessKey, environment,
		experiments.WithLogger(experiments.LoggerFunc(log.Printf)),
		experiments.WithDomain(address))

	if err != nil {
		log.Fatal(err)
	}
	return client
}
