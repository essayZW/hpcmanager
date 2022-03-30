# hpc 服务接口文档

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

## AddUserWithGroup

描述：创建用户组并添加用户到这个组中,一般用于新建用户组的时候顺便添加导师用户到组中

原型定义：`rpc AddUserWithGroup(AddUserWithGroupRequest) returns (AddUserWithGroupResponse) {}`

请求参数：

```protobuf
message AddUserWithGroupRequest {
    request.BaseRequest baseRequest = 1;
    string tutorUsername = 2;
    string groupName = 3;
    string queueName = 4;
}
```

响应参数：

```protobuf
message AddUserWithGroupResponse {
    string groupName = 1;
    int32 gid = 2;
    string userName = 3;
    int32 uid = 4;
    int32 hpcGroupID = 5;
    int32 hpcUserID = 6;
}
```

## AddUserToGroup

描述：创建一个新用户并添加到一个已经存在的用户组中

原型定义：`rpc AddUserToGroup(AddUserToGroupRequest) returns (AddUserToGroupResponse) {}`

请求参数：

```protobuf
message AddUserToGroupRequest {
    request.BaseRequest baseRequest = 1;
    string userName = 2;
    int32 hpcGroupID = 3;
}
```

响应参数：

```protobuf
message AddUserToGroupResponse {
    int32 hpcUserID = 1;
    string userName = 2;
    int32 uid = 3;
}
```

## GetUserInfoByID

描述: 通过 ID 查询一个 hpc_user 的信息

原型定义: `rpc GetUserInfoByID(GetUserInfoByIDRequest) returns (GetUserInfoByIDResponse) {}`

请求参数:

```protobuf
message GetUserInfoByIDRequest {
    request.BaseRequest baseRequest = 1;
    int32 hpcUserID = 2;
}
```

响应参数:

```protobuf
message GetUserInfoByIDResponse {
    hpc.HpcUser user = 1;
}
```

## GetGroupInfoByID

描述: 通过 ID 查询 hpc_group 的信息

原型定义: `rpc GetGroupInfoByID(GetGroupInfoByIDRequest) returns (GetGroupInfoByIDResponse) {}`

请求参数:

```protobuf
message GetGroupInfoByIDRequest {
    request.BaseRequest baseRequest = 1;
    int32 hpcGroupID = 2;
}
```

响应参数:

```protobuf
message GetGroupInfoByIDResponse {
    hpc.HpcGroup group = 1;
}
```

## GetUserInfoByUsername

描述: 通过计算节点用户名查询计算节点信息

原型定义: `rpc GetUserInfoByUsername(GetUserInfoByUsernameRequest) returns (GetUserInfoByUsernameResponse) {}`

请求参数:

```protobuf
message GetUserInfoByUsernameRequest {
    request.BaseRequest baseRequest = 1;
    string username = 2;
}
```

响应参数:

```protobuf
message GetUserInfoByUsernameResponse {
    hpc.HpcUser user = 1;
}
```

## GetGroupInfoByGroupName

描述: 通过计算节点用户组的组名查询计算节点信息

原型定义: `rpc GetGroupInfoByGroupName(GetGroupInfoByGroupNameRequest) returns (GetGroupInfoByGroupNameResponse) {}`

请求参数:

```protobuf
message GetGroupInfoByGroupNameRequest {
    request.BaseRequest baseRequest = 1;
    string name = 2;
}
```

响应参数:

```protobuf
message GetGroupInfoByGroupNameResponse {
    hpc.HpcGroup group = 1;
}
```

# 附录

## HpcUser

描述: hpc_user 表的消息

```protobuf
// HpcUser hpc_user表的消息映射
message HpcUser {
    int32 id = 1;
    string nodeUsername = 2;
    int32 nodeUID = 3;
    int32 nodeMaxQuota = 4;
    int64 quotaStartTime = 5;
    int64 quotaEndTime = 6;
    string extraAttributes = 7;
}
```

## HpcGroup

描述: hpc_group 表的消息映射

```protobuf
// HpcGroup hpc_group表的消息映射
message HpcGroup {
    int32 id = 1;
    string name = 2;
    string queueName = 3;
    int32 gID = 4;
    string extraAttributes = 5;
}
```
