# fss 服务接口文档

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

## StoreFile

描述: 存储文件接口,返回存储的文件名

原型定义: `rpc StoreFile(StoreFileRequest) returns (StoreFileResponse) {}`

请求参数:

```protobuf
message StoreFileRequest {
    request.BaseRequest baseRequest = 1;
    string fileName = 2;
    bytes file = 3;
}
```

响应参数:

```protobuf
message StoreFileResponse {
    string filePath = 1;
}

```
