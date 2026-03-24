# Go User Service

基于 Go + Gin + GORM + Mysql 开发的用户中心RESTful API 服务。


## 技术栈

- **框架**: Gin
- **ORM**: GORM
- **配置管理**: Viper
- **数据库**: Mysql
- **认证**: JWT

## 功能列表
- [x] 用户注册
- [x] 用户登录
- [x] 获取用户信息
- [ ] JWT认证中间件

## 项目结构
go_code/
├── main.go # 入口文件
├── config/ # 配置文件
├── controller/ # 控制器层
├── service/ # 业务逻辑层
├── model/ # 数据模型层
├── middleware/ # 中间件
└── go.mod # 依赖管理