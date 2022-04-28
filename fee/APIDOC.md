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

## GetNodeDistributeFeeRate

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

## CreateNodeWeekUsageBill

描述: 创建机器节点时长周账单

原型定义: `rpc CreateNodeWeekUsageBill(CreateNodeWeekUsageBillRequest) returns (CreateNodeWeekUsageBillResponse) {}`

请求参数:

```protobuf
message CreateNodeWeekUsageBillRequest {
    request.BaseRequest baseRequest = 1;
    int32 nodeWeekUsageRecordID = 2;
}
```

响应参数:

```protobuf
message CreateNodeWeekUsageBillResponse {
    int32 id = 1;
}
```

## PaginationGetNodeWeekUsageBillRecords

描述: 分页查询机器节点机时周账单

原型定义: `rpc PaginationGetNodeWeekUsageBillRecords(PaginationGetNodeWeekUsageBillRecordsResquest) returns (PaginationGetNodeWeekUsageBillRecordsResponse) {}`

请求参数:

```protobuf
message PaginationGetNodeWeekUsageBillRecordsResquest {
    request.BaseRequest baseRequest = 1;
    int32 pageIndex = 2;
    int32 pageSize = 3;
    int64 startTimeUnix = 4;
    int64 endTimeUnix = 5;
}
```

响应参数:

```protobuf
message PaginationGetNodeWeekUsageBillRecordsResponse {
    int32 count = 1;
    repeated fee.NodeWeekUsageBill bills = 2;
}
```

## PaginationGetUserGroupUsageBillRecords

描述: 分组分页查询账单信息

原型定义: `rpc PaginationGetUserGroupUsageBillRecords(PaginationGetUserGroupUsageBillRecordsRequest) returns (PaginationGetUserGroupUsageBillRecordsResponse) {}`

请求参数:

```protobuf
message PaginationGetUserGroupUsageBillRecordsRequest {
    request.BaseRequest baseRequest = 1;
    int32 groupID = 2;
    int32 pageIndex = 3;
    int32 pageSize = 4;
    bool payFlag = 5;
}
```

响应参数:

```protobuf
message PaginationGetUserGroupUsageBillRecordsResponse {
    int32 count = 1;
    repeated fee.NodeWeekUsageBillForUserGroup bills = 2;
}
```

## PayGroupNodeUsageBill

描述: 支付某个用户组的所有的机器节点时长待缴费账单

原型定义: `rpc PayGroupNodeUsageBill(PayGroupNodeUsageBillRequest) returns (PayGroupNodeUsageBillResponse) {}`

请求参数:

```protobuf
message PayGroupNodeUsageBillRequest {
    request.BaseRequest baseRequest = 1;
    int32 userGroupID = 2;
    int32 payType = 3;
    string payMessage = 4;
    double needFee = 5;
}
```

响应参数:

```protobuf
message PayGroupNodeUsageBillResponse {
    int32 payCount = 1;
}
```

## GetNodeUsageFeeRate

描述: 查询机器节点时间费率

原型定义: `rpc GetNodeUsageFeeRate(GetNodeUsageFeeRateRequest) returns (GetNodeUsageFeeRateResponse) {}`

请求参数:

```protobuf
message GetNodeUsageFeeRateRequest {
    request.BaseRequest baseRequest = 1;
}
```

响应参数:

```protobuf
message GetNodeUsageFeeRateResponse {
    double cpu = 1;
    double gpu = 2;
}
```

## CreateNodeQuotaModifyBill

描述: 创建机器存储账单

原型定义: `rpc CreateNodeQuotaModifyBill(CreateNodeQuotaModifyBillRequest) returns (CreateNodeQuotaModifyBillResponse) {}`

请求参数:

```protobuf
message CreateNodeQuotaModifyBillRequest {
    request.BaseRequest baseRequest = 1;
    int32 userID = 2;
    int32 oldSize = 3;
    int32 newSize = 4;
    int64 oldEndTimeUnix = 5;
    int64 newEndTimeUnix = 6;
    bool quotaSizeModify = 7;
}
```

响应参数:

```protobuf
message CreateNodeQuotaModifyBillResponse {
    int32 id = 1;
}
```

## PaginationGetNodeQuotaBill

描述: 分页查询机器存储账单

原型定义: `rpc PaginationGetNodeQuotaBill(PaginationGetNodeQuotaBillRequest) returns (PaginationGetNodeQuotaBillResponse) {}`

请求参数:

```protobuf
message PaginationGetNodeQuotaBillRequest {
    request.BaseRequest baseRequest = 1;
    int32 pageIndex = 2;
    int32 pageSize = 3;
}
```

响应参数:

```protobuf
message PaginationGetNodeQuotaBillResponse {
    int32 count = 1;
    repeated fee.NodeQuotaBill bills = 2;
}
```

## GetNodeQuotaFeeRate

描述: 查询机器存储费率

原型定义: `rpc GetNodeQuotaFeeRate(GetNodeQuotaFeeRateRequest) returns (GetNodeQuotaFeeRateResponse) {}`

请求参数:

```protobuf
message GetNodeQuotaFeeRateRequest {
    request.BaseRequest baseRequest = 1;
}
```

响应参数:

```protobuf
message GetNodeQuotaFeeRateResponse {
    double basic = 1;
    double extra = 2;
}

```

## PayNodeQuotaBill

描述: 支付机器存储账单

原型定义: `rpc PayNodeQuotaBill(PayNodeQuotaBillRequest) returns (PayNodeQuotaBillResponse) {}`

请求参数:

```protobuf
message PayNodeQuotaBillRequest {
    request.BaseRequest baseRequest = 1;
    int32 billID = 2;
    double payMoney = 3;
    int32 payType = 4;
    string payMessage = 5;
}
```

响应参数:

```protobuf
message PayNodeQuotaBillResponse {
    bool success = 1;
}
```

## SetNodeDistributeFeeRate

描述: 设置节点分配账单费率

原型定义: `rpc SetNodeDistributeFeeRate(SetNodeDistributeFeeRateRequest) returns (SetNodeDistributeFeeRateResponse) {}`

请求参数:

```protobuf
message SetNodeDistributeFeeRateRequest {
    request.BaseRequest baseRequest = 1;
    double rate36CPU = 2;
    double rate4GPU = 3;
    double rate8GPU = 4;
}
```

响应参数:

```protobuf
message SetNodeDistributeFeeRateResponse {
    bool success = 1;
}
```

## SetNodeUsageFeeRate

描述: 设置机时费率

原型定义: `rpc SetNodeUsageFeeRate(SetNodeUsageFeeRateRequest) returns (SetNodeUsageFeeRateResponse) {}`

请求参数:

```protobuf
message SetNodeUsageFeeRateRequest {
    request.BaseRequest baseRequest = 1;
    double cpu = 2;
    double gpu = 3;
}
```

响应参数:

```protobuf
message SetNodeUsageFeeRateResponse {
    bool success = 1;
}
```

## SetNodeQuotaFeeRate

描述: 设置机器存储费率

原型定义: `rpc SetNodeQuotaFeeRate(SetNodeQuotaFeeRateRequest) returns (SetNodeQuotaFeeRateResponse) {}`

请求参数:

```protobuf
message SetNodeQuotaFeeRateRequest {
    request.BaseRequest baseRequest = 1;
    double basic = 2;
    double extra = 3;
}
```

响应参数:

```protobuf
message SetNodeQuotaFeeRateResponse {
    bool success = 1;
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

## NodeWeekUsageBill

描述: 机器时长周账单定义

```protobuf
message NodeWeekUsageBill {
    int32 id = 1;
    int32 userID = 2;
    string username = 3;
    string name = 4;
    int32 wallTime = 5;
    int32 gwallTime = 6;
    double fee = 7;
    double payFee = 8;
    int64 startTime = 9;
    int64 endTime = 10;
    int32 payFlag = 11;
    int64 payTime = 12;
    int32 payType = 13;
    string payMessage = 14;
    int32 userGroupID = 15;
    int64 createTime = 16;
    string extraAttributes = 17;
}
```

## NodeWeekUsageBillForUserGroup

描述: 机器时长账单分组查询

```protobuf
message NodeWeekUsageBillForUserGroup {
    int32 wallTime = 1;
    int32 gwallTime = 2;
    double fee = 3;
    double payFee = 4;
    int32 userGroupID = 6;
}
```
