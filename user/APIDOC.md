# user 服务接口文档

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

## Login

描述：登录认证接口，在 redis 中创建对应的登录 token

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

描述：验证登录 token 接口，验证传入的 token 信息并返回对应用户的相关信息以及权限信息

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

描述：添加一个新的用户，返回新添加用户的用户 id

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

描述：传入用户名，为其创建登录 token

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

描述：传入目标用户 ID，查询用户的详细个人信息，对于不同权限的查询者其能查询的范围不同

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

描述：传入目标组的 ID，查询组的基本信息，对于不同权限的查询者其能查询的范围不同

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

## PaginationGetUserInfo

描述：传入分页大小以及页码，查询一系列用户的基本信息，对于导师限定只可以查询其组内的用户，对于管理员可以查询所有的用户信息

需求权限：`Tutor`及以上

原型定义：`rpc PaginationGetUserInfo(PaginationGetUserInfoRequest) returns (PaginationGetUserInfoResponse) {}`

请求参数：

```protobuf
message PaginationGetUserInfoRequest {
    request.BaseRequest baseRequest = 1;
    int32 pageSize = 2;
    int32 pageIndex = 3;
}
```

响应参数：

```protobuf
message PaginationGetUserInfoResponse {
    repeated user.UserInfo userInfos = 1;
    int32 count = 2;
}
```

## CreateJoinGroupApply

描述：传入申请的组的 ID 信息，创建一个当前用户申请加入组的新的申请

需求权限：仅`Guest`权限

原型定义：`rpc CreateJoinGroupApply(CreateJoinGroupApplyRequest) returns (CreateJoinGroupApplyResponse) {}`

请求参数：

```protobuf
message CreateJoinGroupApplyRequest {
    request.BaseRequest baseRequest = 1;
    int32 applyGroupID = 2;
}
```

响应参数：

```protobuf
message CreateJoinGroupApplyResponse {
    bool success = 1;
    int32 applyID = 2;
}
```

## SearchTutorInfo

描述：传入导师的用户名，搜索其对应组的信息以及导师的其他基本信息

需求权限：仅`Guest`权限

原型定义：`rpc SearchTutorInfo(SearchTutorInfoRequest) returns (SearchTutorInfoResponse) {}`

请求参数：

```protobuf
message SearchTutorInfoRequest {
    request.BaseRequest baseRequest = 1;
    string username = 2;
}
```

响应参数：

```protobuf
message SearchTutorInfoResponse {
    int32 tutorID = 1;
    string tutorUsername = 2;
    string tutorName = 3;
    int32 groupID = 4;
    string groupName = 5;
}
```

## PageGetApplyGroupInfo

描述：分页查询申请记录，对于不同权限的用户所能查找的范围不同

需求权限：无

原型定义：`rpc PageGetApplyGroupInfo(PageGetApplyGroupInfoRequest) returns (PageGetApplyGroupInfoResponse) {}`

请求参数：

```protobuf
message PageGetApplyGroupInfoRequest {
    request.BaseRequest baseRequest = 1;
    int32 pageIndex = 2;
    int32 pageSize = 3;
}
```

响应参数：

```protobuf
message PageGetApplyGroupInfoResponse {
    repeated user.UserGroupApply applies = 1;
    int32 count = 2;
}
```

## CheckApply

描述：审核用户加入组申请

需求权限：`Tutor`及以上

原型定义：`rpc CheckApply(CheckApplyRequest) returns (CheckApplyResponse) {}`

请求参数：

```protobuf
message CheckApplyRequest {
    request.BaseRequest baseRequest = 1;
    int32 applyID = 2;
    bool checkStatus = 3;
    string checkMessage = 4;
    bool tutorCheck = 5;
}
```

响应参数：

```protobuf
message CheckApplyResponse {
    bool success = 1;
}
```

## CreateGroup

描述：创建新的用户组

需求权限：`CommonAdmin`及以上

原型定义：`rpc CreateGroup(CreateGroupRequest) returns (CreateGroupResponse) {}`

请求参数：

```protobuf
message CreateGroupRequest {
    request.BaseRequest baseRequest = 1;
    int32 tutorID = 2;
    string name = 3;
    string queueName = 4;
}
```

响应参数：

```protobuf
message CreateGroupResponse {
    bool success = 1;
    int32 groupID = 2;
}
```

## JoinGroup

描述：添加现有的没有组的用户到一个组中

需求权限：`CommonAdmin`及以上

原型定义：`rpc JoinGroup(JoinGroupRequest) returns (JoinGroupResponse) {}`

请求参数：

```protobuf
message JoinGroupRequest {
    request.BaseRequest baseRequest = 1;
    int32 userID = 2;
    int32 groupID = 3;
}
```

响应参数：

```protobuf
message JoinGroupResponse {
    bool success = 1;
}
```

## Logout

描述: 用户退出登录,销毁 token

原型定义: `rpc Logout(LogoutRequest) returns (LogoutResponse) {}`

需求权限: 无

请求参数:

```protobuf
message LogoutRequest {
    request.BaseRequest baseRequest = 1;
    string username = 2;
}
```

响应参数:

```protobuf
message LogoutResponse {
    bool success = 1;
}
```

## GetUserInfoByHpcID

描述: 通过 hpc_user 的 id 查询对应的用户的信息

原型定义: `rpc GetUserInfoByHpcID(GetUserInfoByHpcIDRequest) returns (GetUserInfoByHpcIDResponse) {}`

需求权限: `Common`及以上

请求参数:

```protobuf
message GetUserInfoByHpcIDRequest {
    request.BaseRequest baseRequest = 1;
    int32 hpcUserID = 2;
}
```

响应参数:

```protobuf
message GetUserInfoByHpcIDResponse {
    user.UserInfo info = 1;
}
```

## GetApplyInfoByID

描述: 通过 ID 查询用户申请加入组信息

原型定义: `rpc GetApplyInfoByID(GetApplyInfoByIDRequest) returns (GetApplyInfoByIDResponse) {}`

需求权限: 无

请求参数:

```protobuf
message GetApplyInfoByIDRequest {
    request.BaseRequest baseRequest = 1;
    int32 applyID = 2;
}
```

响应参数:

```protobuf
message GetApplyInfoByIDResponse {
    user.UserGroupApply apply = 1;
}
```

# 附录

## UserInfo

描述：UserInfo 消息

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
    int64 createTime = 10;
    string extraAttributes = 11;
    int32 hpcUserID = 12;
}
```

## GroupInfo

描述：Group 相关的信息

```protobuf
// GroupInfo 用户组基本信息
message GroupInfo {
    int32 id = 1;
    string name = 2;
    int64 createTime = 3;
    int32 createrID = 4;
    string createrName = 5;
    string createrUsername = 6;
    int32 tutorID = 7;
    string tutorName = 8;
    string tutorUsername = 9;
    double balance = 10;
    string extraAttributes = 11;
    int32 hpcGroupID = 12;
}
```

## UserGroupApply

描述：group 申请相关的信息

```protobuf
// UserGroupApply 用户加入组申请记录信息
message UserGroupApply {
    int32 id = 1;
    int32 userID = 2;
    string userUsername = 3;
    string userName = 4;
    int32 applyGroupID = 5;
    int32 tutorID = 6;
    string tutorUsername = 7;
    string tutorName = 8;
    int32 tutorCheckStatus = 9;
    int32 managerCheckStatus = 10;
    int32 status = 11;
    string messageTutor = 12;
    string messageManager = 13;
    int64 tutorCheckTime = 14;
    int64 managerCheckTime = 15;
    int32 managerCheckerID = 16;
    string managerCheckerUsername = 17;
    string managerCheckerName = 18;
    int64 createTime = 19;
    string extraAttributes = 20;
}
```
