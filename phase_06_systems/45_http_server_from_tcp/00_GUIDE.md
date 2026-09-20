# 45 — HTTP/1.1 Server From Raw TCP — Guide (`httpraw`)

> Folder: `phase_06_systems/45_http_server_from_tcp` · Package: `httpraw` · Run via root `main.go` (verifier ignores `main.go`).
> Time budget: **2–3 weeks** (~20–30 hours). You will never again say "the framework handles it" without knowing what *it* is.

## Why it exists

`net/http` hides sockets, parsing, keep-alive, timeouts, and slow-client attacks. Real systems engineers have debugged all five at 2am. This module forces that education in a safe dojo: a small but genuinely usable HTTP/1.1 server built on `net.Conn` that can serve your 44 KV store, pass a subset of real HTTP client checks (`curl`, Go's `net/http` client, Python `urllib`), and survive hostile inputs.

Real-life payoff: you can reason about request smuggling, header limits, keep-alive vs close, graceful shutdown, and why every production server has read/write/idle timeouts — because you implemented each one and attacked your own server.

## What you will learn

- HTTP/1.1 request parsing from a stream: request line, headers (case-insensitive, folded-line rejection), `Content-Length` vs `Transfer-Encoding: chunked`, keep-alive connection reuse.
- Response writing: status line, headers, exact-length vs chunked vs close-delimited bodies, HEAD/GET/POST/PUT/DELETE semantics you actually support.
- Routing + middleware: exact + `:param` + prefix routes, `Use()` chain (logging via your logger, request-id via your event bus or counter, panic-recovery, per-route timeout).
- Production hardening: `ReadHeaderTimeout`, `ReadTimeout`, `WriteTimeout`, `IdleTimeout`, `MaxHeaderBytes`, `MaxBodyBytes` (413/431/408 statuses), slow-loris resistance you *demonstrate*.
- Operating it: graceful `Shutdown(ctx)` draining in-flight requests, access log shape, `/health` + `/stats` endpoints, static-file handler with path-traversal rejection, KV-backed CRUD demo.

## Study first

1. RFC 9110 (HTTP Semantics) §9–10 + RFC 9112 (HTTP/1.1) §3–7 — read the request/response framing sections, not the whole RFC. Take notes on what you will *deliberately not* support.
2. Go `net/http` server source (skim `server.go`: `readRequest`, `serve`, timeout handling) + `net/textproto` reader — steal timeout and state-machine ideas, not code.
3. "Slowloris" + request-smuggling explainers (any two good articles) — you must implement at least the mitigations you can test.
4. Your 37 `httpapi` handlers — port at least one route to this raw server so the reuse loop is visible.

## Milestones (do in order)

### M1 — Parser + single-shot server (Week 1, days 1–4)

- `ParseRequest(r *bufio.Reader) (Request, error)` handles `GET/POST/PUT/DELETE/HEAD`, query strings, headers (limit `MaxHeaderBytes`, reject obs-fold with 400), `Content-Length` bodies up to `MaxBodyBytes` (else 413). Chunked decoding may start as "return 501" if you document it — but M2 must finish it.
- `Server{Addr, Handler, Timeouts...}`, `ListenAndServe()`, per-connection goroutine, one request per connection is enough for M1.
- `ResponseWriter` with `WriteHeader(status)`, `Write([]byte)`, `Header() map`. Correct `Content-Length` on fixed bodies, `Date` + `Server: practice-go/1.0` headers, proper `400/404/413/500` paths.
- Prove with real clients: `curl -v localhost:PORT/health` + Go `net/http` client GET/POST transcript pasted into QUESTIONS. `curl` succeeding against your hand-rolled parser is the M1 thrill — don't skip it.

### M2 — Keep-alive, chunked, router, middleware (Week 1–2)

- Keep-alive: `Connection: keep-alive/close` + `HTTP/1.0` without keep-alive → close; `IdleTimeout` closes quiet keep-alive conns; pipelined requests on one conn work sequentially (no concurrency per conn).
- `Transfer-Encoding: chunked` request decoding + chunked response encoding when handler writes without length. Malformed chunks → 400 and close.
- `Router`: `Handle(method, pattern, HandlerFunc)` with `:name` params (`/kv/:key`) + `Prefix()` static helper; `Query()` accessor; 404 vs 405 distinction (wrong method on known path → 405 + `Allow` header).
- `Use(mw Middleware)` chain: at minimum request logging (your `18_logger`), request-id header (`X-Request-ID`), recover-from-panic → 500. Per-request `context` with timeout helper `WithTimeout(ms)`.
- Fuzz-ish robustness: `DemoAttack()` sends garbage (huge header, bad chunk, negative Content-Length, smuggled `CL+TE` combo) and prints your status codes — all must be 4xx + connection handled, never a server panic or hang.

### M3 — Harden + serve the KV (Week 2–3)

- Timeouts enforced with `net.Conn.SetDeadline`: `ReadHeaderTimeoutMs` (default 5s — the slow-loris guard), `ReadTimeoutMs`, `WriteTimeoutMs`, `IdleTimeoutMs`. `DemoSlowClient()` proves a 10-byte-drip header gets a 408 + close instead of holding a worker forever. Paste timings.
- `Shutdown(ctx)` : stops accepting, waits for in-flight (up to ctx deadline), then closes. Prove with in-flight-sleep handler test (`DemoShutdown()` transcript).
- KV service: `ServeKV(addr, *kvstore.Store)` exposing `GET /health`, `GET /kv/:key`, `PUT /kv/:key` (raw body), `DELETE /kv/:key`, `GET /stats` (JSON via your `23_jsonlite`). Concurrency via your `25_taskpool` (bounded workers) or documented goroutine-per-conn with cap — load it with 50 concurrent `curl`-style clients in `DemoLoad()` and report req/s + failures.
- Static files: `ServeFiles(prefix, dir)` with `..` traversal rejection (test `GET /static/../../go.mod` → 403/404, never 200), content-type sniff-lite, `If-None-Match`-optional stretch.

### M4 — Ship it (Week 3)

- `DemoCLI(args) int`: `serve --addr --dir` (serves KV on disk), `bench --n --c` (concurrent GET/PUT mix through real TCP), `routes` (print table). Same CLI shape as 41/44.
- `README.md` (your file): supported subset table (method × feature → status), timeout guidance, threat model ("what I mitigate / what I don't: no TLS, no h2, no TE+CL dual — I reject"), bench numbers, curl transcripts.
- Seam: 46 may frame RPC over your connection handling; 48 may benchmark against this server. Keep `Server/Router/ParseRequest/ServeKV` signatures stable.

## Reuse (required)

- MUST use `18_logger` (access + error logs), `21_config` (server options), `25_taskpool` (or justify goroutine-per-conn + cite `09_goroutines`/`10_channels`), `23_jsonlite` (`/stats` JSON), `44_kv_store` (`ServeKV` backend).
- SHOULD reuse `19_event` (request-id / request-completed events), `32_tokenizer`-style header scanning discipline, `41_cli` patterns for `DemoCLI`.
- Document in Q1.

## Banned imports (in impl)

`net/http` (server AND client in impl — tests/demos may use `net/http` *client* or `curl` to check you, and that is required), any web framework (`gin`, `echo`, `chi`, `fasthttp`). `net`, `net/textproto` (allowed as a reader helper — cite it), `bufio`, `context`, `sync`, `time` are fine.

Wrapping `http.Server` or `http.ServeMux` = FAIL.

## How to run

```go
// root main.go (TEMPORARY, verifier ignores it)
package main

import (
	"fmt"
	"practice_go/phase_06_systems/45_http_server_from_tcp"
)

func main() { fmt.Println(httpraw.Demo()) }
```

```powershell
go run .
go vet ./phase_06_systems/45_http_server_from_tcp/...
```

## Done when

- All `01_EXERCISES.md` symbols exist; real `curl` + Go-client transcripts prove GET/POST/keep-alive/chunked work.
- Attack demo shows only 4xx + no panic/hang; slow-client demo shows 408 + close; shutdown demo drains in-flight.
- KV-backed service handles 50-concurrent load with 0 unexpected failures (paste `DemoLoad` numbers).
- `go vet` clean; `-race` load run clean (state command in Q6).
