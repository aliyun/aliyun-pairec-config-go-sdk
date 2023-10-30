# aliyun-pairec-config-go-sdk
Go sdk for PA-REC config server. Aliyun product [link](https://pairec.console.aliyun.com/v2).

# Installation

```
go get github.com/aliyun/aliyun-pairec-config-go-sdk/v2
```

# Usage

```golang
// config server info
region := "cn-hangzhou"
instanceId := os.Getenv("INSTANCE_ID")
accessId := os.Getenv("ACCESS_ID")
accessKey := os.Getenv("ACCESS_KEY")
client, err := NewExperimentClient(instanceId, region, accessId, accessKey, environment, WithLogger(LoggerFunc(log.Printf)),  WithErrorLogger(LoggerFunc(log.Fatalf)))

// 具体匹配实验室，构造 ExperimentContext
experimentContext := model.ExperimentContext{
    RequestId: "pvid", // request id
    Uid:       "2115", // uid
    FilterParams: map[string]interface{}{
        "sex": "male",
        "age": 35,
    },
}

// 匹配时，传入场景名称和上下文 ExperimentContext
experimentResult := client.MatchExperiment("home_feed", &experimentContext)

// 打印匹配的信息，可以做日志用
fmt.Println(experimentResult.Info())
// 获取匹配的实验ID
fmt.Println(experimentResult.GetExpId())

// get experiment param value
// ab_param_name 是实验参数名称，如果命中实验，返回实验对应的配置参数，否则返回默认值
param := experimentResult.GetExperimentParams().GetString("ab_param_name", "not_exist")

if param != "not_exist" {
    // 实验逻辑
    ...
} else {
    // 默认值逻辑
    ...
}
```
