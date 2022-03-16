# project 服务接口文档

## Ping

描述：ping 测试，返回请求 ID 等基本信息

原型定义：`rpc Ping(publicproto.Empty) returns (publicproto.PingResponse) {}`

请求参数：

```protobuf
// 空消息
message Empty {
    request.BaseRequest baseRequest = 1;
}
```

响应参数：

```protobuf
// PingResponse ping消息的回复,一般内容为pong
message PingResponse {
    string Msg = 1;
    string Ip = 2;
    string RequestId = 3;
}
```

## CreateProject

描述: 创建一条新的 project 记录

原型定义: `rpc CreateProject(CreateProjectRequest) returns (CreateProjectResponse) {}`

需求权限: guest 及以上

请求参数:

```protobuf
message CreateProjectRequest {
    request.BaseRequest baseRequest = 1;
    project.ProjectInfo projectInfo = 2;
}
```

响应参数:

```protobuf
message CreateProjectResponse {
    int32 projectID = 1;
}
```

# 附录

## ProjectInfo

描述: project 数据的消息映射

```protobuf
message ProjectInfo {
    int32 id = 1;
    string name = 2;
    string from = 3;
    string numbering = 4;
    string expenses = 5;
    string description = 6;
    int32 createrUserID = 7;
    string createrUsername = 8;
    string createrName = 9;
    int64 createTime = 10;
    int32 modifyUserID = 11;
    string modifyUsername = 12;
    string modifyName = 13;
    int64 modifyTime = 14;
    string extraAttributes = 15;
}
```
