# gateway服务接口文档

目前所有HTTP接口都以`/api`为前缀

## /user/ping

Method: GET

描述：进行user服务的ping测试

参数：无

响应：请求ID、PONG

## /user/token

Method: POST

描述：进行用户的登录验证，返回生成的登录 token

参数：

```json
{
    "username": "",
    "password": ""
}
```

响应：

生成的Token以及登录的用户的基础信息