# Draft: 网站重新设计

## Current Site Summary
- 技术栈：Go + Gin + Go html/template + 单文件 CSS
- 页面：首页 (/) + 关于 (/about) + 健康检查 (/health)
- 特性：零 JS，纯 CSS 交互，Docker 部署
- 风格：极简个人品牌页，蓝灰配色

## Requirements (confirmed)
- 全盘重做：视觉 + 内容 + 功能整体升级
- 风格方向：温暖/个人化，暖橙/琥珀色系配色方案
- 配色方案：白天模式 #FFFAF5 + #D97706/#F59E0B；深色模式 #0F0A06 + #F59E0B/#FBBF24
- 技术架构：保留 Go + Gin + html/template，保持零JS/轻量特性
- 页面结构：保持首页 + 关于两页结构
- 内容更新：个人理念更新，头像照片
- 深色模式：需要（prefers-color-scheme CSS 变量切换）
- 交互动效：适度（加载淡入、滚动渐入、悬浮微动效）
- 技能卡片：从 4 项扩展到 6 项（Go, Gin/K8s, 云原生, DevOps, 开源, 系统设计）
- 关于页内容：头像照片 + 个人理念 + 技能标签云
- 测试框架：需要基础测试框架配置
