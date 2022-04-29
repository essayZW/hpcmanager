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

### /user

Method: POST

描述: 管理员手动添加用户并添加到用户组中

参数:

```go
// CreateUserWithGroup 创建用户并添加到对应的组请求参数
type CreateUserWithGroup struct {
    Tel         string `form:"tel" json:"tel"`
    Email       string `form:"email" json:"email"`
    Name        string `form:"name" json:"name" binding:"required"`
    CollegeName string `form:"collegeName" json:"collegeName"`
    GroupID int `form:"groupID" json:"groupID"`
    Username string `form:"username" json:"username" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}
```

响应: 新用户的用户 ID

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

### /hpc/quota/:hpcID

Method: GET

描述: 通过用户 hpc ID 查询用户存储信息

参数: hpcID

响应:

```go
map[string]interface{}{
        "used":          resp.Used,
        "max":           resp.Max,
        "startTimeUnix": resp.StartTimeUnix,
        "endTimeUnix":   resp.EndTimeUnix,
}
```

### /hpc/quota

Method: PUT

描述: 修改用户存储信息

参数:

```go
// SetUserQuotaParam 设置用户存储信息参数
type SetUserQuotaParam struct {
    HpcUserID           int   `form:"hpcUserID"           json:"hpcUserID"           binding:"required"`
    OldSize             int   `form:"oldSize"             json:"oldSize"             binding:"required"`
    NewSize             int   `form:"newSize"             json:"newSize"             binding:"required"`
    OldEndTimeMilliUnix int64 `form:"oldEndTimeMilliUnix" json:"oldEndTimeMilliUnix" binding:"required"`
    NewEndTimeMilliUnix int64 `form:"newEndTimeMilliUnix" json:"newEndTimeMilliUnix" binding:"required"`
    ModifyData          bool  `form:"modifyData"          json:"modifyData"`
}
```

响应: 是否修改成功

## permission 控制器

### /permission/ping

Method: GET

描述：进行 permission 服务的 ping 测试

参数：无

响应：请求 ID、PONG

### /permission/admin

Method: POST

描述: 将某个用户设置为普通管理员

参数:

```go
// ChangeUserPermissionParam 修改普通管理员用户请求参数
type ChangeUserPermissionParam struct {
    UserID int `form:"userID" json:"userID" binding:"required"`
}
```

响应: 是否添加成功

### /permission/admin

Method: DELETE

描述: 取消某个用户的管理员权限

参数:

```go
// ChangeUserPermissionParam 修改普通管理员用户请求参数
type ChangeUserPermissionParam struct {
    UserID int `form:"userID" json:"userID" binding:"required"`
}
```

响应: 是否删除成功

### /permission/user/:id

Method: GET

描述: 查询某个用户的所有的权限

参数:

```text
id: int
```

响应: 查询到的权限信息

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

### /group/apply/:id

Method: DELETE

描述: 撤销某个申请

参数:

```text
id: number
```

响应: 是否成功

### /group/balance

Method: PATCH

描述: 修改用户组的余额

参数:

```go
// AddGroupBalanceParam 修改用户组的余额参数
type AddGroupBalanceParam struct {
    GroupID int     `form:"groupID" json:"groupID" binding:"required"`
    Balance float64 `form:"balance" json:"balance" binding:"required"`
}
```

响应参数: 修改后的余额

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

### /project/:id

Method: GET

描述: 通过 ID 查询项目信息

参数:

```text
id: number
```

响应: 项目信息

## node 控制器

### /node/ping

Method: GET

描述：进行 group 服务的 ping 测试

参数：无

响应：请求 ID、PONG

### /node/apply

Method: POST

描述: 创建新的计算节点申请记录

参数:

```go
// CreateNodeApplyParam 创建机器节点申请记录请求参数
type CreateNodeApplyParam struct {
    ProjectID int    `form:"projectID" json:"projectID" binding:"required"`
    NodeType  string `form:"nodeType" json:"nodeType" binding:"required"`
    NodeNum   int    `form:"nodeNum" json:"nodeNum" binding:"required"`
    StartTime int64  `form:"startTime" json:"startTime" binding:"required"`
    EndTime   int64  `form:"endTime" json:"endTime" binding:"required"`
}
```

响应: 创建申请记录的 ID

### /node/apply

Method: GET

描述: 分页查询计算节点申请记录

参数:

```typescript
type param = {
  pageIndex: number;
  pageSize: number;
};
```

响应: 分页查询的结果

### /node/apply

Method: PATCH

描述: 审核机器节点申请

参数:

```go
// CreateNodeApplyParam 创建机器节点申请记录请求参数
type CreateNodeApplyParam struct {
    ProjectID int    `form:"projectID" json:"projectID" binding:"required"`
    NodeType  string `form:"nodeType" json:"nodeType" binding:"required"`
    NodeNum   int    `form:"nodeNum" json:"nodeNum" binding:"required"`
    StartTime int64  `form:"startTime" json:"startTime" binding:"required"`
    EndTime   int64  `form:"endTime" json:"endTime" binding:"required"`
}
```

### /node/apply

Method: PUT

描述: 修改机器节点申请信息

参数:

```go
// UpdateNodeApplyParam 更新机器节点申请信息表单
type UpdateNodeApplyParam struct {
    ID        int    `form:"id"        json:"id"        binding:"required"`
    NodeType  string `form:"nodeType"  json:"nodeType"  binding:"required"`
    NodeNum   int    `form:"nodeNum"   json:"nodeNum"   binding:"required"`
    StartTime int64  `form:"startTime" json:"startTime" binding:"required"`
    EndTime   int64  `form:"endTime"   json:"endTime"   binding:"required"`
}

```

响应: 是否修改成功

### /node/apply/:id

Method: DELETE

描述: 撤销某一个机器节点申请记录

参数:

```text
id: number
```

响应: 是否成功

响应: 是否撤销成功

### /node/distribute

Method: GET

描述: 分页查询用户节点分配工单信息

请求参数:

```typescript
type param = {
  pageIndex: number;
  pageSize: number;
};
```

响应参数: 分页的信息

### /node/distribute

Method: PATCH

描述: 处理机器处理分配工单

请求参数:

```go
// FinishNodeDistributeParam 处理机器节点分配工单参数
type FinishNodeDistributeParam struct {
    ID int `form:"id" json:"id" binding:"required"`
}
```

响应参数: 是否成功

### /node/apply/:id

Method: GET

描述: 通过 ID 查询机器节点申请信息

请求参数:

```text
id: number
```

响应参数: 申请信息

### /node/usage

Method: GET

描述: 分页查询一段时间内的机器节点使用记录

请求参数:

```protobuf
pageIndex: number
pageSize: number
startDateMilliUnix: number
endDateMilliUnix: number
```

响应参数: 分页的结果以及总的数量

## fee 控制器

### /fee/ping

Method: GET

描述：进行 group 服务的 ping 测试

参数：无

响应：请求 ID、PONG

### /fee/distribute

Method: GET

描述: 分页查询机器节点独占账单

参数: 分页参数

响应: 查询的数据

### /fee/distribute

Method: PUT

描述: 支付机器独占账单

参数:

```go
// PayNodeDistributeBillParam 支付机器独占账单参数
type PayNodeDistributeBillParam struct {
    ID         int     `form:"id"         json:"id"         binding:"required"`
    PayMoney   float64 `form:"payMoney"   json:"payMoney"   binding:"required"`
    PayType    float64 `form:"payType"    json:"payType"    binding:"required"`
    PayMessage string  `form:"payMessage" json:"payMessage" binding:"required"`
}
```

响应: 是否支付成功

### /fee/rate/distribute

Method: GET

描述: 查询机器节点独占费率

参数: 无

响应:

```json
{
    "36CPU": resp.Rate36CPU,
    "4GPU":  resp.Rate4GPU,
    "8GPU":  resp.Rate8GPU,
}
```

### /fee/usage/week

Method: GET

描述: 分页查询机器节点周账单

参数:

```typescript
pageIndex: number;
pageSize: number;
startDateMilliUnix: number;
endDateMilliUnix: number;
```

响应: 分页查询的结果

### /fee/usage/group/week

Method: GET

描述: 按照组 ID 进行分组，并分页查询某个组的账单信息

参数:

```typescript
pageIndex: number;
pageSize: number;
payFlag: boolean;
```

响应: 分页查询的结果

### /fee/usage/group/bill

Method: PUT

描述: 通过用户组 ID 支付用户组的机器节点机时未支付账单

参数:

```go
// PayGroupNodeUsageBillParam 支付用户组机器节点时长账单参数
type PayGroupNodeUsageBillParam struct {
    UserGroupID int     `form:"userGroupID" json:"userGroupID" binding:"required"`
    PayType     float64 `form:"payType"     json:"payType"     binding:"required"`
    PayMessage  string  `form:"payMessage"  json:"payMessage"`
    NeedFee     float64 `form:"needFee"     json:"needFee"     binding:"required"`
}
```

响应: 支付成功的账单数目

### /fee/rate/usage

Method: GET

描述: 查询机器时长费率信息

参数: 无

响应:

```go
map[string]float64{
    "cpu": 0,
    "gpu": 0,
}
```

### /fee/quota

Method: GET

描述: 分页查询机存储账单

参数:

```typescript
pageIndex: number;
pageSize: number;
```

响应: 分页查询的结果

### /fee/rate/quota

Method: GET

描述: 查询机器存储费率

参数: 无

响应: 费率信息

### /fee/quota/bill

Method: PUT

描述: 支付机器存储账单

参数:

```protobuf
// PayNodeQuotaBillParam 支付机器存储账单参数
type PayNodeQuotaBillParam struct {
    BillID     int     `form:"billID"     json:"billID"     binding:"required"`
    PayType    int     `form:"payType"    json:"payType"    binding:"required"`
    PayMessage string  `form:"payMessage" json:"payMessage"`
    PayMoney   float64 `form:"payMoney"   json:"payMoney"   binding:"required"`
}
```

响应: 是否支付成功

### /fee/rate/distribute

Method: PUT

描述: 修改机器分配费率

参数:

```go
// SetNodeDistributeFeeRateParam 设置机器分配费率参数
type SetNodeDistributeFeeRateParam struct {
    Rate36CPU float64 `form:"rate36CPU" json:"rate36CPU"`
    Rate4GPU  float64 `form:"rate4GPU"  json:"rate4GPU"`
    Rate8GPU  float64 `form:"rate8GPU"  json:"rate8GPU"`
}
```

响应: 是否操作成功

### /fee/rate/usage

Method: PUT

描述: 修改机器时长费率

参数:

```go
// SetNodeUsageFeeRateParam 设置机器节点使用机时费率参数
type SetNodeUsageFeeRateParam struct {
    Cpu float64 `form:"cpu" json:"cpu"`
    Gpu float64 `form:"gpu" json:"gpu"`
}
```

响应: 是否操作成功

## fss 控制器

### /fss/ping

Method: GET

描述：进行 group 服务的 ping 测试

参数：无

响应：请求 ID、PONG

### /fss/file

Method: POST

描述: 单个文件上传接口

参数: file 文件

响应: 新的文件名称

## award 控制器

### /award/ping

Method: GET

描述：进行 group 服务的 ping 测试

参数：无

响应：请求 ID、PONG

### /award/paper

Method: POST

描述: 创建论文奖励申请

参数:

```go
// CreatePaperAwardApplyParam 创建论文奖励申请参数
type CreatePaperAwardApplyParam struct {
    Title               string `form:"title"               json:"title"               binding:"required"`
    Category            string `form:"category"            json:"category"            binding:"required"`
    Partition           string `form:"partition"           json:"partition"           binding:"required"`
    FirstPageImageName  string `form:"firstPageImageName"  json:"firstPageImageName"  binding:"required"`
    ThanksPageImageName string `form:"thanksPageImageName" json:"thanksPageImageName" binding:"required"`
    RemarkMessage       string `form:"remarkMessage"       json:"remarkMessage"`
}
```

响应: 创建的申请记录的 ID

### /award/paper

Method: GET

描述: 分页查询论文奖励申请记录

请求参数: 标准的分页参数

响应参数: 查询的信息

### /award/paper

Method: PUT

描述: 审核论文奖励申请

参数:

```go
// CheckPaperApplyParam 审核论文奖励申请参数
type CheckPaperApplyParam struct {
    ID           int     `form:"id"           json:"id"           binding:"required"`
    CheckMoney   float64 `form:"checkMoney"   json:"checkMoney"`
    CheckMessage string  `form:"checkMessage" json:"checkMessage"`
    Accept       bool    `form:"accept"       json:"accept"`
}
```

响应参数: 是否成功

### /award/technology

描述: 创建科技奖励申请

参数:

```go
// CreateTechnologyAwardApplyParam 创建科技奖励申请参数
type CreateTechnologyAwardApplyParam struct {
    ProjectID      int    `form:"projectID"      json:"projectID"      binding:"required"`
    PrizeLevel     string `form:"prizeLevel"     json:"prizeLevel"     binding:"required"`
    PrizeImageName string `form:"prizeImageName" json:"prizeImageName" binding:"required"`
    RemarkMessage  string `form:"remarkMessage"  json:"remarkMessage"`
}
```

响应: 创建的记录的 ID

### /award/technology

Method: GET

描述: 分页查询科技奖励申请记录

请求参数: 标准的分页参数

响应参数: 查询的信息

### /award/technology

Method: PUT

描述: 审核科技奖励申请

参数:

```go
// CheckTechnologyApplyParam 审核科技奖励申请参数
type CheckTechnologyApplyParam struct {
    ID           int     `form:"id"           json:"id"           binding:"required"`
    CheckMoney   float64 `form:"checkMoney"   json:"checkMoney"`
    CheckMessage string  `form:"checkMessage" json:"checkMessage"`
    Accept       bool    `form:"accept"       json:"accept"`
}
```

响应参数: 是否成功
