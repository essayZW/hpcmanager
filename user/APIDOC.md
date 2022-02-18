# user服务接口文档

## Ping

描述：ping测试，返回请求ID等基本信息

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

## Login

描述：登录认证接口，在redis中创建对应的登录token

需求权限：无

原型定义:`rpc Login(LoginRequest) returns (LoginResponse) {}`

请求参数：

```protobuf
message LoginRequest {
    request.BaseRequest baseRequest = 1;
    string username = 2;
    string password = 3;
}
```

响应参数：

```protobuf
message LoginResponse {
    string token = 1;
    user.UserInfo userInfo = 2;
}
```

## CheckLogin

描述：验证登录token接口，验证传入的token信息并返回对应用户的相关信息以及权限信息

需求权限：无

原型定义：`rpc CheckLogin(CheckLoginRequest) returns (CheckLoginResponse) {}`

请求参数：

```protobuf
message CheckLoginRequest {
    request.BaseRequest baseRequest = 1;
    string token = 2;
}
```

响应参数：

```protobuf
message CheckLoginResponse {
    bool login = 1;
    user.UserInfo userInfo = 2;
    repeated int32 permissionLevel = 3;
}
```

## ExistUsername

描述：查询某个用户名的用户是否存在

需求权限：无

原型定义：`rpc ExistUsername(ExistUsernameRequest) returns (ExistUsernameResponse) {}`

请求参数：

```protobuf
message ExistUsernameRequest {
    request.BaseRequest baseRequest = 1;
    string username = 2;
}
```

响应参数：

```protobuf
message ExistUsernameResponse {
    bool exist = 1;
}
```

## AddUser

描述：添加一个新的用户，返回新添加用户的用户id

需求权限：`SuperAdmin`,`CommonAdmin`

原型定义：`rpc AddUser(AddUserRequest) returns (AddUserResponse) {}`

请求参数：

```protobuf
message AddUserRequest {
    request.BaseRequest baseRequest = 1;
    user.UserInfo userInfo = 2;
}
```

响应参数：

```protobuf
message AddUserResponse {
    int32 userid = 1;
}
```

## CreateToken

描述：传入用户名，为其创建登录token

需求权限：无

原型定义：`rpc createtoken(createtokenrequest) returns (createtokenresponse) {}`

请求参数：

```protobuf
message CreateTokenRequest {
    request.BaseRequest baseRequest = 1;
    string username = 2;
}
```

响应参数：

```protobuf
message CreateTokenResponse {
    string token = 1;
    user.UserInfo userInfo = 2;
}
```

## GetUserInfo

描述：传入目标用户ID，查询用户的详细个人信息，对于不同权限的查询者其能查询的范围不同

需求权限：`Common`及以上

原型定义：`rpc GetUserInfo(GetUserInfoRequest) returns (GetUserInfoResponse) {}`

请求参数：

```protobuf
message GetUserInfoRequest {
    request.BaseRequest baseRequest = 1;
    int32 userid = 2;
}
```

响应参数：

```protobuf
message GetUserInfoResponse {
    user.UserInfo userInfo = 1;
}
```

## GetGroupInfoByID

描述：传入目标组的ID，查询组的基本信息，对于不同权限的查询者其能查询的范围不同

需求权限：`Tutor`及以上

原型定义：`rpc GetGroupInfoByID(GetGroupInfoByIDRequest) returns (GetGroupInfoByIDResponse) {}`

请求参数：

```protobuf
message GetGroupInfoByIDRequest {
    request.BaseRequest baseRequest = 1;
    int32 groupID = 2;
}
```

响应参数：

```protobuf
message GetGroupInfoByIDResponse {
    user.GroupInfo groupInfo = 1;
}
```

## PaginationGetGroupInfo

描述：传入分页大小以及页码，查询一系列组的基本信息，只支持管理员进行查询

需求权限：`CommonAdmin`及以上

原型定义：`rpc PaginationGetGroupInfo(PaginationGetGroupInfoRequest) returns (PaginationGetGroupInfoResponse) {}`

请求参数：

```protobuf
message PaginationGetGroupInfoRequest {
    request.BaseRequest baseRequest = 1;
    int32 pageSize = 2;
    int32 pageIndex = 3;
}
```

响应参数：

```protobuf
message PaginationGetGroupInfoResponse {
    repeated user.GroupInfo groupInfos = 1;
}
```

# 附录

## UserInfo

描述：UserInfo消息

```protobuf
// UserInfo 用户基本信息
message UserInfo {
    int32 id = 1;
    string username = 2;
    string password = 3;
    string tel = 4;
    string email = 5;
    string name = 6;
    string pyName = 7;
    string college = 8;
    int32 groupId = 9;
    int32 createTime = 10;
}
```

## GroupInfo

描述：Group相关的信息

```protobuf
// GroupInfo 用户组基本信息
message GroupInfo {
    int32 id = 1;
    string name = 2;
    string queueName = 3;
    string nodeGroupName = 4;
    int64 createTime = 5;
    int32 createrID = 6;
    string createrName = 7;
    string createrUsername = 8;
    int32 tutorID = 9;
    string tutorName = 10;
    string tutorUsername = 11;
    double balance = 12;
    string extraAttributes = 13;
}
```

