# MaskLR-Go 项目结构

## masklr-go-auth/

### `cmd/`
- **server/**  
  - `main.go`  # 应用入口

### `internal/`  
应用核心逻辑（非公共代码）

- **config/**  
  - `config.go`  # 配置加载（env, JSON, .env）

- **db/**  
  - `mysql.go`  # 数据库连接（MySQL）

- **user/**  
  - `model.go`  # 用户结构体和数据库方法  
  - `service.go`  # 用户业务逻辑（注册、登录等）  
  - `handler.go`  # HTTP 请求处理（注册、登录）

- **middleware/**  
  - `auth.go`  # JWT 验证中间件

- **router/**  
  - `router.go`  # Gin 路由配置

- **util/**  
  - `hash.go`  # 密码哈希加密  
  - `token.go`  # JWT 生成与验证

### `api/`  
Swagger 文档或 API 接口定义（可选）

### `sql/`
- `init.sql`  # MySQL 用户表建表语句（已建好可忽略）

### `go.mod`  
Go 模块文件

### `go.sum`  
依赖校验
