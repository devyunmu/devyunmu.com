# Draft: 网站重构

## 当前项目状态

- **技术栈**: Go 1.26 + Gin v1.10.0 + html/template
- **页面**: 首页 (index.html) + 关于页 (about.html) + Health 端点
- **模板**: 6个模板文件 (base + 3 partials + 2 pages)
- **CSS**: 单文件 ~267行，BEM-ish 命名
- **JS**: 无 - 所有交互纯 CSS
- **中间件链**: Gzip → Cache → Security
- **数据层**: 无数据库，无外部 API，纯静态站点
- **测试**: 无 (_test.go 不存在)
- **构建**: go mod vendor + go build -mod=vendor (中国网络适配)
- **部署**: Docker 多阶段构建 → scratch 镜像

## Open Questions（待确认）

1. 重构的目标是什么？视觉 redesign？代码结构重组？还是两者都有？
2. 是否要保持"零 JS"的约束？
3. 是否有新增页面的需求？（如博客、作品集等）
4. CSS 是否考虑模块化/拆分？
5. 是否要引入 Tailwind 或其他 CSS 框架？
6. 是否需要引入测试？

## Final Decisions（已确认）
- **重构范围**: 视觉翻新 + 代码重构 + 工程化改进
- **功能扩展**: 本次不做（无博客、无作品集页）
- **暗色模式**: 纳入，作为视觉翻新的一部分（纯 CSS）
- **CSS 组织**: 单文件进化（内部用 CSS 变量/分区组织）
- **测试**: 引入 Go 单元测试（handler + middleware）
- **技术约束**: 保持 Gin + html/template；零 JS；无数据库；无外部 API
- **中国网络**: 保持 vendored 依赖

## User Decisions（已确认）
- **重构范围**: 全部包含（视觉 + 代码结构 + 功能扩展 + 工程化）
- **视觉方向**: 听我建议，需要先做调研
- **技术约束**: 保持 Gin + html/template；保持零 JS（推荐，用户已采纳）
- **功能扩展**: 没有具体想法，用户希望我给建议
- **时间规划**: 不着急，先出方案

## 调研发现

### 视觉设计方向
- 现代零 JS 个人站趋势: 干净排版、玻璃态/毛玻璃效果、Bento-grid 布局
- 深色模式: 纯 CSS 方案（`prefers-color-scheme` + CSS 自定义属性）
- 无框架 CSS: CSS Grid、Flexbox、Container Queries、`clamp()` 流式排版
- 动效: CSS scroll-driven animations、hover 过渡

### Go 项目结构最佳实践
- cmd/internal 布局是行业标准
- 推荐分层: handler → service → repository（尽管本站无数据库）
- 路由独立文件管理

### 博客功能方案
- goldmark（Go 的 CommonMark 解析器）实现 Markdown → HTML
- `embed.FS` 嵌入模板/静态资源
- 内容以 .md 文件管理，无需数据库
- 编译时 / 运行时渲染两种路线均可

## User Decisions（已确认）
- **重构范围**: 全部包含（视觉 + 代码结构 + 功能扩展 + 工程化）
- **视觉方向**: 听我建议，需要先做调研
- **技术约束**: 保持 Gin + html/template；保持零 JS（推荐，用户已采纳）
- **功能扩展**: 没有具体想法，用户希望我给建议
- **时间规划**: 不着急，先出方案
