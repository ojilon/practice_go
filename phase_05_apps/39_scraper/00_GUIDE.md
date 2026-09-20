# scraper — Polite Concurrent Fetcher — Guide (`scraper`)

> Folder: `phase_05_apps/39_scraper` · Package: `scraper` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Be polite: rate limit + timeout + UA; respect robots.txt delay note (no aggressive crawling).
- Step 1 — Fetch(url)(title string, links []string, err) via net/http + timeout.
- Step 2 — Crawl(seed, maxPages, delay) concurrent with YOUR taskpool logic (or WaitGroup+sem chan).
- Step 3 — context cancel + dedup via YOUR set.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package scraper` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_05_apps/39_scraper"` (package name `scraper`), call your funcs, `go run .`.
3. `go vet ./phase_05_apps/39_scraper/...` must be clean before moving on.

### Details
- Be polite: rate limit + timeout + UA; respect robots.txt delay note (no aggressive crawling).
- Step 1 — Fetch(url)(title string, links []string, err) via net/http + timeout.
- Step 2 — Crawl(seed, maxPages, delay) concurrent with YOUR taskpool logic (or WaitGroup+sem chan).
- Step 3 — context cancel + dedup via YOUR set.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_05_apps/39_scraper")
func main() { fmt.Println(scraper.Demo()) }
```
```powershell
go run .
go vet ./phase_05_apps/39_scraper/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
