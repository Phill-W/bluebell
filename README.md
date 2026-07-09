# bluebell

基于 Go + Gin 的论坛社区类 Web API 项目。提供用户认证、社区分类、帖子发布与列表、点赞投票等功能，包含完整的前后端与 Docker 部署方案。

## 技术栈

| 分类 | 技术 | 说明 |
|------|------|------|
| Web 框架 | [Gin](https://github.com/gin-gonic/gin) | HTTP 路由与中间件 |
| 数据库 | MySQL 8.0 + [sqlx](https://github.com/jmoiron/sqlx) | 关系型存储，sqlx 增强 database/sql |
| 缓存 | Redis 5.0 + [go-redis](https://github.com/go-redis/redis) | 投票排行榜（ZSet）、会话缓存 |
| 认证 | [jwt-go](https://github.com/dgrijalva/jwt-go) | JWT Token 鉴权 |
| ID 生成 | [snowflake](https://github.com/bwmarrin/snowflake) | 雪花算法生成分布式唯一 ID |
| 配置 | [viper](https://github.com/spf13/viper) | 读取 YAML 配置，支持热加载 |
| 日志 | [zap](https://go.uber.org/zap) + [lumberjack](https://github.com/natefinch/lumberjack) | 高性能日志 + 按大小/时间切割 |
| 参数校验 | [validator/v10](https://github.com/go-playground/validator) | 结构体标签校验，内置中文翻译 |
| 接口文档 | [swaggo/swag](https://github.com/swaggo/swag) | 根据注解自动生成 Swagger 文档 |
| 性能分析 | [pprof](https://github.com/gin-contrib/pprof) | 内建性能剖析端点 |

## 功能特性

- **用户系统**：注册、登录（JWT 签发与校验）
- **社区模块**：社区列表、社区详情
- **帖子模块**：发帖（需登录）、帖子列表（按时间/分数排序）、帖子详情
- **投票模块**：基于 Redis ZSet 实现的点赞投票与排行榜
- **接口文档**：Swagger UI 在线查看与调试
- **性能剖析**：内置 pprof，支持 CPU/内存/Goroutine 分析
- **前端页面**：Vue 打包的 SPA，随服务一起部署

## 项目结构

```
bluebell/
├── main.go                  # 程序入口，加载配置并初始化各组件
├── conf/
│   ├── config.dev.yaml      # 本地开发配置（host: 127.0.0.1, mode: debug）
│   └── config.docker.yaml   # Docker 部署配置（host: 服务名, mode: release）
├── controller/              # 控制层：解析请求、参数校验、调用 logic、返回响应
├── logic/                   # 业务逻辑层：核心业务处理
├── dao/                     # 数据访问层
│   ├── mysql/              #   MySQL 操作（CRUD、密码加密）
│   └── redis/              #   Redis 操作（投票、缓存）
├── middlewares/             # 中间件：JWT 认证、限流
├── models/                  # 数据模型、请求参数结构、建表 SQL
├── pkg/                     # 工具包
│   ├── jwt/                #   JWT 生成与解析
│   └── snowflake/          #   雪花算法 ID 生成器
├── router/                  # 路由注册
├── setting/                 # 配置加载（基于 viper）
├── logger/                  # 日志初始化（zap + lumberjack）
├── docs/                    # Swagger 文档（swag init 自动生成）
├── static/                  # 前端静态资源（Vue 打包产物）
├── templates/               # HTML 模板（入口 index.html）
├── Dockerfile               # 多阶段构建：golang:alpine 编译 + debian:bookworm-slim 运行
├── docker-compose.yml       # 编排 MySQL + Redis + bluebell_app
├── wait-for.sh              # 等待 MySQL/Redis 端口就绪后再启动应用
├── bluebell_user.sql        # 建表脚本：user 表 + 测试数据
├── bluebell_community.sql   # 建表脚本：community 表 + 4 条初始数据
├── bluebell_post.sql        # 建表脚本：post 表 + 15 条测试数据
├── Makefile                 # 构建/运行/格式化命令
├── go.mod
└── go.sum
```

## 快速启动

### 方式一：Docker Compose（推荐）

一键启动 MySQL + Redis + 应用，首次启动自动建库、建表并灌入测试数据：

```bash
docker compose up
```

启动完成后：

- 应用地址：<http://localhost:8888>
- Swagger 文档：<http://localhost:8888/swagger/index.html>
- pprof 性能分析：<http://localhost:8888/debug/pprof/>

> MySQL 容器通过 `/docker-entrypoint-initdb.d/` 在首次初始化时自动执行 `bluebell_user.sql`、`bluebell_community.sql`、`bluebell_post.sql`，无需手动导入。

**常用命令：**

```bash
docker compose up          # 前台启动（看日志）
docker compose up -d       # 后台启动
docker compose down        # 停止并移除容器（保留数据）
docker compose down -v     # 停止并清空数据卷（重置数据库，下次启动重新初始化）
```

### 方式二：本地运行

适用于开发调试。需先在本地启动 MySQL（3306）和 Redis（6379）。

1. **配置文件**：本地开发使用 `conf/config.dev.yaml`（已预置 `host: 127.0.0.1`、`mode: debug`），无需修改。如需调整密码或端口，编辑该文件即可。

2. **建库建表**：在 MySQL 中执行：

   ```sql
   CREATE DATABASE IF NOT EXISTS bluebell;
   USE bluebell;
   SOURCE bluebell_user.sql;
   SOURCE bluebell_community.sql;
   SOURCE bluebell_post.sql;
   ```

3. **启动应用**：

   ```bash
   make run
   # 或
   go run main.go conf/config.dev.yaml
   ```

启动后监听 `conf/config.dev.yaml` 中 `port` 指定的端口（默认 8084）。

## 接口列表

| 方法 | 路径 | 说明 | 鉴权 |
|------|------|------|------|
| POST | `/api/v1/signup` | 用户注册 | 否 |
| POST | `/api/v1/login` | 用户登录 | 否 |
| GET | `/api/v1/community` | 社区列表 | 否 |
| GET | `/api/v1/community/:id` | 社区详情 | 否 |
| GET | `/api/v1/posts` | 帖子列表（按时间/分数） | 否 |
| GET | `/api/v1/posts2` | 帖子列表（新版，带作者与社区信息） | 否 |
| GET | `/api/v1/post/:id` | 帖子详情 | 否 |
| POST | `/api/v1/post` | 发布帖子 | 是 |
| POST | `/api/v1/vote` | 投票 | 是 |

> 需鉴权的接口请在请求头携带 `Authorization: <token>`，token 通过登录接口获取。

## 账号说明

数据库预置的测试账号密码加密方式与当前代码（MD5）不匹配，无法直接登录。**首次使用请注册新账号**：

```bash
# 注册账号
curl -X POST http://localhost:8888/api/v1/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456","re_password":"123456"}'

# 登录获取 token
curl -X POST http://localhost:8888/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'
```

## 配置说明

项目包含两份配置文件，按运行环境选用：

| 文件 | 用途 | mode | mysql/redis host |
|------|------|------|------------------|
| `conf/config.dev.yaml` | 本地开发 | debug | 127.0.0.1 |
| `conf/config.docker.yaml` | Docker 部署 | release | 服务名（mysql8019 / redis507） |

`make run` 默认使用 `config.dev.yaml`，`docker compose` 使用 `config.docker.yaml`，两者互不干扰。也可手动指定：`go run main.go conf/config.dev.yaml`。

配置字段说明（两份文件结构一致）：

```yaml
name: "bluebell"       # 应用名
mode: "release"        # 运行模式：debug / release
port: 8084             # 监听端口
version: "v0.0.1"
start_time: "2020-07-01"
machine_id: 1          # 雪花算法机器标识（分布式部署时各节点需唯一）

auth:
  jwt_expire: 8760     # JWT 有效期（小时）

log:
  level: "info"        # 日志级别
  filename: "web_app.log"
  max_size: 200        # 单文件最大 MB
  max_age: 30          # 保留天数
  max_backups: 7       # 保留备份数

mysql:
  host: mysql8019      # docker 版用服务名；dev 版为 127.0.0.1
  port: 3306
  user: root
  password: 123456
  dbname: bluebell
  max_open_conns: 200
  max_idle_conns: 50

redis:
  host: redis507       # docker 版用服务名；dev 版为 127.0.0.1
  port: 6379
  password: ""
  db: 0
  pool_size: 100
```

## 重新生成 Swagger 文档

修改 controller 中的注解后，重新生成文档：

```bash
swag init
```

生成结果在 `docs/` 目录，访问 <http://localhost:8084/swagger/index.html>（Docker 模式为 8888 端口）查看。

## Makefile 命令

| 命令 | 说明 |
|------|------|
| `make` | 格式化代码 + 编译生成二进制文件 |
| `make build` | 交叉编译 Linux amd64 二进制到 `bin/bluebell` |
| `make run` | 直接运行（`go run main.go conf/config.dev.yaml`） |
| `make gotool` | 执行 `go fmt` + `go vet` |
| `make clean` | 清理二进制文件 |
