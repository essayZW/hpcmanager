# gateway 服务接口文档

目前所有 HTTP 接口都以`/api`为前缀

**所有的 API 请求必须附加 access_token=token 进行身份验证**

## user 控制器

### /user/ping

Method: GET

描述：进行 user 服务的 ping 测试

参数：无

响应：请求 ID、PONG

### /user/token

Method: POST

描述：进行用户的登录验证，返回生成的登录 token

参数：

```typescript
interface Login {
  username: string;
  password: string;
}
```

响应：

生成的 Token 以及登录的用户的基础信息

### /user/token

Method: GET

描述：通过用户的 token 查询对应的用户信息

参数：

无

响应：

用户的基本信息

### /user/token

Method: DELETE

描述: 删除用户的登录 token,使用户退出登录

参数:

无

响应:

无

## hpc 控制器

### /hpc/ping

Method: GET

描述：进行 hpc 服务的 ping 测试

参数：无

响应：请求 ID、PONG

## permission 控制器

### /permission/ping

Method: GET

描述：进行 permission 服务的 ping 测试

参数：无

响应：请求 ID、PONG

## group 控制器

### /group/ping

Method: GET

描述：进行 group 服务的 ping 测试

参数：无

响应：请求 ID、PONG

## system 控制器

### /sys/install

Method: POST

描述：进行系统的初始化

参数：

```typescript
interface CreateUserParam {
  username: string;
  password: string;
  tel?: string;
  email?: string;
  name: string;
  collegeName: string;
}
```

响应：无

### /sys/install

Method: GET

描述：查询系统是否已经初始化

参数：无

响应：states 表明是否已经初始化

### /sys/cas/config

Method: GET

描述: 获取系统 cas 配置参数

参数:

```text
// 服务的地址
serviceHost=
```

响应:

```go
type casConfig struct {
    Enable      bool
    AuthServer  string
    ValidPath   string
    ServiceAddr string
}
```

### /sys/cas/valid

Method: GET

描述: 进行 cas 验证的回调验证

参数:

```text
// 票据
ticket=
```

响应:

若验证成功跳转到主页,否则返回错误信息
