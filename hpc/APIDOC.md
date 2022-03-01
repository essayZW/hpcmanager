# hpc服务接口文档

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
