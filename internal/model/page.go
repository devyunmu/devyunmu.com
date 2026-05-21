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
	ICP         string
}

type Skill struct {
	Title string
	Desc  string
	Icon  template.HTML
}

var NavItems = []NavItem{
	{Label: "首页", Href: "/"},
	{Label: "关于", Href: "/about"},
	{Label: "文档", Href: "/tutorials"},
}

var Info = SiteInfo{
	Email:       "m@devyunmu.com",
	Description: "云沐 (Yun Mu) — 全栈开发工程师 / 云原生后端架构师，5年+ 经验，深耕 Spring Cloud 微服务架构。具备全栈开发与性能调优能力，系统性整合 Cursor + OpenCode 构建 AI 辅助开发工作流。后端擅长微服务与数据库优化，前端可独立交付小程序。",
	SiteName:    "Dev云沐",
	SiteURL:     "https://devyunmu.com",
	AuthorName:  "云沐",
	AuthorDesc:  "全栈开发工程师 / 云原生后端架构师 · 深度交付 · 效能革新",
	ICP:         "湘ICP备2024062354号-3",
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

var DesignChallenges = []DesignChallenge{}

type Doc struct {
	Slug        string
	Title       string
	Subtitle    string
	Category    string
	CategoryColor string
	Summary     string
	Content     template.HTML
	TechStack   []string
	Difficulty  string
	ReadTime    string
	CreatedAt   string
	UpdatedAt   string
}

var Docs = []Doc{
	{
		Slug:          "java-spring-boot-quickstart",
		Title:         "Spring Boot 快速入门指南",
		Subtitle:      "从零构建生产级 Java 后端服务",
		Category:      "Spring Boot",
		CategoryColor: "emerald",
		Summary:       "手把手带你从零搭建 Spring Boot 项目，涵盖依赖管理、配置分层、RESTful API 设计、统一异常处理与日志规范，适合初学者快速上手。",
		TechStack:     []string{"Java 17", "Spring Boot 3.x", "Maven", "JUnit 5"},
		Difficulty:    "入门",
		ReadTime:      "25 分钟",
		CreatedAt:     "2026-05-21",
		UpdatedAt:     "2026-05-21",
	},
	{
		Slug:          "spring-cloud-microservices",
		Title:         "Spring Cloud 微服务架构实战",
		Subtitle:      "基于 Spring Cloud Alibaba 构建企业级微服务平台",
		Category:      "微服务",
		CategoryColor: "blue",
		Summary:       "深入讲解 Nacos 服务发现与配置中心、OpenFeign 声明式调用、Gateway 网关路由、Sentinel 流量控制等核心组件，配合实战案例完整落地微服务体系。",
		TechStack:     []string{"Spring Cloud", "Nacos", "Gateway", "Sentinel", "OpenFeign"},
		Difficulty:    "进阶",
		ReadTime:      "45 分钟",
		CreatedAt:     "2026-05-21",
		UpdatedAt:     "2026-05-21",
	},
	{
		Slug:          "mysql-performance-tuning",
		Title:         "MySQL 性能优化实战手册",
		Subtitle:      "从慢查询到高并发，系统性调优方法论",
		Category:      "数据库",
		CategoryColor: "orange",
		Summary:       "涵盖索引设计原则、执行计划分析、慢查询定位、连接池调优、分库分表策略及读写分离实践，解决生产环境常见的数据库性能瓶颈。",
		TechStack:     []string{"MySQL 8.0", "MyBatis Plus", "ShardingSphere"},
		Difficulty:    "进阶",
		ReadTime:      "35 分钟",
		CreatedAt:     "2026-05-21",
		UpdatedAt:     "2026-05-21",
	},
	{
		Slug:          "redis-cache-patterns",
		Title:         "Redis 缓存设计与应用模式",
		Subtitle:      "高并发场景下的缓存最佳实践",
		Category:      "缓存",
		CategoryColor: "purple",
		Summary:       "详解缓存穿透、缓存击穿、缓存雪崩的解决方案，介绍布隆过滤器、分布式锁、缓存一致性等高级模式，结合 Spring Data Redis 提供完整代码示例。",
		TechStack:     []string{"Redis 7", "Spring Data Redis", "Redisson", "Lettuce"},
		Difficulty:    "中级",
		ReadTime:      "30 分钟",
		CreatedAt:     "2026-05-21",
		UpdatedAt:     "2026-05-21",
	},
	{
		Slug:          "jvm-deep-dive",
		Title:         "JVM 原理与调优深度解析",
		Subtitle:      "理解内存模型、垃圾回收与线上问题排查",
		Category:      "底层原理",
		CategoryColor: "yellow",
		Summary:       "从 JVM 内存结构出发，深入讲解 G1/ZGC 垃圾回收器原理、常见 OOM 场景分析、Arthas 线上诊断工具使用，以及 GC 日志分析与 JVM 参数调优策略。",
		TechStack:     []string{"Java 17", "JVM", "Arthas", "G1/ZGC"},
		Difficulty:    "高级",
		ReadTime:      "50 分钟",
		CreatedAt:     "2026-05-21",
		UpdatedAt:     "2026-05-21",
	},
	{
		Slug:          "java-concurrency-patterns",
		Title:         "Java 并发编程核心模式",
		Subtitle:      "线程池、锁机制与并发容器实战",
		Category:      "并发编程",
		CategoryColor: "cyan",
		Summary:       "系统梳理 Java 并发编程知识体系：从 synchronized/volatile 底层实现到 JUC 工具类使用，重点讲解线程池参数调优、CompletableFuture 异步编排及 Fork/Join 框架。",
		TechStack:     []string{"Java 17", "JUC", "CompletableFuture", "Fork/Join"},
		Difficulty:    "高级",
		ReadTime:      "40 分钟",
		CreatedAt:     "2026-05-21",
		UpdatedAt:     "2026-05-21",
	},
}
