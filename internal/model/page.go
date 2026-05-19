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
	Description: "云沐 (Yun Mu) — 全栈开发工程师 / 云原生后端架构师，5年+ 经验，深耕 Spring Cloud 微服务架构。具备全栈开发与性能调优能力，系统性整合 Cursor + OpenCode 构建 AI 辅助开发工作流。后端擅长微服务与数据库优化，前端可独立交付小程序。",
	SiteName:    "Dev云沐",
	SiteURL:     "https://devyunmu.com",
	AuthorName:  "云沐",
	AuthorDesc:  "全栈开发工程师 / 云原生后端架构师 · 深度交付 · 效能革新",
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
		Title: "AI 辅助开发工作流",
		Desc:  "系统性整合 Cursor IDE 与 OpenCode 平台，构建从需求分析到部署的 AI 协同开发流程，独立产品交付效率提升超 40%",
		Icon:  template.HTML(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2zm0 18a8 8 0 1 1 8-8 8 8 0 0 1-8 8z"/><path d="M12 6v6l4 2"/></svg>`),
	},
}

type CaseStudy struct {
	Slug           string
	Title          string
	Subtitle       string
	Tag            string
	TagColor       string
	Summary        string
	Challenge      string
	Approach       string
	Result         string
	TechStack      []string
	Highlight      string
	HighlightColor string
}

var CaseStudies = []CaseStudy{
	{
		Slug:           "performance-optimization",
		Title:          "时序数据查询性能优化",
		Subtitle:       "微电网物联网运营中台",
		Tag:            "性能调优",
		TagColor:       "emerald",
		Summary:        "基于 InfluxDB 时序数据库的高频电量数据采集系统，针对单月 2076 万条数据的查询瓶颈，通过分批并行查询结合本地缓存，将响应时间从 60 秒骤降至 200 毫秒。",
		Challenge:      "电力物联网平台需支撑 1 万+ 测点的高频数据采集（日采集 96 点），单月数据量达 2076 万条。历史数据查询接口响应时间长达 60 秒，严重影响业务报表生成和实时监控体验。",
		Approach:       "采用 Stopwatch 工具精准定位查询瓶颈，设计分批并行查询策略降低单次查询数据量，引入线程池管理控制并发度，配合本地缓存层缓存热点数据，避免重复查询数据库。",
		Result:         "响应时间从 60 秒降至 200 毫秒，性能提升 300 倍。系统可稳定支撑每日近亿级数据点的高并发读写，为业务提供了秒级报表生成能力。",
		TechStack:      []string{"Spring Boot", "InfluxDB", "线程池", "本地缓存"},
		Highlight:      "60s → 200ms",
		HighlightColor: "emerald",
	},
	{
		Slug:           "low-latency-architecture",
		Title:          "实时交易低延迟架构",
		Subtitle:       "电力现货市场决策系统",
		Tag:            "系统架构",
		TagColor:       "orange",
		Summary:        "针对电力现货市场 15 分钟粒度（日 96 点）的高频交易决策需求，设计 Caffeine + Redis + DB 三级缓存架构，实现毫秒级数据处理和决策输出。",
		Challenge:      "电力现货市场要求日前市场（D-1 日）和实时市场（T-15 分钟）的快速决策能力。节点电价、新能源出力预测等数据高频变化，传统数据库查询无法满足毫秒级响应要求。",
		Approach:       "设计三级缓存架构：Caffeine 本地缓存优先响应热点数据，Redis 缓存接口返回结果免除数据库查询开销，MySQL 兜底保证数据一致性。Kafka 消费者同步刷新缓存，采用预加载策略提前准备核心页面数据。",
		Result:         "核心接口响应时间从数百毫秒降至 10 毫秒以内，支撑现货市场高频交易决策的实时性需求，缓存命中率达 95%+，数据库压力降低 80%。",
		TechStack:      []string{"Caffeine", "Redis", "Kafka", "Spring Boot"},
		Highlight:      "毫秒级决策",
		HighlightColor: "orange",
	},
	{
		Slug:           "fullstack-delivery",
		Title:          "独立产品全栈交付",
		Subtitle:       "沐晴记账小程序",
		Tag:            "独立产品",
		TagColor:       "purple",
		Summary:        "从 0 到 1 独立完成产品设计、UI 实现、后端开发到部署上线全流程。重度引入 AI Agent 辅助开发（Cursor + OpenCode 双引擎），开发效率提升约 40%。",
		Challenge:      "需要独立交付一款完整的微信小程序，涵盖用户体系、账单记录、分类统计、数据可视化等核心功能，且要求代码结构清晰、可维护性强，具备持续迭代能力。",
		Approach:       "使用小程序原生框架 + 微信云开发，独立设计产品原型和交互逻辑。引入 Cursor AI Agent 作为主力开发工具，系统性整合 OpenCode 平台辅助代码审查和优化。采用组件化开发模式，确保代码复用率和可维护性。",
		Result:         "产品成功上线并稳定运行，核心功能完整可用。AI 辅助开发使整体效率提升 40%，代码质量达到生产标准，持续迭代中。",
		TechStack:      []string{"小程序原生", "微信云开发", "Cursor", "OpenCode"},
		Highlight:      "0→1 独立交付",
		HighlightColor: "purple",
	},
}

type DesignChallenge struct {
	Slug           string
	Title          string
	Subtitle       string
	Tag            string
	TagColor       string
	Summary        string
	Challenge      string
	Approach       string
	Result         string
	TechStack      []string
	Highlight      string
	HighlightColor string
}

var DesignChallenges = []DesignChallenge{
	{
		Slug:           "data-migration",
		Title:          "百亿行级数据迁移",
		Subtitle:       "电力交易辅助决策系统",
		Tag:            "架构设计推演",
		TagColor:       "blue",
		Summary:        "基于真实业务场景的架构设计推演：海量时序数据累计超百亿行、单表 TB 级，如何设计分阶段迁移方案（历史全量 + 增量同步 + 灰度切换），实现业务零停机、数据零丢失。",
		Challenge:      "系统使用 ClickHouse 存储百亿行时序数据（负荷、电价、新能源出力等），需迁移至 MySQL 8.0 以满足事务支持和复杂关联查询需求。迁移窗口期短，业务不能停，数据一致性要求极高。",
		Approach:       "划分数据预处理、全量迁移、增量同步、灰度切换、旧系统下线五个阶段。历史数据按月分片并行导出，使用 LOAD DATA INFILE 高效导入；增量阶段采用单向同步避免双向写入冲突；切换阶段通过 Nginx/Nacos 灰度流量控制。",
		Result:         "3 年全量数据无损迁移，业务零停机。迁移后查询性能满足业务需求，成功将时序数据纳入 MySQL 事务体系，支持复杂关联分析和行级更新。",
		TechStack:      []string{"ClickHouse", "MySQL 8.0", "DataX", "Spring Boot"},
		Highlight:      "百亿行 · 零丢失 · 零停机",
		HighlightColor: "blue",
	},
}
