# Dev云沐 Portfolio

个人作品集网站，使用 Go + Gin 构建的服务端渲染应用。零数据库、零 CMS，所有内容硬编码在 Go 结构体中，开箱即用。

## 特性

- **服务端渲染** — Go `html/template` 模板引擎，首屏即完整页面
- **深色主题** — Tailwind CSS CDN 驱动，品牌色 `#06b6d4`（青色）
- **SEO 优化** — 内嵌 JSON-LD 结构化数据（Person Schema）
- **安全优先** — HSTS（2年）、严格 CSP、Permissions-Policy 全禁用
- **响应式设计** — 移动端适配，`prefers-reduced-motion` 无障碍支持
- **Docker 部署** — 纯 scratch 镜像，不含 Go 运行时，镜像极小（~15MB）
- **零依赖运行** — 编译后单二进制文件，无需运行时环境

## 快速开始

### 本地开发

```bash
# 克隆仓库
git clone <repository-url>
cd devyunmu.com

# 运行（默认端口 8080）
PORT=8080 go run ./cmd/server
```

访问 `http://localhost:8080`

### Docker

```bash
# 本地编译二进制
go build -o devyunmu.com ./cmd/server

# 构建镜像（无需 golang 编译镜像）
docker build -t devyunmu.com .

# 运行容器
docker run -p 8080:8080 devyunmu.com
```

### 部署到远程服务器

```bash
# 本地：编译二进制 → 构建镜像 → 导出
go build -o devyunmu.com ./cmd/server
docker build -t devyunmu.com .
docker save devyunmu.com:latest | gzip > /tmp/devyunmu.tar.gz
scp /tmp/devyunmu.tar.gz root@<IP>:/tmp/

# 远程：加载镜像 → 替换容器
ssh root@<IP> "docker load < /tmp/devyunmu.tar.gz && \
  docker stop devyunmu && docker rm devyunmu && \
  docker run -d --name devyunmu -p 8080:8080 --restart unless-stopped devyunmu.com:latest && \
  rm /tmp/devyunmu.tar.gz"
```

## 项目结构

```
./
├── cmd/server/         # 入口：main.go 启动 Gin、注册中间件与路由
├── config/             # 环境变量配置（PORT）
├── internal/
│   ├── handler/        # 路由处理器：Home、About、Health
│   ├── middleware/     # 中间件栈：Gzip → Security → ErrorHandler → Cache
│   └── model/          # 数据模型 + 硬编码站点内容
├── templates/          # Go HTML 模板（_ 前缀为局部模板）
├── static/             # 静态资源：SVG 图标、robots.txt、sitemap.xml
├── Dockerfile          # 纯 scratch 部署镜像
└── go.mod              # 模块路径：devyunmu.com，Go 1.26
```

## 技术栈

| 类别 | 技术 |
|------|------|
| 语言 | Go 1.26 |
| 框架 | Gin v1.10 |
| 模板 | Go `html/template` |
| 样式 | Tailwind CSS CDN（深色主题，品牌色 `#06b6d4`） |
| 部署 | Docker scratch 镜像（本地编译 → 纯部署镜像） |
| 依赖管理 | Go Modules |

## 路由

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/health` | 健康检查，返回 204 No Content |
| `GET` | `/` | 首页：Hero、技能、项目、联系信息 |
| `GET` | `/about` | 关于页，含个人简介元数据 |
| `GET` | `/static/*` | 静态文件服务（图片、robots.txt、sitemap.xml） |

## 中间件栈

中间件按以下顺序执行（顺序很重要）：

1. **Gzip** — 响应压缩
2. **Security** — 安全头（HSTS、CSP、X-Frame-Options 等）
3. **ErrorHandler** — 统一错误处理（结构化 `HTTPError` + `ErrResponse`）
4. **Cache** — 缓存控制（静态资源 1 年，页面 1 小时）

## 配置

通过环境变量配置：

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `PORT` | `8080` | 服务监听端口 |

```bash
# 自定义端口
PORT=3000 go run ./cmd/server
```

站点内容（技能、简介、导航）硬编码在 `internal/model/page.go` 中，修改后重新构建即可。

## 构建与测试

```bash
# 运行测试
go test ./...

# 编译二进制
CGO_ENABLED=0 go build -o devyunmu.com ./cmd/server
```

## 其他说明

- 所有站点内容（技能、描述、个人简介）均硬编码在 `internal/model/page.go`，无 CMS、无数据库
- Git 提交格式：`type: description`（例如 `fix: resolve all project issues`）

## License

MIT
