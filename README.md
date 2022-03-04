# hpcmannager

高性能计算平台管理系统

为基于go语言的微服务系统，主要使用go-micro框架进行搭建

## 服务划分

* 机器管理: 处理包机申请、存储申请、机时管理查看以及存储情况查看
* 费用管理: 处理账单的生成以及统计以及缴费
* 用户管理: 处理用户和用户组的增删改查
* 权限服务: 处理用户权限、系统权限方面的信息(操作鉴权功能通过工具包的形式提供)
* 奖励管理: 处理论文奖励申请和科技成果奖励申请
* 作业调度服务:对接上游作业调度系统，提供接口供其他服务获取机时信息、存储信息，以及机器用户、用户组的操作和机器节点的操作
* 项目管理服务: 项目创建等信息的管理
* 网关服务:初步鉴权、协议转换、接口聚合
* 代理服务: 使用Nginx暴露网关接口以及静态资源的请求
* 静态资源服务: 负责处理静态资源文件的存储以及下载
