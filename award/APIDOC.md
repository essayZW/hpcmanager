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

## CreateTechnologyAwardApply

描述: 创建科技奖励申请

原型定义: `rpc CreateTechnologyAwardApply(CreateTechnologyAwardApplyRequest) returns (CreateTechnologyAwardApplyResponse) {}`

请求参数:

```protobuf
message CreateTechnologyAwardApplyRequest {
    request.BaseRequest baseRequest = 1;
    int32 projectID = 2;
    string prizeLevel = 3;
    string prizeImageName = 4;
    string remarkMessage = 5;
}
```

响应参数:

```protobuf
message CreateTechnologyAwardApplyResponse {
    int32 id = 1;
}
```

## PaginationGetTechnologyApply

描述: 分页查询科技奖励申请信息

原型定义: `rpc PaginationGetTechnologyApply(PaginationGetTechnologyApplyRequest) returns (PaginationGetTechnologyApplyResponse) {}`

请求参数:

```protobuf
message PaginationGetTechnologyApplyRequest {
    request.BaseRequest baseRequest = 1;
    int32 pageIndex = 2;
    int32 pageSize = 3;
}
```

响应参数:

```protobuf
message PaginationGetTechnologyApplyResponse {
    int32 count = 1;
    repeated award.TechnologyApply applies = 2;
}
```

# 附录

## PaperApply

描述: 论文奖励申请信息

```protobuf
message PaperApply {
    int32 id = 1;
    int32 createrID = 2;
    string createrUsername = 3;
    string createrName = 4;
    int32 userGroupID = 5;
    int32 tutorID = 6;
    string tutorUsername =  7;
    string tutorName = 8;
    string paperTitle = 9;
    string paperCategory = 10;
    string paperPartition = 11;
    string paperFirstPageImageName = 12;
    string paperThanksPageImageName = 13;
    string remarkMessage = 14;
    int32 checkStatus = 15;
    int32 checkerID = 16;
    string checkerUsername = 17;
    string checkerName = 18;
    double checkMoney = 19;
    string checkMessage = 20;
    int64 checkTimeUnix = 21;
    string extraAttributes = 22;
    int64 createTimeUnix = 23;
}
```

## TechnologyApply

描述: 科技奖励申请

```protobuf
message TechnologyApply {
    int32 id = 1;
    int32 CreaterID = 2;
    string CreaterUsername = 3;
    string CreaterName = 4;
    int64 CreateTimeUnix = 5;
    int32 UserGroupID = 6;
    int32 TutorID = 7;
    string TutorUsername = 8;
    string TutorName = 9;
    int32 ProjectID = 10;
    string ProjectName = 11;
    string ProjectDescription = 12;
    string PrizeLevel = 13;
    string PrizeImageName = 14;
    string RemarkMessage = 15;
    int32 CheckStatus = 16;
    int32 CheckerID = 17;
    string CheckerName = 18;
    string CheckerUsername = 19;
    string CheckMessage = 20;
    int64 CheckTimeUnix = 21;
    double CheckMoney = 22;
    string ExtraAttributes = 23;
}
```
