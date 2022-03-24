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

## CreateNodeApply

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

## PaginationGetNodeApply

描述: 分页查询用户申请机器节点申请信息

原型定义: `rpc PaginationGetNodeApply(PaginationGetNodeApplyRequest) returns (PaginationGetNodeApplyResponse) {}`

需求权限: `Common` 及以上

请求参数:

```protobuf
message PaginationGetNodeApplyRequest {
    request.BaseRequest baseRequest = 1;
    int32 pageIndex = 2;
    int32 pageSize = 3;
}
```

响应参数:

```protobuf
message PaginationGetNodeApplyResponse {
    repeated node.NodeApply applies = 1;
    int32 count = 2;
}
```

## CheckNodeApply

描述: 审核机器节点包机申请

原型定义: `rpc CheckNodeApply(CheckNodeApplyRequest) returns (CheckNodeApplyResponse) {}`

需求权限: `Tutor` 及以上

请求参数:

```protobuf
message CheckNodeApplyRequest {
    request.BaseRequest baseRequest = 1;
    int32 applyID = 2;
    bool checkStatus = 3;
    string checkMessage = 4;
    bool tutorCheck = 5;
}
```

响应参数:

```protobuf
message CheckNodeApplyResponse {
    bool success = 1;
}
```

## CreateNodeDistributeWO

描述: 创建机器节点分配处理工单

原型定义: `rpc CreateNodeDistributeWO(CreateNodeDistributeWORequest) returns (CreateNodeDistributeWOResponse) {}`

需求权限: `CommonAdmin` 及以上

请求参数:

```protobuf
message CreateNodeDistributeWORequest {
    request.BaseRequest baseRequest = 1;
    int32 applyID = 2;
}
```

响应参数:

```protobuf
message CreateNodeDistributeWOResponse {
    int32 id = 1;
}
```

## PaginationGetNodeDistributeWO

描述: 分页查询机器节点分配处理工单

原型定义: `rpc PaginationGetNodeDistributeWO(PaginationGetNodeDistributeWORequest) returns (PaginationGetNodeDistributeWOResponse) {}`

需求权限: `CommonAdmin` 及以上

请求参数:

```protobuf
message PaginationGetNodeDistributeWORequest {
    request.BaseRequest baseRequest = 1;
    int32 pageIndex = 2;
    int32 pageSize = 3;
}
```

响应参数:

```protobuf
message PaginationGetNodeDistributeWOResponse {
    repeated node.NodeDistribute wos = 1;
    int32 count = 2;
}
```

# 附录

### NodeApply

描述: NodeApply 数据库 node_appl 对应的消息映射

```protobuf
// NodeApply 数据库node_appl对应的消息映射
message NodeApply {
    int32 id = 1;
    int64 createTime = 2;
    int32 createrID = 3;
    string createrUsername = 4;
    string createrName = 5;
    int32 projectID = 6;
    int32 tutorCheckStatus = 7;
    int32 managerCheckStatus = 8;
    int32 status = 9;
    string messageTutor = 10;
    string messageManager = 11;
    int64 tutorCheckTime = 12;
    int32 tutorID = 13;
    string tutorUsername = 14;
    string tutorName = 15;
    int64 managerCheckTime = 16;
    int32 managerCheckerID = 17;
    string managerCheckerUsername = 18;
    string managerCheckerName = 19;
    int64 modifyTime = 20;
    int32 modifyUserID = 21;
    string modifyName = 22;
    string modifyUsername = 23;
    string nodeType = 24;
    int32 nodeNum = 25;
    int64 startTime = 26;
    int64 endTime = 27;
    string extraAttributes = 28;
}
```

### NodeDistribute

描述: NodeDistribute 节点分配处理工单对应的消息映射

```protobuf
// NodeDistribute 节点分配处理工单对应的消息映射
message NodeDistribute {
    int32 id = 1;
    int32 applyID = 2;
    int32 handlerFlag = 3;
    int32 handlerUserID = 4;
    string handlerUsername = 5;
    string handlerName = 6;
    int32 distributeBillID = 7;
    int64 createTime = 8;
    string extraAttributes = 9;
}
```
