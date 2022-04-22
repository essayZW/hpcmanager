# award 服务接口文档

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

## CreatePaperAward

描述: 创建论文奖励申请

原型定义: `rpc CreatePaperAward(CreatePaperAwardRequest) returns (CreatePaperAwardResponse) {}`

请求参数:

```protobuf
message CreatePaperAwardRequest {
    request.BaseRequest baseRequest = 1;
    string title = 2;
    string category = 3;
    string partition = 4;
    string firstPageImageName = 5;
    string thanksPageImageName = 6;
    string remarkMessage = 7;
}
```

响应参数:

```protobuf
message CreatePaperAwardResponse {
    int32 id = 1;
}
```

## PaginationGetPaperApply

描述: 分页查询论文奖励申请信息

原型定义: `rpc PaginationGetPaperApply(PaginationGetPaperApplyRequest) returns (PaginationGetPaperApplyResponse) {}`

请求参数:

```protobuf
message PaginationGetPaperApplyRequest {
    request.BaseRequest baseRequest = 1;
    int32 pageIndex = 2;
    int32 pageSize = 3;
}
```

响应参数:

```protobuf
message PaginationGetPaperApplyResponse {
    int32 count = 1;
    repeated award.PaperApply applies = 2;
}
```

## CheckPaperApplyByID

描述: 审核论文奖励申请信息

原型定义: `rpc CheckPaperApplyByID(CheckPaperApplyByIDRequest) returns (CheckPaperApplyByIDResponse) {}`

请求参数:

```protobuf
message CheckPaperApplyByIDRequest {
    request.BaseRequest baseRequest = 1;
    int32 applyID = 2;
    double money = 3;
    string checkMessage = 4;
}
```

响应参数:

```protobuf
message CheckPaperApplyByIDResponse {
    bool success = 1;
}
```
