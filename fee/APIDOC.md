# fee 服务接口文档

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

## CreateNodeDistributeBill

描述: 创建新的机器节点独占账单

原型定义: `rpc CreateNodeDistributeBill(CreateNodeDistributeBillRequest) returns (CreateNodeDistributeBillResponse) {}`

请求参数:

```protobuf
message CreateNodeDistributeBillRequest {
    request.BaseRequest baseRequest = 1;
    int32 nodeDistributeID = 2;
}
```

响应参数:

```protobuf
message CreateNodeDistributeBillResponse {
    int32 id = 1;
}
```

## PaginationGetNodeDistributeBill

描述: 分页查询计算节点独占账单

原型定义: `rpc PaginationGetNodeDistributeBill(PaginationGetNodeDistributeBillRequest) returns (PaginationGetNodeDistributeBillResponse) {}`

请求参数:

```protobuf
message PaginationGetNodeDistributeBillRequest {
    request.BaseRequest baseRequest = 1;
    int32 pageIndex = 2;
    int32 pageSize = 3;
}
```

响应参数:

```protobuf
message PaginationGetNodeDistributeBillResponse {
    repeated fee.NodeDistributeBill bills = 1;
    int32 count = 2;
}
```
