# Core Note — Goroutines

> Prerequisite for: `phase_01_core/09_goroutines`. Used by: `scheduler`, `taskpool`, `event` bus, `scraper`, `http_api` server, `tui` tick loop.

## Essentials
- `go f()` starts a lightweight concurrent function. Main exiting kills all — coordinate shutdown.
- **`sync.WaitGroup` pattern:** `wg.Add(1)` BEFORE `go`, `defer wg.Done()` inside, `wg.Wait()` after loop.
- Shared memory needs sync: `sync.Mutex` around counter/map, or better, share via channels (next note).
- `go run -race .` detects data races — run it on every concurrent exercise; a clean race report is part of DONE.

## Gotchas
- Loop-var capture: `for _, u := range urls { go func() { fetch(u) }() }` — pass `u` as a param (`go func(u string){...}(u)`) unless you know your Go version's per-iteration semantics.
- NEVER sync with `time.Sleep` ("wait 100ms hoping worker finished") — use WaitGroup/done-channel/Future.
- Goroutine leaks: every launched goroutine must have an exit path (done channel, context cancel, pool Shutdown).

## Building-block role
- Worker pools = N goroutines × job channel. Schedulers = goroutine × Ticker. Scrapers = goroutines × semaphore + WaitGroup. All the same three primitives.
