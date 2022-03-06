# permission服务接口文档

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

## GetUserPermission

描述：查询指定用户的拥有的权限信息，返回权限信息，包括权限ID，权限等级，权限名称以及权限描述等信息

原型定义：`rpc GetUserPermission(GetUserPermissionRequest) returns (GetUserPermissionResponse) {}`

请求参数：

```protobuf
message GetUserPermissionRequest {
    request.BaseRequest baseRequest = 1;
    int32 id = 2;
}
```

响应参数：

```protobuf
message GetUserPermissionResponse {
    repeated permission.PermissionInfo info = 1;
}
```

## AddUserPermission

描述：赋予用户新的权限信息，返回操作是否成功状态

原型定义：`rpc AddUserPermission(AddUserPermissionRequest) returns (AddUserPermissionResponse) {}`

请求参数：

```protobuf
message AddUserPermissionRequest {
    request.BaseRequest baseRequest = 1;
    int32 userid = 2;
    int32 userGroupID = 3;
    int32 level = 4;
}
```

响应参数：

```protobuf
message AddUserPermissionResponse {
    bool success = 1;
}
```

## RemoveUserPermission

描述：删除用户的某一个权限，返回操作是否成功的状态

原型定义：`rpc RemoveUserPermission(RemoveUserPermissionRequest) returns (RemoveUserPermissionResponse) {}`

请求参数：

```protobuf
message RemoveUserPermissionRequest {
    request.BaseRequest baseRequest = 1;
    int32 userid = 2;
    int32 level = 3;
}
```

响应参数：

```protobuf
message RemoveUserPermissionResponse {
    bool success = 1;
}
```

## AddPermission

描述：添加一个新的权限等级描述，返回新添加的权限等级的ID

原型定义：`rpc AddPermission(AddPermissionRequest) returns (AddPermissionResponse) {}`

请求参数：

 ```protobuf
 message AddPermissionRequest {
     request.BaseRequest baseRequest = 1;
     permission.PermissionInfo info = 2;
 }
 ```

响应参数：

```protobuf
message AddPermissionResponse {
    int32 permissionID = 1;
}
```

# 附录

## PermissionInfo

描述：permissionInfo消息

```protobuf
message PermissionInfo {
    int32 id = 1;
    string name = 2;
    int32 level = 3;
    string description = 4;
    int64 createTime = 5;
    string extraAttributes = 6;
}
```

