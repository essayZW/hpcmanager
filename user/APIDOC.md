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

