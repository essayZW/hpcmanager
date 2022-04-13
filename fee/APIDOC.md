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

## PayNodeDistributeBill

描述: 支付机器独占账单

原型定义: `rpc PayNodeDistributeBill(PayNodeDistributeBillRequest) returns (PayNodeDistributeBillResponse) {}`

请求参数:

```protobuf
message PayNodeDistributeBillRequest {
    request.BaseRequest baseRequest = 1;
    int32 id = 2;
    double payMoney = 3;
    int32 payType = 4;
    string payMessage = 5;
}
```

响应参数:

```protobuf
message PayNodeDistributeBillResponse {
    bool success = 1;
}
```

### GetNodeDistributeFeeRate

描述: 查询机器节点独占费率

原型定义: `rpc GetNodeDistributeFeeRate(GetNodeDistributeFeeRateRequest) returns (GetNodeDistributeFeeRateResponse) {}`

请求参数:

```protobuf
message GetNodeDistributeFeeRateRequest {
    request.BaseRequest baseRequest = 1;
}
```

响应参数:

```protobuf
message GetNodeDistributeFeeRateResponse {
    double rate36CPU = 1;
    double rate4GPU = 2;
    double rate8GPU = 3;
}
```

# 附录

## NodeDistributeBill

描述: 机器独占账单消息定义

```protobuf
message NodeDistributeBill {
    int32 id = 1;
    int32 applyID = 2;
    int32 nodeDistributeID = 3;
    double fee = 4;
    double payFee = 5;
    int32 payFlag = 6;
    int64 payTimeMilliUnix = 7;
    int32 payType = 8;
    string payMessage = 9;
    int32 userID = 10;
    string userUsername = 11;
    string userName = 12;
    int32 userGroupID = 13;
    int64 createTimeMilliUnix = 14;
    string extraAttributes = 15;
}
```
