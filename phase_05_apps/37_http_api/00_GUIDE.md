# http API (net/http + your pkgs) — Guide (`httpapi`)

> Folder: `phase_05_apps/37_http_api` · Package: `httpapi` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- APP QUEST: build JSON CRUD with ONLY net/http + your logger/cache/config/event.
- Step 1 — GET /health, GET/POST /todos (memory store behind your cache or mutex).
- Step 2 — logging middleware using YOUR logger; request-id via event bus stretch.
- Step 3 — graceful shutdown with context.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package httpapi` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_05_apps/37_http_api"` (package name `httpapi`), call your funcs, `go run .`.
3. `go vet ./phase_05_apps/37_http_api/...` must be clean before moving on.

### Details
- APP QUEST: build JSON CRUD with ONLY net/http + your logger/cache/config/event.
- Step 1 — GET /health, GET/POST /todos (memory store behind your cache or mutex).
- Step 2 — logging middleware using YOUR logger; request-id via event bus stretch.
- Step 3 — graceful shutdown with context.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_05_apps/37_http_api")
func main() { fmt.Println(httpapi.Demo()) }
```
```powershell
go run .
go vet ./phase_05_apps/37_http_api/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
