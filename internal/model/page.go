package model

import "html/template"

type NavItem struct {
	Label string
	Href  string
}

type SiteInfo struct {
	Email       string
	Description string
	SiteName    string
	SiteURL     string
	AuthorName  string
	AuthorDesc  string
}

type Skill struct {
	Title string
	Desc  string
	Icon  template.HTML
}

var NavItems = []NavItem{
	{Label: "首页", Href: "/"},
	{Label: "关于", Href: "/about"},
}

var Info = SiteInfo{
	Email:       "m@devyunmu.com",
	Description: "云沐 (Yun Mu) — AI 原生开发工程师 / 云原生后端架构师，5年+ 经验，深耕 Spring Cloud 微服务架构。具备全栈开发与性能调优能力，系统性整合 Cursor + OpenCode 构建 AI 原生开发工作流。后端擅长微服务与数据库优化，前端可独立交付小程序。",
	SiteName:    "Dev云沐",
	SiteURL:     "https://devyunmu.com",
	AuthorName:  "云沐",
	AuthorDesc:  "AI 原生开发工程师 / 云原生后端架构师 · 全栈交付 · 效能革新",
}

var Skills = []Skill{
	{
		Title: "Java / Spring Boot",
		Desc:  "高性能后端服务开发，熟练运用 Spring Boot 与并发原语",
		Icon:  template.HTML(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2L2 7l10 5 10-5-10-5z"/><path d="M2 17l10 5 10-5"/><path d="M2 12l10 5 10-5"/></svg>`),
	},
	{
		Title: "Spring Cloud / 微服务",
		Desc:  "Spring Cloud Alibaba、Nacos 服务发现、RabbitMQ 消息队列、XXL-JOB 分布式调度",
		Icon:  template.HTML(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06A1.65 1.65 0 0 0 4.68 15a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06A1.65 1.65 0 0 0 9 4.68a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06A1.65 1.65 0 0 0 15 4.68a1.65 1.65 0 0 0-1 1.51V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/></svg>`),
	},
	{
		Title: "数据库 / 缓存",
		Desc:  "MySQL 关系数据库、Redis 缓存、ClickHouse / InfluxDB 时序数据库、Elasticsearch 搜索",
		Icon:  template.HTML(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3"/><path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"/></svg>`),
	},
	{
		Title: "DevOps / 容器化",
		Desc:  "Docker 容器化、Kubernetes 编排、Jenkins / GitLab CI 自动化部署",
		Icon:  template.HTML(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"/></svg>`),
	},
	{
		Title: "前端 / 小程序",
		Desc:  "Vue 前端开发、UniApp 跨端开发、微信小程序独立交付",
		Icon:  template.HTML(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="3" width="20" height="14" rx="2" ry="2"/><line x1="8" y1="21" x2="16" y2="21"/><line x1="12" y1="17" x2="12" y2="21"/></svg>`),
	},
	{
		Title: "性能优化 / 架构",
		Desc:  "大数据量接口性能诊断与调优，独立负责完整项目从 0 到 1 交付",
		Icon:  template.HTML(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polygon points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"/></svg>`),
	},
	{
		Title: "AI 原生开发工作流",
		Desc:  "系统性整合 Cursor IDE 与 OpenCode 平台，构建从需求分析到部署的 AI 协同开发流程，独立产品交付效率提升超 40%",
		Icon:  template.HTML(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2zm0 18a8 8 0 1 1 8-8 8 8 0 0 1-8 8z"/><path d="M12 6v6l4 2"/></svg>`),
	},
}
