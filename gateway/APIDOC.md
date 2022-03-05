# gateway服务接口文档

目前所有HTTP接口都以`/api`为前缀

**所有的API请求必须附加access_token=token进行身份验证**

## user控制器

### /user/ping

Method: GET

描述：进行user服务的ping测试

参数：无

响应：请求ID、PONG

### /user/token

Method: POST

描述：进行用户的登录验证，返回生成的登录 token

参数：

```typescript
interface Login {
    "username": string,
    "password": string
}
```

响应：

生成的Token以及登录的用户的基础信息

### /user/token

Method: GET

描述：通过用户的token查询对应的用户信息

参数：

无

响应：

用户的基本信息

## hpc控制器

### /hpc/ping

Method: GET

描述：进行hpc服务的ping测试

参数：无

响应：请求ID、PONG

## permission控制器

### /permission/ping

Method: GET

描述：进行permission服务的ping测试

参数：无

响应：请求ID、PONG

## group控制器

### /group/ping

Method: GET

描述：进行group服务的ping测试

参数：无

响应：请求ID、PONG

## system控制器

### /sys/install

Method: POST

描述：进行系统的初始化

参数：

```typescript
interface CreateUserParam {
    "username": string;
    "password": string;
    "tel"?: string;
    "email"?: string;
    "name": string;
    "collegeName": string;
}
```

响应：无