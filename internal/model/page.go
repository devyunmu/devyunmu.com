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
	{Label: "文档", Href: "/docs"},
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
		Content: template.HTML(`<div class="space-y-8">
<h3 class="text-xl font-bold text-white">什么是 Spring Boot</h3>
<p class="text-slate-400 leading-relaxed">Spring Boot 是 Spring 框架的扩展，旨在简化新 Spring 应用的初始搭建和开发过程。它采用"约定优于配置"的理念，通过自动配置和起步依赖，让你可以在几分钟内创建一个可运行的、生产级的 Spring 应用，而无需编写大量的样板代码和 XML 配置。</p>

<h3 class="text-xl font-bold text-white mt-8">搭建开发环境</h3>
<p class="text-slate-400 leading-relaxed">在开始之前，请确保已安装以下工具：</p>
<ul class="list-disc list-inside text-slate-400 space-y-2 mt-2">
<li>JDK 17 或更高版本（推荐使用 Eclipse Temurin 或 Amazon Corretto）</li>
<li>Apache Maven 3.9+ 或 Gradle 8+</li>
<li>IntelliJ IDEA（社区版或 Ultimate 版均可）</li>
</ul>

<h3 class="text-xl font-bold text-white mt-8">创建第一个 Spring Boot 项目</h3>
<p class="text-slate-400 leading-relaxed">最简单的方式是使用 Spring Initializr（<a href="https://start.spring.io" class="text-brand-light hover:underline">https://start.spring.io</a>）。选择以下配置：</p>
<ul class="list-disc list-inside text-slate-400 space-y-2 mt-2">
<li>Project: Maven</li>
<li>Language: Java</li>
<li>Spring Boot: 3.2.x</li>
<li>Dependencies: Spring Web, Spring Data JPA, H2 Database</li>
</ul>

<h3 class="text-xl font-bold text-white mt-8">项目结构解析</h3>
<p class="text-slate-400 leading-relaxed">标准的 Spring Boot 项目采用分层架构，各层职责清晰：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">src/main/java/com/example/demo/
├── DemoApplication.java      启动类
├── config/                   配置类
├── controller/               控制器层（REST API）
├── service/                  业务逻辑层
│   └── impl/                 实现类
├── repository/               数据访问层
├── entity/                   实体类（ORM 映射）
├── dto/                      数据传输对象
├── exception/                自定义异常
└── util/                     工具类

src/main/resources/
├── application.yml           主配置文件
├── application-dev.yml       开发环境配置
├── application-prod.yml      生产环境配置
└── static/                   静态资源</code></pre>

<h3 class="text-xl font-bold text-white mt-8">核心概念：依赖注入</h3>
<p class="text-slate-400 leading-relaxed">Spring Boot 的核心是控制反转（IoC）容器。通过 <code>@Autowired</code> 或构造函数注入，容器会自动装配所需的 Bean：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@Service
public class UserService {
    private final UserRepository userRepository;

    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public User findById(Long id) {
        return userRepository.findById(id)
            .orElseThrow(() -> new ResourceNotFoundException("User not found"));
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">配置分层管理</h3>
<p class="text-slate-400 leading-relaxed">使用多环境配置文件实现不同环境的差异化管理：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300"># application.yml（公共配置）
spring:
  profiles:
    active: dev
  application:
    name: demo-service

---
# application-dev.yml（开发环境）
spring:
  datasource:
    url: jdbc:h2:mem:testdb
    driver-class-name: org.h2.Driver
  jpa:
    hibernate:
      ddl-auto: create-drop
    show-sql: true

---
# application-prod.yml（生产环境）
spring:
  datasource:
    url: jdbc:mysql://localhost:3306/prod_db
    username: ${DB_USERNAME}
    password: ${DB_PASSWORD}
  jpa:
    hibernate:
      ddl-auto: validate</code></pre>

<h3 class="text-xl font-bold text-white mt-8">统一异常处理</h3>
<p class="text-slate-400 leading-relaxed">使用 <code>@RestControllerAdvice</code> 实现全局异常拦截，保证 API 返回格式统一：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(ResourceNotFoundException.class)
    public ResponseEntity&lt;ApiResponse&lt;?&gt;&gt; handleNotFound(ResourceNotFoundException ex) {
        return ResponseEntity.status(HttpStatus.NOT_FOUND)
            .body(ApiResponse.error(ex.getMessage(), HttpStatus.NOT_FOUND.value()));
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity&lt;ApiResponse&lt;?&gt;&gt; handleValidation(MethodArgumentNotValidException ex) {
        String message = ex.getBindingResult().getFieldErrors().stream()
            .map(error -> error.getField() + ": " + error.getDefaultMessage())
            .collect(Collectors.joining(", "));
        return ResponseEntity.badRequest()
            .body(ApiResponse.error(message, HttpStatus.BAD_REQUEST.value()));
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">日志规范</h3>
<p class="text-slate-400 leading-relaxed">统一使用 SLF4J + Logback，避免直接调用 <code>System.out.println</code>：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@Slf4j
@RestController
@RequestMapping("/api/users")
public class UserController {

    @GetMapping("/{id}")
    public ApiResponse&lt;User&gt; getUser(@PathVariable Long id) {
        log.info("Fetching user with id: {}", id);
        User user = userService.findById(id);
        log.debug("User found: {}", user);
        return ApiResponse.success(user);
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">打包与运行</h3>
<p class="text-slate-400 leading-relaxed">Spring Boot 应用可以打包为可执行的 JAR 文件，内置 Tomcat，无需外部部署：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300"># 打包
mvn clean package -DskipTests

# 运行
java -jar target/demo-0.0.1-SNAPSHOT.jar

# 或直接使用 Maven 插件
mvn spring-boot:run</code></pre>
</div>`),
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
		Content: template.HTML(`<div class="space-y-8">
<h3 class="text-xl font-bold text-white">为什么需要微服务架构</h3>
<p class="text-slate-400 leading-relaxed">当单体应用规模不断扩大时，会面临代码耦合严重、团队并行开发困难、技术栈单一、部署风险高等问题。微服务架构将系统拆分为一组小型、自治的服务，每个服务围绕业务能力构建，独立部署、独立扩展，使用最适合的技术栈。</p>

<h3 class="text-xl font-bold text-white mt-8">服务注册与发现：Nacos</h3>
<p class="text-slate-400 leading-relaxed">Nacos 是阿里巴巴开源的服务发现、配置管理和服务管理平台。在微服务架构中，服务实例动态上下线，需要一个注册中心来维护服务地址的实时映射。</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// 服务提供者注册到 Nacos
@SpringBootApplication
@EnableDiscoveryClient
public class OrderServiceApplication {
    public static void main(String[] args) {
        SpringApplication.run(OrderServiceApplication.class, args);
    }
}

# application.yml
spring:
  application:
    name: order-service
  cloud:
    nacos:
      discovery:
        server-addr: 127.0.0.1:8848
        namespace: prod
        group: DEFAULT_GROUP</code></pre>

<h3 class="text-xl font-bold text-white mt-8">声明式服务调用：OpenFeign</h3>
<p class="text-slate-400 leading-relaxed">OpenFeign 让你可以用声明式接口调用远程服务，就像调用本地方法一样简洁：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@FeignClient(name = "user-service", fallbackFactory = UserClientFallbackFactory.class)
public interface UserClient {

    @GetMapping("/api/users/{id}")
    User getUserById(@PathVariable("id") Long id);

    @PostMapping("/api/users")
    User createUser(@RequestBody UserDTO user);
}

// 降级处理
@Component
@Slf4j
public class UserClientFallbackFactory implements FallbackFactory&lt;UserClient&gt; {
    @Override
    public UserClient create(Throwable cause) {
        log.error("User service fallback, reason: {}", cause.getMessage());
        return new UserClient() {
            @Override
            public User getUserById(Long id) {
                return User.builder().id(id).name("Unknown").build();
            }
            @Override
            public User createUser(UserDTO user) {
                throw new ServiceUnavailableException("User service is unavailable");
            }
        };
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">API 网关：Spring Cloud Gateway</h3>
<p class="text-slate-400 leading-relaxed">Gateway 作为系统的统一入口，负责路由转发、鉴权、限流、日志记录等横切关注点：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@Configuration
public class GatewayConfig {

    @Bean
    public RouteLocator customRouteLocator(RouteLocatorBuilder builder) {
        return builder.routes()
            .route("user-service", r -> r.path("/api/users/**")
                .filters(f -> f.stripPrefix(1)
                    .addRequestHeader("X-Gateway", "gateway-v1")
                    .circuitBreaker(config -> config
                        .setName("userCircuitBreaker")
                        .setFallbackUri("forward:/fallback")))
                .uri("lb://user-service"))
            .route("order-service", r -> r.path("/api/orders/**")
                .uri("lb://order-service"))
            .build();
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">流量控制与熔断：Sentinel</h3>
<p class="text-slate-400 leading-relaxed">Sentinel 提供流量控制、熔断降级、系统负载保护等功能。通过注解方式即可实现细粒度的限流：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@RestController
public class OrderController {

    @SentinelResource(
        value = "createOrder",
        blockHandler = "createOrderBlockHandler",
        fallback = "createOrderFallback"
    )
    @PostMapping("/orders")
    public Order createOrder(@RequestBody OrderDTO dto) {
        return orderService.create(dto);
    }

    public Order createOrderBlockHandler(OrderDTO dto, BlockException ex) {
        log.warn("Order creation blocked: {}", ex.getMessage());
        throw new RateLimitExceededException("Server is busy, please try again later");
    }

    public Order createOrderFallback(OrderDTO dto, Throwable ex) {
        log.error("Order creation failed: {}", ex.getMessage());
        return Order.builder().status(OrderStatus.PENDING).build();
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">分布式事务：Seata</h3>
<p class="text-slate-400 leading-relaxed">在微服务架构中，一个业务操作往往涉及多个服务的数据更新。Seata 提供 AT、TCC、Saga 和 XA 四种模式，其中 AT 模式对业务代码侵入最小：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@Service
public class BusinessService {

    @Autowired
    private StorageClient storageClient;

    @Autowired
    private OrderClient orderClient;

    @Autowired
    private AccountClient accountClient;

    @GlobalTransactional(name = "purchase-tx", rollbackFor = Exception.class)
    public void purchase(String userId, String commodityCode, int orderCount) {
        // 扣减库存
        storageClient.deduct(commodityCode, orderCount);
        // 创建订单
        orderClient.create(userId, commodityCode, orderCount);
        // 扣减账户余额
        accountClient.debit(userId, totalPrice);
        // 任意步骤失败，全局回滚
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">配置中心动态刷新</h3>
<p class="text-slate-400 leading-relaxed">将配置外置到 Nacos，支持热更新而无需重启服务：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@RefreshScope
@RestController
public class ConfigController {

    @Value("${app.timeout:5000}")
    private int timeout;

    @Value("${app.feature-flag:false}")
    private boolean featureEnabled;

    @GetMapping("/config")
    public Map&lt;String, Object&gt; getConfig() {
        return Map.of("timeout", timeout, "featureEnabled", featureEnabled);
    }
}</code></pre>
</div>`),
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
		Content: template.HTML(`<div class="space-y-8">
<h3 class="text-xl font-bold text-white">性能优化的核心思路</h3>
<p class="text-slate-400 leading-relaxed">数据库性能优化的本质是用最小的系统资源消耗，在最短时间内返回正确的结果。遵循"由浅入深、由点及面"的原则，先定位瓶颈，再对症下药。</p>

<h3 class="text-xl font-bold text-white mt-8">索引设计原则</h3>
<p class="text-slate-400 leading-relaxed">索引是数据库优化中最常用、最有效的手段。合理的索引可以将查询时间从秒级降至毫秒级，但不合理的索引会降低写入性能并浪费存储空间。</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">-- 最左前缀原则示例
CREATE TABLE user_orders (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    status TINYINT NOT NULL DEFAULT 0,
    create_time DATETIME NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    -- 联合索引：user_id + status + create_time
    INDEX idx_user_status_time (user_id, status, create_time)
);

-- 能命中索引的查询
SELECT * FROM user_orders WHERE user_id = 100 AND status = 1;
SELECT * FROM user_orders WHERE user_id = 100;
SELECT * FROM user_orders WHERE user_id = 100 AND status = 1 AND create_time > '2024-01-01';

-- 不能命中索引的查询（缺少最左列）
SELECT * FROM user_orders WHERE status = 1;
SELECT * FROM user_orders WHERE create_time > '2024-01-01';</code></pre>

<h3 class="text-xl font-bold text-white mt-8">执行计划分析：EXPLAIN</h3>
<p class="text-slate-400 leading-relaxed">EXPLAIN 是 MySQL 提供的查询分析工具，通过分析执行计划可以判断查询是否走了索引、扫描了多少行、是否使用了临时表等关键信息。</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">EXPLAIN SELECT * FROM user_orders 
WHERE user_id = 100 
ORDER BY create_time DESC 
LIMIT 20;

-- 关注以下字段：
-- type: range / ref / index / ALL（ALL 代表全表扫描，需优化）
-- key: 实际使用的索引名
-- rows: 预估扫描行数
-- Extra: Using index（覆盖索引）/ Using filesort（需要排序优化）</code></pre>

<h3 class="text-xl font-bold text-white mt-8">慢查询定位与优化</h3>
<p class="text-slate-400 leading-relaxed">开启慢查询日志，持续监控超过阈值的 SQL：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">-- my.cnf 配置
[mysqld]
slow_query_log = 1
slow_query_log_file = /var/log/mysql/slow.log
long_query_time = 1    -- 超过 1 秒的 SQL 会被记录
log_queries_not_using_indexes = 1

-- 常用分析命令
mysqldumpslow -s t -t 10 /var/log/mysql/slow.log   -- 按时间排序 Top 10
pt-query-digest /var/log/mysql/slow.log             -- Percona 工具详细分析</code></pre>

<h3 class="text-xl font-bold text-white mt-8">连接池调优</h3>
<p class="text-slate-400 leading-relaxed">HikariCP 是目前性能最好的连接池。合理配置连接池参数对高并发系统至关重要：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">spring:
  datasource:
    hikari:
      maximum-pool-size: 20          -- 最大连接数（公式：CPU核数 * 2 + 有效磁盘数）
      minimum-idle: 10               -- 最小空闲连接
      connection-timeout: 30000      -- 获取连接超时时间（毫秒）
      idle-timeout: 600000         -- 空闲连接存活时间
      max-lifetime: 1800000        -- 连接最大生命周期（小于 MySQL wait_timeout）
      pool-name: HikariPool-Primary
      leak-detection-threshold: 60000  -- 连接泄漏检测阈值</code></pre>

<h3 class="text-xl font-bold text-white mt-8">分库分表策略</h3>
<p class="text-slate-400 leading-relaxed">当单表数据量超过 500 万行或单库写入 QPS 超过 2000 时，需要考虑分库分表。ShardingSphere 提供了透明化的分片能力：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">-- 水平分表示例：按 user_id 取模分 4 表
CREATE TABLE user_order_0 LIKE user_order_template;
CREATE TABLE user_order_1 LIKE user_order_template;
CREATE TABLE user_order_2 LIKE user_order_template;
CREATE TABLE user_order_3 LIKE user_order_template;

-- ShardingSphere 配置
spring:
  shardingsphere:
    datasource:
      names: ds0, ds1
    rules:
      sharding:
        tables:
          user_order:
            actual-data-nodes: ds$->{0..1}.user_order_$->{0..3}
            table-strategy:
              standard:
                sharding-column: user_id
                sharding-algorithm-name: mod-4
        sharding-algorithms:
          mod-4:
            type: INLINE
            props:
              algorithm-expression: user_order_$->{user_id % 4}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">读写分离</h3>
<p class="text-slate-400 leading-relaxed">读多写少的场景下，读写分离可以显著提升查询吞吐量。ShardingSphere 支持自动路由：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">-- 强制走主库（避免延迟）
HintManager hintManager = HintManager.getInstance();
hintManager.setWriteRouteOnly();
try {
    return orderMapper.selectById(orderId);
} finally {
    hintManager.close();
}</code></pre>
</div>`),
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
		Content: template.HTML(`<div class="space-y-8">
<h3 class="text-xl font-bold text-white">为什么使用缓存</h3>
<p class="text-slate-400 leading-relaxed">在高并发系统中，数据库往往是性能瓶颈。Redis 作为内存数据库，读写性能是 MySQL 的 10-50 倍。通过合理的缓存策略，可以显著降低数据库压力，提升系统响应速度。</p>

<h3 class="text-xl font-bold text-white mt-8">三种经典缓存问题</h3>
<p class="text-slate-400 leading-relaxed">生产环境中，缓存使用不当会引发三种典型问题：</p>

<p class="text-slate-400 leading-relaxed mt-4"><strong class="text-white">1. 缓存穿透</strong> — 查询一个不存在的数据，缓存和数据库都无法命中，导致每次请求都打到数据库。</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// 解决方案：缓存空值 + 布隆过滤器
@Component
public class CachePenetrationSolution {

    @Autowired
    private StringRedisTemplate redisTemplate;

    @Autowired
    private RBloomFilter&lt;String&gt; bloomFilter;

    public User getUser(Long id) {
        String key = "user:" + id;
        String cached = redisTemplate.opsForValue().get(key);

        if (cached != null) {
            // 缓存命中（包括空值标记 "null"）
            return "null".equals(cached) ? null : JSON.parseObject(cached, User.class);
        }

        // 布隆过滤器拦截
        if (!bloomFilter.contains(key)) {
            // 数据一定不存在，直接返回
            return null;
        }

        // 查询数据库
        User user = userMapper.selectById(id);
        if (user == null) {
            // 缓存空值，过期时间设置较短
            redisTemplate.opsForValue().set(key, "null", 60, TimeUnit.SECONDS);
        } else {
            redisTemplate.opsForValue().set(key, JSON.toJSONString(user), 30, TimeUnit.MINUTES);
        }
        return user;
    }
}</code></pre>

<p class="text-slate-400 leading-relaxed mt-4"><strong class="text-white">2. 缓存击穿</strong> — 热点数据在缓存失效的瞬间，大量并发请求同时打到数据库。</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// 解决方案：互斥锁 + 逻辑过期
@Component
public class CacheBreakdownSolution {

    @Autowired
    private RedissonClient redissonClient;

    public User getHotUser(Long id) {
        String key = "user:" + id;
        String json = redisTemplate.opsForValue().get(key);

        if (json != null) {
            RedisData redisData = JSON.parseObject(json, RedisData.class);
            LocalDateTime expireTime = redisData.getExpireTime();

            if (expireTime.isAfter(LocalDateTime.now())) {
                // 逻辑未过期，直接返回
                return JSON.parseObject(redisData.getData(), User.class);
            }

            // 逻辑已过期，尝试获取锁重建缓存
            RLock lock = redissonClient.getLock("lock:user:" + id);
            if (lock.tryLock()) {
                try {
                    // 双重检查
                    json = redisTemplate.opsForValue().get(key);
                    redisData = JSON.parseObject(json, RedisData.class);
                    if (redisData.getExpireTime().isAfter(LocalDateTime.now())) {
                        return JSON.parseObject(redisData.getData(), User.class);
                    }
                    // 重建缓存
                    User user = userMapper.selectById(id);
                    cacheUserWithLogicalExpire(id, user);
                    return user;
                } finally {
                    lock.unlock();
                }
            }
        }

        // 未获取到锁，返回旧数据（保证可用性）
        return JSON.parseObject(JSON.parseObject(json, RedisData.class).getData(), User.class);
    }
}</code></pre>

<p class="text-slate-400 leading-relaxed mt-4"><strong class="text-white">3. 缓存雪崩</strong> — 大量缓存同时失效或 Redis 宕机，导致数据库瞬时压力过大。</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// 解决方案：过期时间加随机值 + 多级缓存
@Component
public class CacheAvalancheSolution {

    @Autowired
    private CacheManager caffeineCacheManager;

    public User getUserWithMultiLevel(Long id) {
        // L1: Caffeine 本地缓存（进程内，毫秒级）
        Cache localCache = caffeineCacheManager.getCache("userLocal");
        User user = localCache.get(id, User.class);
        if (user != null) return user;

        // L2: Redis 分布式缓存（微秒级）
        String key = "user:" + id;
        String json = redisTemplate.opsForValue().get(key);
        if (json != null) {
            user = JSON.parseObject(json, User.class);
            localCache.put(id, user);  // 回填本地缓存
            return user;
        }

        // L3: 数据库
        user = userMapper.selectById(id);
        if (user != null) {
            // 过期时间加随机值，避免同时失效
            long expire = 30 * 60 + ThreadLocalRandom.current().nextInt(300);
            redisTemplate.opsForValue().set(key, JSON.toJSONString(user), expire, TimeUnit.SECONDS);
            localCache.put(id, user);
        }
        return user;
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">分布式锁</h3>
<p class="text-slate-400 leading-relaxed">Redisson 提供了可重入、可续期的分布式锁，是 Redis 实现分布式锁的最佳实践：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@Component
public class DistributedLockDemo {

    @Autowired
    private RedissonClient redissonClient;

    public void deductStock(String productId, int count) {
        RLock lock = redissonClient.getLock("stock:" + productId);

        try {
            // 尝试获取锁，最多等待 10 秒，锁持有时长 30 秒（看门狗自动续期）
            boolean acquired = lock.tryLock(10, 30, TimeUnit.SECONDS);
            if (!acquired) {
                throw new BusinessException("获取库存锁失败，请稍后重试");
            }

            int stock = getStock(productId);
            if (stock < count) {
                throw new BusinessException("库存不足");
            }
            updateStock(productId, stock - count);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new BusinessException("操作被中断");
        } finally {
            if (lock.isHeldByCurrentThread()) {
                lock.unlock();
            }
        }
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">缓存一致性策略</h3>
<p class="text-slate-400 leading-relaxed">缓存与数据库的数据一致性是分布式系统的经典难题。推荐采用 Cache-Aside 模式：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// 读：先查缓存，未命中再查数据库并回填
public User getById(Long id) {
    String key = "user:" + id;
    User user = cache.get(key);
    if (user == null) {
        user = userMapper.selectById(id);
        if (user != null) cache.set(key, user, 30, MINUTES);
    }
    return user;
}

// 写：先更新数据库，再删除缓存（非更新缓存）
@Transactional
public void update(User user) {
    userMapper.updateById(user);
    cache.delete("user:" + user.getId());
}

// 删除：先删数据库，再删缓存
@Transactional
public void delete(Long id) {
    userMapper.deleteById(id);
    cache.delete("user:" + id);
}</code></pre>
</div>`),
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
		Content: template.HTML(`<div class="space-y-8">
<h3 class="text-xl font-bold text-white">JVM 内存结构</h3>
<p class="text-slate-400 leading-relaxed">Java 虚拟机在执行 Java 程序时会把它管理的内存划分为若干个不同的数据区域。理解这些区域的用途和特性是 JVM 调优的基础。</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">┌─────────────────────────────────────────┐
│              堆（Heap）                    │
│  ┌─────────────┐    ┌─────────────────┐ │
│  │  年轻代      │    │     老年代       │ │
│  │  ┌─┐┌─┐┌─┐  │    │                 │ │
│  │  │E│S0│S1│  │    │                 │ │
│  │  └─┘└─┘└─┘  │    │                 │ │
│  └─────────────┘    └─────────────────┘ │
│              -Xms -Xmx                   │
├─────────────────────────────────────────┤
│  元空间（Metaspace）-XX:MaxMetaspaceSize  │
├─────────────────────────────────────────┤
│  虚拟机栈（VM Stack）-Xss                 │
│  本地方法栈（Native Stack）              │
│  程序计数器（PC Register）               │
├─────────────────────────────────────────┤
│  直接内存（Direct Memory）               │
│  -XX:MaxDirectMemorySize                 │
└─────────────────────────────────────────┘</code></pre>

<h3 class="text-xl font-bold text-white mt-8">垃圾回收器演进</h3>
<p class="text-slate-400 leading-relaxed">从 Serial 到 ZGC，JVM 垃圾回收器经历了多次重大演进，核心矛盾始终是"吞吐量 vs 延迟"之间的权衡：</p>
<ul class="list-disc list-inside text-slate-400 space-y-2 mt-2">
<li><strong>Serial GC</strong>：单线程，适合客户端应用</li>
<li><strong>Parallel GC</strong>：多线程并行，追求高吞吐量</li>
<li><strong>CMS</strong>：并发标记清除，降低停顿时间，但碎片化严重</li>
<li><strong>G1</strong>：区域化收集，可预测停顿时间（Java 9 默认）</li>
<li><strong>ZGC</strong>：低延迟（&lt;10ms），可扩展至 TB 级堆（Java 15 生产可用）</li>
<li><strong>Shenandoah</strong>：Red Hat 开发的低延迟 GC</li>
</ul>

<h3 class="text-xl font-bold text-white mt-8">G1 垃圾回收器详解</h3>
<p class="text-slate-400 leading-relaxed">G1（Garbage First）将堆划分为多个大小相等的 Region，每个 Region 根据需要动态扮演 Eden、Survivor 或 Old 角色。</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300"># G1 推荐参数（Java 17，8GB 堆）
-XX:+UseG1GC
-XX:MaxGCPauseMillis=200          # 目标最大停顿时间
-Xms8g -Xmx8g                      # 固定堆大小避免扩容
-XX:+AlwaysPreTouch               # 启动时预分配内存
-XX:InitiatingHeapOccupancyPercent=45  # 老年代占比触发并发标记
-XX:G1HeapRegionSize=16m          # Region 大小（根据堆大小自动计算）
-XX:+UseStringDeduplication        # 字符串去重（Java 8u20+）</code></pre>

<h3 class="text-xl font-bold text-white mt-8">ZGC：亚毫秒级停顿</h3>
<p class="text-slate-400 leading-relaxed">ZGC 使用染色指针和读屏障技术，实现了全并发垃圾回收，停顿时间几乎不受堆大小影响：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300"># ZGC 启动参数（Java 17+）
-XX:+UseZGC
-Xms32g -Xmx32g
-XX:+ZGenerational                 # Java 21 分代 ZGC（性能更优）

# ZGC 日志分析
-Xlog:gc*:file=/var/log/gc.log:time,uptime,level,tags:filecount=5,filesize=100m

# 关键指标
[gc,start] GC(0) Garbage Collection (Proactive)
[gc,phases] GC(0) Pause Mark Start 0.012ms      # 停顿时间
[gc,phases] GC(0) Pause Mark End 0.015ms
[gc,phases] GC(0) Pause Relocate Start 0.018ms
[gc,phases] GC(0) Pause Relocate End 0.021ms
[gc,heap] GC(0) Heap: 8192M(8192M) -> 4096M(8192M)  # 回收前后堆占用</code></pre>

<h3 class="text-xl font-bold text-white mt-8">常见 OOM 场景诊断</h3>
<p class="text-slate-400 leading-relaxed">OOM 是生产环境最常见的问题之一，不同区域的 OOM 有不同的排查思路：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// 1. Java heap space
// 原因：内存泄漏或堆设置过小
// 排查：
jmap -dump:format=b,file=heap.hprof &lt;pid&gt;
# 使用 MAT / VisualVM 分析 Dominator Tree

// 2. Metaspace
// 原因：动态生成类过多（如 CGLIB、反射）
// 排查：
-XX:+TraceClassLoading -XX:+TraceClassUnloading
# 观察类加载日志

// 3. Direct buffer memory
// 原因：NIO 直接内存泄漏（Netty 堆外内存未释放）
// 排查：
# 开启 Native Memory Tracking
-XX:NativeMemoryTracking=summary
jcmd &lt;pid&gt; VM.native_memory summary

// 4. Unable to create new native thread
// 原因：线程数超过系统限制（ulimit -u）
// 解决：
ulimit -u 65535  # 修改系统限制
# 或使用线程池替代无限创建线程</code></pre>

<h3 class="text-xl font-bold text-white mt-8">Arthas 线上诊断</h3>
<p class="text-slate-400 leading-relaxed">Arthas 是阿里巴巴开源的 Java 诊断工具，可以实时查看 JVM 状态、方法耗时、线程堆栈等：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300"># 连接目标进程
java -jar arthas-boot.jar

# 热更新代码（无需重启）
redefine /path/to/UserService.class

# 方法执行监控（查看入参、返回值、耗时）
watch com.example.service.UserService getUserById \
  "{params, returnObj, throwExp}" \
  -x 2 \
  '#cost>100'    # 只监控耗时超过 100ms 的调用

# 火焰图生成
profiler start --event cpu
profiler stop --file /tmp/flamegraph.html

# 线程死锁检测
thread --block

# 反编译查看线上代码
jad com.example.service.UserService</code></pre>

<h3 class="text-xl font-bold text-white mt-8">JVM 调优实战建议</h3>
<p class="text-slate-400 leading-relaxed">调优不是一蹴而就的，遵循"测量-分析-调整"的闭环：</p>
<ol class="list-decimal list-inside text-slate-400 space-y-2 mt-2">
<li>先加 GC 日志和监控，收集运行数据</li>
<li>分析 GC 频率和停顿时间是否符合预期</li>
<li>通过 MAT 分析堆 dump，确认是否存在内存泄漏</li>
<li>调整参数后对比前后指标变化</li>
<li>避免盲目调优，每次只改一个参数</li>
</ol>
</div>`),
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
		Content: template.HTML(`<div class="space-y-8">
<h3 class="text-xl font-bold text-white">并发编程的核心挑战</h3>
<p class="text-slate-400 leading-relaxed">并发编程的难点在于正确性、可见性和有序性。Java 内存模型（JMM）定义了线程和主内存之间的抽象关系，synchronized 和 volatile 是保障并发安全的两个基础关键字。</p>

<h3 class="text-xl font-bold text-white mt-8">volatile：轻量级同步</h3>
<p class="text-slate-400 leading-relaxed">volatile 保证变量的可见性和有序性，但不保证原子性。适合作为状态标记或单次写多次读的场景：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">public class VolatileExample {
    private volatile boolean running = true;

    public void shutdown() {
        running = false;  // 所有线程立即可见
    }

    public void doWork() {
        while (running) {  // 不会读取到本地缓存的旧值
            // 执行任务
        }
    }
}

// volatile 的双重检查锁定（DCL）
public class Singleton {
    private volatile static Singleton instance;

    public static Singleton getInstance() {
        if (instance == null) {              // 第一次检查（无锁）
            synchronized (Singleton.class) {
                if (instance == null) {        // 第二次检查（有锁）
                    instance = new Singleton(); // volatile 禁止指令重排序
                }
            }
        }
        return instance;
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">synchronized 与锁升级</h3>
<p class="text-slate-400 leading-relaxed">Java 6 之后，synchronized 不再是重量级锁的代名词。JVM 引入了锁升级机制：无锁 → 偏向锁 → 轻量级锁 → 重量级锁：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// 锁对象头的 Mark Word 结构（64 位 JVM）
// 无锁：    [hashcode:25] | [age:4] | [biased_lock:0] | [lock:01]
// 偏向锁：  [thread:54]   | [epoch:2] | [age:4] | [biased_lock:1] | [lock:01]
// 轻量锁：  [ptr_to_lock_record:62] | [lock:00]
// 重量锁：  [ptr_to_monitor:62]     | [lock:10]

// 生产环境中，可通过以下参数观察锁状态
-XX:+PrintBiasedLockingStatistics  # 打印偏向锁统计
-XX:BiasedLockingStartupDelay=0   # 启动即开启偏向锁（默认 4 秒延迟）</code></pre>

<h3 class="text-xl font-bold text-white mt-8">ReentrantLock：更灵活的锁</h3>
<p class="text-slate-400 leading-relaxed">相比 synchronized，ReentrantLock 提供了公平锁、可中断、超时获取等高级特性：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">public class ReentrantLockDemo {
    private final ReentrantLock lock = new ReentrantLock(true);  // 公平锁
    private final Condition condition = lock.newCondition();
    private int count = 0;

    public void increment() {
        lock.lock();
        try {
            count++;
            if (count >= 10) {
                condition.signalAll();  // 唤醒等待的线程
            }
        } finally {
            lock.unlock();  // 必须在 finally 中释放
        }
    }

    public boolean tryIncrement(long timeout, TimeUnit unit) {
        try {
            if (lock.tryLock(timeout, unit)) {  // 带超时的获取
                try {
                    count++;
                    return true;
                } finally {
                    lock.unlock();
                }
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        return false;
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">线程池：并发的基础设施</h3>
<p class="text-slate-400 leading-relaxed">线程池是管理线程生命周期的最佳实践。核心参数的选择直接影响系统稳定性：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// 线程池参数设计公式
// corePoolSize = CPU 核数 + 1（计算密集型）
// corePoolSize = CPU 核数 * 2（IO 密集型）
// maximumPoolSize = corePoolSize + (核心数 / (1 - 阻塞系数))
// workQueue = 有界队列（ArrayBlockingQueue），避免无限堆积
// keepAliveTime = 60s（非核心线程存活时间）
// rejectionPolicy = CallerRunsPolicy（让提交者自己执行，起到背压效果）

@Component
public class ThreadPoolConfig {

    @Bean("ioExecutor")
    public ThreadPoolExecutor ioExecutor() {
        int core = Runtime.getRuntime().availableProcessors();
        return new ThreadPoolExecutor(
            core * 2,                       // 核心线程数
            core * 4,                       // 最大线程数
            60L, TimeUnit.SECONDS,
            new ArrayBlockingQueue<>(1000), // 有界队列
            new ThreadFactoryBuilder()
                .setNameFormat("io-pool-%d")
                .build(),
            new ThreadPoolExecutor.CallerRunsPolicy()
        );
    }

    @Bean("computeExecutor")
    public ThreadPoolExecutor computeExecutor() {
        int core = Runtime.getRuntime().availableProcessors();
        return new ThreadPoolExecutor(
            core + 1,
            core + 1,
            0L, TimeUnit.MILLISECONDS,       // 计算密集型线程不回收
            new LinkedBlockingQueue<>(100),
            new ThreadFactoryBuilder()
                .setNameFormat("compute-pool-%d")
                .build(),
            new ThreadPoolExecutor.AbortPolicy()
        );
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">CompletableFuture：异步编排</h3>
<p class="text-slate-400 leading-relaxed">CompletableFuture 是 Java 8 引入的异步编程工具，支持链式调用、组合多个异步任务、异常处理等：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">@Service
public class OrderService {

    @Autowired
    @Qualifier("ioExecutor")
    private ExecutorService executor;

    public CompletableFuture&lt;OrderDetail&gt; getOrderDetail(Long orderId) {
        // 并行查询：订单、用户、商品、物流
        CompletableFuture&lt;Order&gt; orderFuture = CompletableFuture
            .supplyAsync(() -> orderMapper.selectById(orderId), executor);

        CompletableFuture&lt;User&gt; userFuture = orderFuture
            .thenCompose(order -> CompletableFuture
                .supplyAsync(() -> userService.getById(order.getUserId()), executor));

        CompletableFuture&lt;List&lt;Product&gt;&gt; productFuture = orderFuture
            .thenCompose(order -> CompletableFuture
                .supplyAsync(() -> productService.getByIds(order.getProductIds()), executor));

        CompletableFuture&lt;Logistics&gt; logisticsFuture = orderFuture
            .thenCompose(order -> CompletableFuture
                .supplyAsync(() -> logisticsService.getByOrderId(orderId), executor)
                .exceptionally(ex -> {
                    log.warn("Logistics query failed: {}", ex.getMessage());
                    return Logistics.EMPTY;  // 降级处理
                }));

        // 等待所有任务完成并组装结果
        return CompletableFuture.allOf(userFuture, productFuture, logisticsFuture)
            .thenApply(v -> {
                Order order = orderFuture.join();
                return OrderDetail.builder()
                    .order(order)
                    .user(userFuture.join())
                    .products(productFuture.join())
                    .logistics(logisticsFuture.join())
                    .build();
            });
    }
}</code></pre>

<h3 class="text-xl font-bold text-white mt-8">并发集合：线程安全的数据结构</h3>
<p class="text-slate-400 leading-relaxed">JUC 包提供了多种线程安全的集合实现，选择合适的并发集合可以大幅提升性能：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// ConcurrentHashMap：分段锁 → CAS + synchronized（Java 8+）
ConcurrentHashMap&lt;String, User&gt; map = new ConcurrentHashMap&lt;&gt;();
// computeIfAbsent 是原子操作，无需额外同步
User user = map.computeIfAbsent(userId, id -> loadFromDatabase(id));

// CopyOnWriteArrayList：读多写少场景（如配置列表）
CopyOnWriteArrayList&lt;String&gt; configList = new CopyOnWriteArrayList&lt;&gt;();
// 读操作无锁，写操作复制整个数组（适合读多写少）

// LinkedBlockingQueue：生产者-消费者模式
LinkedBlockingQueue&lt;Task&gt; queue = new LinkedBlockingQueue&lt;&gt;(1000);
// 生产者
queue.put(task);  // 队列满时阻塞
// 消费者
Task task = queue.take();  // 队列空时阻塞

// LongAdder：高并发计数（比 AtomicLong 性能更优）
LongAdder counter = new LongAdder();
counter.increment();  // 分散到多个 Cell 减少竞争
long sum = counter.sum();  // 读取时汇总</code></pre>

<h3 class="text-xl font-bold text-white mt-8">Fork/Join 框架：分治并行</h3>
<p class="text-slate-400 leading-relaxed">Fork/Join 框架适合可分解的递归任务，利用工作窃取算法平衡线程负载：</p>
<pre class="bg-dark-bg border border-dark-border rounded-lg p-4 mt-2 overflow-x-auto"><code class="text-sm text-slate-300">// 并行计算大数据数组的和
public class SumTask extends RecursiveTask&lt;Long&gt; {
    private static final int THRESHOLD = 10000;
    private final int[] array;
    private final int start;
    private final int end;

    @Override
    protected Long compute() {
        if (end - start <= THRESHOLD) {
            // 直接计算
            long sum = 0;
            for (int i = start; i < end; i++) {
                sum += array[i];
            }
            return sum;
        }

        // 拆分为两个子任务
        int mid = (start + end) / 2;
        SumTask left = new SumTask(array, start, mid);
        SumTask right = new SumTask(array, mid, end);

        left.fork();   // 异步执行左任务
        long rightResult = right.compute();  // 同步执行右任务
        long leftResult = left.join();  // 等待左任务结果

        return leftResult + rightResult;
    }
}

// 使用方式
ForkJoinPool pool = new ForkJoinPool();
long result = pool.invoke(new SumTask(array, 0, array.length));</code></pre>
</div>`),
		TechStack:     []string{"Java 17", "JUC", "CompletableFuture", "Fork/Join"},
		Difficulty:    "高级",
		ReadTime:      "40 分钟",
		CreatedAt:     "2026-05-21",
		UpdatedAt:     "2026-05-21",
	},
}
