# Project Fix & Docker Deploy — devyunmu.com

## TL;DR

> **Quick Summary**: Fix all identified issues (Dockerfile duplication, template glob crash, broken tests, dead code, missing configs) and deploy a working Docker container to localhost:8080.
>
> **Deliverables**:
> - All `go test ./...` passing (zero failures)
> - Clean Dockerfile (no duplicate COPY)
> - Running Docker container serving on `:8080`
> - `.dockerignore`, `.env.example` created
> - Dead code removed (logo-font.css, error_test.go, AppError type)
>
> **Estimated Effort**: Short
> **Parallel Execution**: YES — 2 waves
> **Critical Path**: Task 5 → Task 9 → Task 10 → Task 11

---

## Context

### Original Request
用户要求全面检查 devyunmu.com 项目问题，修复所有发现的问题，然后部署到本地 Docker。

### Interview Summary
**Key Discussions**:
- 修复范围: 全面修复（所有已发现问题 + 测试修复 + Docker 部署）
- 测试策略: Tests-after + agent QA（修复已有测试，确保通过；所有交付物由 agent 进行 QA 验证）
- Go 编译使用 `-mod=vendor`（中国网络限制）

**Research Findings**:
- Docker 容器崩溃根因: Go 的 `filepath.Glob` 不支持 `**` 递归匹配，`templates/**/*.html` 匹配零文件导致 panic
- `middleware_test.go:34` — `TestGzip` 测试函数缺少 `func` 声明（裸代码编译失败）
- `error_test.go` — 仅 2 行（package + 注释），无实际测试
- `AppError` 类型定义在 `handler/home.go` 但从未被使用
- `logo-font.css` — AGENTS.md 确认为死代码，无模板引用
- Dockerfile L13-15 与 L17-19 完全重复
- 项目无 `.dockerignore`、无 `.env.example`
- Git 仓库零提交历史

### Metis Review
**Identified Gaps** (addressed):
- Go `filepath.Glob` 不支持 `**` → 回退到 `templates/*.html`（所有模板扁平目录）
- `AppError` 死代码 → 纳入修复范围
- `.dockerignore` 内容 → 排除 `.git/`, `.sisyphus/`, `devyunmu.com` 二进制, `.env`（保留 `vendor/`，Dockerfile 的 COPY 需要它）
- Git 提交策略 → 单次初始提交，包含所有修复

---

## Work Objectives

### Core Objective
Fix all identified issues in the devyunmu.com project (Dockerfile duplication, template glob crash, broken tests, dead code, missing config files) and deploy a working Docker container locally.

### Concrete Deliverables
- `internal/middleware/middleware_test.go` — `TestGzip` 函数声明修复
- `internal/middleware/error_test.go` — 删除（空文件）
- `static/css/logo-font.css` — 删除（死代码）
- `internal/handler/home.go` — 移除 `AppError` 类型
- `cmd/server/main.go` — 模板 glob 修复：`templates/**/*.html` → `templates/*.html`
- `Dockerfile` — 移除重复 COPY 行
- `.dockerignore` — 新建
- `.env.example` — 新建
- `devyunmu.com` Docker 镜像 — 本地构建
- Docker 容器 — 运行在 localhost:8080

### Definition of Done
- [ ] `go test -mod=vendor -count=1 ./...` → 全部 PASS
- [ ] `docker build -t devyunmu.com .` → 成功
- [ ] `curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/` → 200
- [ ] `curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/about` → 200
- [ ] `curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/health` → 204

### Must Have
- 所有测试通过（零失败）
- Docker 容器启动不崩溃（无 panic）
- 首页 `/` 返回 200 且包含 HTML 内容
- `/health` 返回 204
- `/about` 返回 200

### Must NOT Have (Guardrails)
- 不修改任何业务逻辑或路由结构
- 不修改样式文件（style.css）
- 不修改模板内容（只修复引用路径）
- 不新增功能或依赖
- 不新增测试（只修复已有测试）
- 不改变中间件顺序
- 不修改 handler 返回逻辑

---

## Verification Strategy

> **ZERO HUMAN INTERVENTION** — ALL verification is agent-executed.

### Test Decision
- **Infrastructure exists**: YES (`go test` works, but tests currently broken)
- **Automated tests**: Tests-after
- **Framework**: `go test` (standard library `testing`)
- **Agent-Executed QA**: Mandatory for ALL tasks — curl for HTTP, bash for CLI

### QA Policy
Every task MUST include agent-executed QA scenarios.
Evidence saved to `.sisyphus/evidence/task-{N}-{scenario-slug}.{ext}`.

- **API/Backend**: Use Bash (curl) — Send requests, assert status + response fields
- **Build/CLI**: Use Bash — `go test`, `docker build`, `docker run`
- **File verification**: Use Bash (grep/ls/cat) — Verify files exist/have expected content

---

## Execution Strategy

### Parallel Execution Waves

```
Wave 1 (Start Immediately — 8 independent fixes, MAX PARALLEL):
├── Task 1: Fix middleware_test.go (TestGzip declaration) [quick]
├── Task 2: Remove error_test.go (empty file) [quick]
├── Task 3: Remove logo-font.css (dead code) [quick]
├── Task 4: Remove AppError type from handler/home.go [quick]
├── Task 5: Fix main.go template glob [quick]
├── Task 6: Fix Dockerfile (remove duplicate COPY) [quick]
├── Task 7: Create .dockerignore [quick]
└── Task 8: Create .env.example [quick]

Wave 2 (After Wave 1 — verification + deploy, SEQUENTIAL):
├── Task 9: Run all tests, verify PASS [unspecified-high]
├── Task 10: Build Docker image [unspecified-high]
└── Task 11: Deploy to local Docker + verify all routes [unspecified-high]

Critical Path: Task 5 → Task 9 → Task 10 → Task 11
Parallel Speedup: ~87% faster than sequential (8 parallel → 3 sequential)
Max Concurrent: 8 (Wave 1)
```

### Dependency Matrix

| Task | Depends On | Blocks | Wave |
|------|-----------|--------|------|
| 1 | — | 9 | 1 |
| 2 | — | 9 | 1 |
| 3 | — | 10 | 1 |
| 4 | — | 9 | 1 |
| 5 | — | 9, 10 | 1 |
| 6 | — | 10 | 1 |
| 7 | — | 10 | 1 |
| 8 | — | — | 1 |
| 9 | 1, 2, 4, 5 | 10 | 2 |
| 10 | 3, 5, 6, 7, 9 | 11 | 2 |
| 11 | 10 | — | 2 |

### Agent Dispatch Summary

- **Wave 1**: 8 × `quick` — T1-T8
- **Wave 2**: 3 × `unspecified-high` — T9-T11

---

## TODOs

- [x] 1. Fix middleware_test.go - TestGzip function declaration

  What to do: Insert func TestGzip(t *testing.T) before the bare code at line 34 in middleware_test.go
  Must NOT do: Do not modify other test functions or add new test cases
  Agent: quick
  Refs: middleware_test.go:17-28 (func pattern), :34-51 (bare code)
  AC: go vet passes; go test -run TestGzip -v passes (exit 0)
  QA: bash: go test -mod=vendor -run TestGzip -v ./internal/middleware/; expect exit 0; evidence task-1.txt
  Commit: see Commit Strategy section

- [x] 2. Remove empty error_test.go

  What to do: Delete internal/middleware/error_test.go (only package decl, no tests)
  Agent: quick
  AC: file gone; go build ./... passes
  QA: bash: ls internal/middleware/error_test.go returns ENOENT; evidence task-2.txt
  Commit: see Commit Strategy section

- [x] 3. Remove dead logo-font.css

  What to do: Delete static/css/logo-font.css (AGENTS.md confirmed orphaned)
  Agent: quick
  AC: file gone; grep -r logo-font templates/ returns nothing
  QA: bash: ls static/css/logo-font.css returns ENOENT; grep -r logo-font templates/ has no output; evidence task-3.txt
  Commit: see Commit Strategy section

- [x] 4. Remove dead AppError type from handler/home.go

  What to do: Remove AppError struct (lines 11-14) and Error() method (lines 16-18) from handler/home.go
  Agent: quick
  Ref: handler/home.go:11-18; Metis confirmed zero references
  AC: AppError gone; go build passes; grep -r AppError internal/ cmd/ returns nothing
  QA: bash: go build -mod=vendor ./... (exit 0); grep -r AppError internal/ cmd/ (no output); evidence task-4.txt
  Commit: see Commit Strategy section

- [x] 5. Fix main.go template glob pattern

  What to do: Change r.LoadHTMLGlob(templates/**/*.html) to r.LoadHTMLGlob(templates/*.html) in cmd/server/main.go line 33
  Root cause: Go filepath.Glob does not support ** recursive matching; pattern matches zero files causing panic
  Agent: quick
  Ref: cmd/server/main.go:33; Go documentation on filepath.Glob
  AC: go build passes; binary starts without panic when templates/ exists
  QA scenario 1 (happy): bash: go build -mod=vendor -o /tmp/devyunmu-test ./cmd/server && PORT=9999 /tmp/devyunmu-test > /tmp/test.log 2>&1 & sleep 1; curl -s -o /dev/null -w %{http_code} http://localhost:9999/; expect 200; kill %1; evidence task-5-startup.txt
  QA scenario 2 (health): bash: same startup; curl -s -o /dev/null -w %{http_code} http://localhost:9999/health; expect 204; evidence task-5-health.txt
  Commit: see Commit Strategy section

- [x] 6. Fix Dockerfile - remove duplicate COPY lines

  What to do: Remove duplicated COPY lines 17-19 from Dockerfile (exact duplicate of lines 13-15)
  Agent: quick
  Ref: Dockerfile:13-19
  AC: Dockerfile has clean COPY block; no duplicate instructions
  QA: bash: grep -c COPY --from=builder /app/devyunmu.com Dockerfile; expect 1 (no duplicate); evidence task-6.txt
  Commit: see Commit Strategy section

- [x] 7. Create .dockerignore

  What to do: Create .dockerignore with entries: .git/ .sisyphus/ devyunmu.com .env
  Agent: quick
  AC: .dockerignore exists with correct entries
  QA: bash: cat .dockerignore; verify each required entry present; evidence task-7.txt
  Commit: see Commit Strategy section

- [x] 8. Create .env.example

  What to do: Create .env.example containing PORT=8080
  Agent: quick
  AC: .env.example exists; contains PORT=8080
  QA: bash: grep PORT .env.example; expect PORT=8080; evidence task-8.txt
  Commit: see Commit Strategy section

- [ ] 9. Run all tests and verify PASS

  What to do: Run go test -mod=vendor -count=1 ./... across all packages; verify ALL packages pass with zero failures
  Depends on: Tasks 1, 2, 4, 5 (all Go code fixes)
  Agent: unspecified-high
  AC: go test ./... exit code 0; zero FAIL lines in output
  QA scenario (all tests): bash: go test -mod=vendor -v -count=1 ./... 2>&1; expect exit 0; grep FAIL output should be empty; evidence task-9-all-tests.txt
  QA scenario (middleware): bash: go test -mod=vendor -v -count=1 ./internal/middleware/ 2>&1; expect all TestXxx PASS; no FAIL; evidence task-9-middleware.txt
  QA scenario (server): bash: go test -mod=vendor -v -count=1 ./cmd/server/ 2>&1; expect TestHealth PASS; evidence task-9-server.txt
  Commit: NO (verification step, no code changes)

- [ ] 10. Build Docker image

  What to do: Build Docker image from the fixed codebase: docker build -t devyunmu.com .
  Preconditions: Tasks 1-9 all pass; Dockerfile fixed; .dockerignore created; all templates present
  Agent: unspecified-high
  AC: docker build succeeds (exit 0); image size is reasonable (approx 18-20 MB)
  QA scenario (build): bash: docker build -t devyunmu.com . 2>&1; expect exit 0; docker images devyunmu.com --format '{{.Size}}' should show reasonable size; evidence task-10-build.txt
  QA scenario (image exists): bash: docker inspect devyunmu.com:latest --format '{{.Created}}'; expect valid timestamp; evidence task-10-inspect.txt
  Commit: NO (build artifact verification)

- [ ] 11. Deploy to local Docker and verify all routes

  What to do:
    1. Stop and remove any existing devyunmu container: docker stop devyunmu 2>/dev/null; docker rm devyunmu 2>/dev/null
    2. Start new container: docker run -d -p 8080:8080 --restart unless-stopped --name devyunmu devyunmu.com
    3. Verify all routes return expected HTTP status codes
  Agent: unspecified-high
  AC: Container running (docker ps shows devyunmu with status Up); all routes verified
  QA scenario 1 (home page): bash: curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/; expect 200; evidence task-11-home.txt
  QA scenario 2 (about page): bash: curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/about; expect 200; evidence task-11-about.txt
  QA scenario 3 (health endpoint): bash: curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/health; expect 204; evidence task-11-health.txt
  QA scenario 4 (static assets): bash: curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/static/css/style.css; expect 200; evidence task-11-static.txt
  QA scenario 5 (favicon): bash: curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/favicon.ico; expect 200; evidence task-11-favicon.txt
  QA scenario 6 (robots.txt): bash: curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/robots.txt; expect 200; evidence task-11-robots.txt
  QA scenario 7 (container not restarting): bash: sleep 2; docker inspect devyunmu --format '{{.State.Restarting}}'; expect false; evidence task-11-stable.txt
  Commit: NO (deployment verification)

---

## Final Verification Wave

- [ ] F1. Plan Compliance Audit - oracle: read plan + verify all Must Have / Must NOT Have
- [ ] F2. Code Quality Review - unspecified-high: go vet + go test + review diffs
- [ ] F3. Real Manual QA - unspecified-high: execute all QA scenarios from tasks 1-11
- [ ] F4. Scope Fidelity Check - deep: verify every task delivered what was specified, no scope creep

---

## Commit Strategy

All Wave 1 fixes (Tasks 1-8) committed together as a single commit:

fix: resolve all project issues and prepare for deployment

- Fix TestGzip function declaration in middleware_test.go
- Remove empty error_test.go
- Remove dead logo-font.css
- Remove unused AppError type from handler/home.go
- Fix template glob pattern (filepath.Glob does not support **)
- Remove duplicate COPY instructions in Dockerfile
- Add .dockerignore
- Add .env.example

Pre-commit verification: go test -mod=vendor -count=1 ./...

---

## Success Criteria

### Verification Commands
```bash
go test -mod=vendor -count=1 ./...          # All tests pass
docker build -t devyunmu.com .              # Image builds successfully
docker run -d -p 8080:8080 --name devyunmu devyunmu.com  # Container starts without panic
curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/         # 200
curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/about    # 200
curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/health   # 204
```

### Final Checklist
- [ ] All Must Have items present
- [ ] All Must NOT Have items absent
- [ ] All tests pass (go test ./... exit 0)
- [ ] Docker container running and stable (not restarting)
- [ ] All routes return correct HTTP status codes
- [ ] Git commit created with all changes
