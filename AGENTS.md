# devyunmu.com — Agent Guide

**Generated:** 2026-05-13
**Branch:** master

## Quick start

```sh
go build -o devyunmu.com ./cmd/server
./devyunmu.com              # serves on :8080 (override with PORT env)
```

No tests, linter, or formatter config exist. No Makefile, no CI.

## Structure

```
devyunmu.com/
├── cmd/server/main.go       # Entrypoint — wires Gin engine, middleware, routes
├── config/config.go         # Env-based config (PORT only)
├── internal/
│   ├── handler/home.go      # HTTP handlers: Home, About, Health
│   ├── middleware/middleware.go  # Gzip, Cache, Security headers
│   └── model/page.go        # NavItem, SiteInfo types
├── templates/               # Go html/template; base.html layout + partials
├── static/
│   ├── css/style.css        # All styling (267 lines)
│   ├── css/logo-font.css    # @font-face rule (unused)
│   ├── images/              # SVG icons + full_logo
│   ├── js/                  # Empty (no JS)
│   ├── robots.txt           # Allow all
│   └── sitemap.xml          # / and /about
├── vendor/                  # Vendored deps (go mod vendor)
└── Dockerfile               # Multi-stage: go build → scratch
```

## Where to look

| Task | Location |
|------|----------|
| Add route/handler | `cmd/server/main.go` + `internal/handler/home.go` |
| Change middleware | `internal/middleware/middleware.go` |
| Update site info | `internal/model/page.go` |
| Edit layout | `templates/base.html` |
| Adjust style | `static/css/style.css` |
| Deploy | `Dockerfile` |

## Conventions

- **Middleware stacking**: `gin.Default()` → Gzip → Cache → Security. Order matters (Gzip wraps ResponseWriter before Cache sets headers).
- **Templates**: `{{ template "base.html" . }}` with `{{ define "content" }}` blocks. Partials prefixed `_` (`_head.html`, `_nav.html`, `_footer.html`). Custom `now` FuncMap exposes `{{ now.Year }}`.
- **CSS**: `style.css` has all styling (267 lines, BEM-ish naming). No framework. `logo-font.css` is unused (see anti-patterns).
- **Handlers**: Return `c.HTML()` or `c.Status()` directly. No service layer — logic lives in handler.
- **Config**: Env-only, no config files. `getEnv(key, fallback)` helper.
- **Go proxy blocked** (China) — use `go mod vendor` + `go build -mod=vendor`. No network during Docker build.
- **Binary** (`devyunmu.com`): `.gitignore`d but present in working dir (pre-built).
- **Site name**: Dev云沐 (used in logo, page title).
- **Owner**: 云沐 (Yun Mu), m@devyunmu.com.
- **Doc comments**: Intentionally absent from all exported Go symbols. Project docs live in this file.
- **Scratch Docker**: Uses `FROM scratch` for minimal image size (no libc).
- **Zero HTML/CSS comments**: Templates and stylesheets contain no inline comments whatsoever.
- **Dual favicon**: Both `/favicon.ico` and `/favicon.svg` serve `static/images/icon.svg`.
- **Binary name**: Built as `devyunmu.com` (not `server` despite `cmd/server/` dir).

## Anti-patterns (this project)

- **No JS** — all interactivity is CSS-only (hover, transitions). Don't add JS dependencies.
- **No database** — static site. Don't add DB drivers or ORMs.
- **No external services** — no API calls, no CDN, no analytics.
- **No linter/formatter override** — relies on `gofmt` defaults.
- **No test files** — no `_test.go` anywhere.
- **Orphaned `logo-font.css`**: `static/css/logo-font.css` is never loaded by any template and duplicates `.logo-text` from `style.css`. Dead code.
- **No `.env.example`**: `.env` is gitignored but no sample env file exists for reference.
- **Vendor dir not committed**: `vendor/` is in `.gitignore` but required for build. New contributors must run `go mod vendor` after checkout.
- **No error wrapping**: Handlers return `c.HTML()`/`c.Status()` directly with no structured error handling or logging.

## Commands

```bash
go build -o devyunmu.com ./cmd/server           # Build
./devyunmu.com                                   # Run (port :8080)
go build -mod=vendor -o devyunmu.com ./cmd/server  # Build with vendored deps
docker build -t devyunmu.com .                   # Docker build
```

## Deployment

```bash
# Scp + deploy on remote (118.178.194.95)
docker save devyunmu.com | gzip > image.tar.gz
scp image.tar.gz root@118.178.194.95:/tmp/
ssh root@118.178.194.95 "docker load < /tmp/image.tar.gz && docker run -d -p 8080:8080 --restart unless-stopped --name devyunmu devyunmu.com"
```

## Code map

| Symbol | Kind | File | Role |
|--------|------|------|------|
| `main` | func | `cmd/server/main.go` | Entrypoint — engine setup, route registration |
| `Config.Load` | func | `config/config.go` | Read PORT env, default 8080 |
| `Home` | func | `internal/handler/home.go` | GET / → index.html |
| `About` | func | `internal/handler/home.go` | GET /about → about.html |
| `Health` | func | `internal/handler/home.go` | GET /health → 204 |
| `Gzip` | func | `internal/middleware/middleware.go` | Gzip compression middleware |
| `Cache` | func | `internal/middleware/middleware.go` | Cache-Control (1h static, 1y /static) |
| `Security` | func | `internal/middleware/middleware.go` | nosniff, DENY, HSTS, CSP headers |
| `NavItem` | type | `internal/model/page.go` | Nav link {Label, Href} |
| `SiteInfo` | type | `internal/model/page.go` | Site metadata struct |
| `NavItems` | var | `internal/model/page.go` | [{Home, /}, {About, /about}] |
| `Info` | var | `internal/model/page.go` | Site owner, desc, email, URLs |
