# http API (net/http + your pkgs) — Exercises

> Work in THIS folder, `package httpapi`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 server.go** — New(addr, logger) with routes; Todo{ID,Title,Done}; Start(ctx) error.
- **E2 curl transcript** — E2 curl transcript: document 3 curl commands + pasted responses in QUESTIONS.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package httpapi`.

## Build-chain (reuse — at least 2 of your packages)
- MUST wire ≥2 of YOUR packages: `18_logger` (request logs), `20_cache` or mutex store, `21_config` (port), `19_event` (request events), `23_jsonlite` (compare against `encoding/json` responses).
- Impl uses `net/http` (allowed — it's the server, not the thing you're rebuilding). Handlers return YOUR errors mapped to status codes.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
