# MaskLR-Go

masklr-go-auth/
├── cmd/
│   └── server/
│       └── main.go             # 应用入口
│
├── internal/                   # 应用核心逻辑（非公共代码）
│   ├── config/                 # 配置加载（env, JSON, .env）
│   │   └── config.go
│   │
│   ├── db/                     # 数据库连接（MySQL）
│   │   └── mysql.go
│   │
│   ├── user/                   # 用户模块（model + service）
│   │   ├── model.go            # 用户结构体和数据库方法
│   │   ├── service.go          # 用户业务逻辑（注册、登录等）
│   │   └── handler.go          # HTTP 请求处理（注册、登录）
│   │
│   ├── middleware/             # JWT 验证中间件
│   │   └── auth.go
│   │
│   ├── router/                 # Gin 路由配置
│   │   └── router.go
│   │
│   └── util/                   # 通用工具函数
│       ├── hash.go             # 密码哈希加密
│       └── token.go            # JWT 生成与验证
│
├── api/                        # Swagger 文档或 API 接口定义（可选）
│
├── sql/
│   └── init.sql                # MySQL 用户表建表语句（已建好可忽略）
│
├── go.mod                      # Go 模块文件
└── go.sum                      # 依赖校验
