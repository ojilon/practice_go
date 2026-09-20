# httpraw — HTTP/1.1 From TCP — Exercises

> Work in THIS folder, `package httpraw`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)

- **E1 parse.go** — `type Request{Method, Path, RawQuery string; Proto string; Headers map[string]string; Body []byte; Params map[string]string}`, `func ParseRequest(r *bufio.Reader, maxHeader, maxBody int) (Request, error)`, errors `ErrBadRequest`, `ErrTooLarge`, `ErrUnsupportedTE`. Header names case-insensitive; obs-fold rejected.
- **E2 response.go** — `type ResponseWriter struct...` with `(w *ResponseWriter) Header() map[string]string`, `(w *ResponseWriter) WriteHeader(int)`, `(w *ResponseWriter) Write([]byte) (int, error)`, chunked-encode helper `WriteChunked(w io.Writer, chunks [][]byte) error`.
- **E3 router.go** — `type HandlerFunc func(*Request, *ResponseWriter)`, `type Middleware func(HandlerFunc) HandlerFunc`, `type Router struct...` with `(m *Router) Handle(method, pattern string, h HandlerFunc)`, `(m *Router) ServeHTTP(req *Request, w *ResponseWriter)`, `(m *Router) Use(mw Middleware)`.
- **E4 server.go** — `type ServerOptions{Addr string; ReadHeaderTimeoutMs, ReadTimeoutMs, WriteTimeoutMs, IdleTimeoutMs, MaxHeaderBytes, MaxBodyBytes int}`, `func DefaultOptions(addr string) ServerOptions`, `type Server struct...` with `func NewServer(opt ServerOptions, h HandlerFunc) *Server`, `(s *Server) ListenAndServe() error`, `(s *Server) Shutdown(ctx context.Context) error`, `(s *Server) Addr() string`.
- **E5 kv_service.go** — `func ServeKV(addr, dir string, opt ServerOptions) (*Server, error)` wiring `GET /health`, `GET /kv/:key`, `PUT /kv/:key`, `DELETE /kv/:key`, `GET /stats`; `func ServeFiles(prefix, dir string) HandlerFunc` with traversal rejection.
- **E6 demo_cli.go** — `func Demo() string`, `func DemoAttack() string`, `func DemoSlowClient() string`, `func DemoShutdown() string`, `func DemoLoad(n, concurrency int) string`, `func DemoCLI(args []string) int`.

## Edge cases (will be spot-checked)

- Bad request line / unknown method → 400; unknown path → 404; known path wrong method → 405 + `Allow`; oversize header → 431; oversize body → 413; negative/invalid Content-Length → 400; bad chunk → 400 + close; `..` traversal → 403/404 never 200.
- HEAD returns headers, no body. Keep-alive reuses conn; `Connection: close` or HTTP/1.0 closes. Handler panic → 500, conn stays usable for next keep-alive request.
- `Shutdown` waits for in-flight up to ctx deadline; new accepts refused after shutdown starts.
- No `fmt.Print*` in library code; exported symbols have `//` doc comments.
- 50-concurrent mixed load: no server panic, no hang, failures only documented 4xx.

## Suggested files

- `parse.go`, `response.go`, `router.go`, `server.go`, `kv_service.go`, `demo_cli.go`. Any layout OK as long as symbols exist.
- Package clause MUST be `package httpraw`.

## Build-chain (reuse + banned imports)

- MUST use your `18_logger`, `21_config`, `23_jsonlite` (/stats), `25_taskpool` (or documented per-conn goroutine + cite), `44_kv_store` backend. SHOULD cite `19_event`, `32_tokenizer`, `41_cli`.
- Banned in impl: `net/http` (any use), `gin/echo/chi/fasthttp`. Demos/tests MUST use a real client (`net/http` client or curl) to prove interop — that use stays in `demo_cli.go` only.
- Wrapping `http.Server`/`ServeMux` = FAIL.

## Stretch (optional, credited)

- `If-None-Match`/ETag on static files; `Range` single-suffix support.
- Access-log in Common Log Format + `routes` table printer.
- `Benchmark`-style latency histogram in `DemoLoad` (p50/p95/p99).
