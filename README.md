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
- [x] JWT认证中间件

## 项目结构
```
go_code/
├── main.go           # 入口文件
├── config/           # 配置文件
├── handler/          # 数据处理层
├── service/          # 业务逻辑层
├── model/            # 数据模型层
├── middleware/       # 中间件
├── pkg/              # 公共工具包
└── go.mod            # 依赖管理
```

## 快速启动

### 1. 配置数据库

修改`config.yaml`参数，修改数据库连接信息：

```yaml
database:
  host: 127.0.0.1
  port: 3306
  user: root
  password: your_password
  name: user_service
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 运行

```bash
go run main.go
```

### 4. API 接口

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| 注册 | POST | /api/users/register | 用户注册 |
| 登录 | POST | /api/users/login | 用户登录 |
| 获取信息 | GET | /api/users/info | 获取当前用户信息 |