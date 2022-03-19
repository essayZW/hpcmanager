# node 服务接口文档

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

# CreateNodeApply

描述: 创建新的机器节点包机申请

原型定义: `rpc CreateNodeApply(CreateNodeApplyRequest) returns (CreateNodeApplyResponse) {}`

需求权限: `Common` 及以上

请求参数:

```protobuf
message CreateNodeApplyRequest {
    request.BaseRequest baseRequest = 1;
    int32 projectID = 2;
    string nodeType = 3;
    int32 nodeNum = 4;
    int64 startTime = 5;
    int64 endTime = 6;
}
```

响应参数:

```protobuf
message CreateNodeApplyResponse {
   int32 id = 1;
}
```
