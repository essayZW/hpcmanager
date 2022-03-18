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

### /user/:id

Method: GET

描述: 根据用户 ID 查询用户信息

参数: 用户 ID

响应: 用户信息

### /user/name/:username

Method: GET

描述: 根据用户的账户查询用户对应的 ID

参数: 用户账户

响应: 用户的基础信息包括 ID 姓名以及组 ID

### /user

Method: GET

描述: 分页查询用户信息

参数:

```text
pageIndex: number,
pageSize: number
```

响应: 分页的用户信息

### /user

Method: PATCH

描述: 修改用户信息，包括用户的邮箱地址，电话以及学院信息

参数:

```go
// UpdateUserInfoParam 用户信息更新参数
type UpdateUserInfoParam struct {
    ID      int    `form:"id" json:"id"`
    Tel     string `form:"tel" json:"tel"`
    Email   string `form:"email" json:"email"`
    College string `form:"college" json:"college"`
}
```

响应: 是否修改成功

## hpc 控制器

### /hpc/ping

Method: GET

描述：进行 hpc 服务的 ping 测试

参数：无

响应：请求 ID、PONG

### /hpc/user/:id

Method: GET

描述: 通过 ID 查询 hpc 用户信息

参数: id,hpc_user 表的主键 ID

响应:

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

### /hpc/group/:id

Method: GET

描述: 通过 ID 查询 hpc 用户组信息

参数: id, hpc_group 表的主键 ID

响应:

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

### /group

Method: GET

描述: 分页查询用户组信息

参数:

```text
pageIndex: number 页码
pageSize: number 页大小
```

响应:

```go
type PaginationQueryResponse struct {
    Count int
    Data  []*userpb.GroupInfo
}
```

### /group/:id

Method: GET

描述: 通过组 ID 查询组信息

参数:

```text
id: 组ID
```

响应: 组信息

### /group

Method: POST

描述: 新建组信息

参数:

```go
// CreateGroupParam 创建组的参数
type CreateGroupParam struct {
    GroupName string `form:"groupName" json:"groupName" binding:"required"`
    QueueName string `form:"queueName" json:"queueName" binding:"required"`
    TutorID   int    `form:"tutorID" json:"tutorID" binding:"required"`
}
```

响应: 组 ID 以及操作是否成功

### /group/apply

Method: GET

参数:

```text
pageIndex: number 页码
pageSize: number 页大小
```

响应:

```go
type PaginationQueryResponse struct {
    Count int
    Data  []*userpb.GroupInfo
}
```

### /group/apply

Method: POST

参数:

```go
// CreateJoinGroupApplyParam 创建加入组申请参数
type CreateJoinGroupApplyParam struct {
    ApplyGroupID int
}
```

响应:

```go
map[string]interface{}{
        "applyID": resp.ApplyID,
}
```

### /group/apply

Method: PATCH

参数:

```go
// CheckJoinGroupApplyParam 审核
type CheckJoinGroupApplyParam struct {
    ApplyID      int    `form:"applyID" json:"applyID" binding:"required"`
    CheckStatus  bool   `form:"checkStatus" json:"checkStatus" binding:"required"`
    CheckMessage string `form:"checkMessage" json:"checkMessage" binding:"required"`
    TutorCheck   bool   `form:"tutorCheck" json:"tutorCheck" binding:"required"`
}
```

响应: 无

### /group/tutor/:username

Method: GET

参数:

```text
:username 用户帐号
```

响应:

```go
map[string]interface{}{
    "tutorUsername": resp.TutorUsername,
    "tutorName":     resp.TutorName,
    "tutorID":       resp.TutorID,
    "groupID":       resp.GroupID,
    "groupName":     resp.GroupName,
}
```

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

## project 控制器

### /project/ping

Method: GET

描述：进行 group 服务的 ping 测试

参数：无

响应：请求 ID、PONG

### /project

Method: POST

描述: 创建新的 project 项目

参数:

```go
// CreateProjectParam 创建新的项目记录请求参数
type CreateProjectParam struct {
    Name        string `form:"name" json:"name" binding:"required"`
    From        string `form:"from" json:"from"`
    Numbering   string `form:"numbering" json:"numbering"`
    Expenses    string `form:"expenses" json:"expenses"`
    Description string `form:"description" json:"description"`
}
```

响应:

```go
map[string]interface{
    "id": id,
}
```

### /project

Method: GET

描述: 分页查询项目信息

参数:

```typescript
{
    pageSize: number,
    pageIndex: number
}
```

响应: 分页查询的结果

## node 控制器

### /node/ping

Method: GET

描述：进行 group 服务的 ping 测试

参数：无

响应：请求 ID、PONG
