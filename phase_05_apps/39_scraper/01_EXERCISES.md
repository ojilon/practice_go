# scraper — Polite Concurrent Fetcher — Exercises

> Work in THIS folder, `package scraper`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 fetch.go** — Fetch with 5s timeout, <title> extract (strings, no external deps).
- **E2 crawl.go** — concurrent crawl of local test server (httptest) — no internet required.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package scraper`.

## Build-chain (reuse + banned imports)
- MUST crawl concurrently with your `25_taskpool` pattern (import or mirror + cite), dedup URLs with your `16_set`, extract titles with your `26_stringsx` logic.
- Banned: external scrape frameworks (`colly` etc.) — `net/http` + your packages only. Test against `httptest` server (no internet needed).

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
