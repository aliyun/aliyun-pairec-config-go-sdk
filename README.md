# aliyun-pairec-config-go-sdk
Go sdk for PA-REC config server. Aliyun product [link](https://pairec.console.aliyun.com/cn-hangzhou/instances).

# Installation

```
go get github.com/aliyun/aliyun-pairec-config-go-sdk 
```

# Usage

```golang
// config server info
host := "" 
token := ""
// 初始化 client, 必须指定 host 和 environment 字段
// environment 合法值为 daily, prepub, product
client, err := NewExperimentClient(host, common.Environment_Prepub_Desc, WithLogger(LoggerFunc(log.Printf)), WithErrorLogger(LoggerFunc(log.Fatalf)), WithToken(token))	
if err != nil {
		t.Fatal(err)
}

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

// get experiment param 
param := experimentResult.GetExperimentParams().GetString("a", "not_exist")

if param != "not_exist" {
    // 实验逻辑
    ...
} else {
    // 默认值逻辑
    ...
}
```
