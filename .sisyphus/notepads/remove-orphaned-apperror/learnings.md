# Learnings

## Task 4 — Remove AppError

- `AppError` was defined at `internal/handler/home.go:11-18` (struct + Error() method) with zero references across the codebase
- Removal confirmed by grep returning empty across `internal/` and `cmd/`
- Build succeeds with `go build -mod=vendor ./...`
- No LSP diagnostics errors after removal
