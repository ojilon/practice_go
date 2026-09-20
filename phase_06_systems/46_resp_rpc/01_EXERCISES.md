# resprpc — RESP + Toy RPC — Exercises

> Work in THIS folder, `package resprpc`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)

- **E1 resp_codec.go** — `type ValueType int` (`SimpleString, Error, Integer, BulkString, Array, Null`), `type Value{Type ValueType; Str string; Int int64; Bulk []byte; Array []Value}`, `func Marshal(v Value) []byte`, `type Reader struct...` + `func NewReader(r *bufio.Reader, maxFrameBytes int) *Reader` + `func (rd *Reader) ReadValue() (Value, error)`, errors `ErrFrameTooLarge`, `ErrBadFrame`, `ErrTruncated`.
- **E2 resp_server.go** — `type RESPOptions{Addr, Dir string; MaxFrameBytes, MaxArgs int}`, `func DefaultRESPOptions(addr, dir string) RESPOptions`, `type RESPServer struct...` + `func ServeRESP(opt RESPOptions) (*RESPServer, error)` + `(s *RESPServer) Addr() string` + `(s *RESPServer) Close() error`. Commands PING/ECHO/SET/GET/DEL/INCR/EXISTS/COMMAND (subset documented).
- **E3 rpc.go** — `type Envelope{ID uint64; Method string; Payload []byte}`, `func MarshalEnvelope(e Envelope) []byte`, `func UnmarshalEnvelope(b []byte) (Envelope, error)`, `type RPCHandler func(ctx context.Context, payload []byte) ([]byte, error)`, `type RPCServer struct...` with `func NewRPCServer() *RPCServer` + `(s *RPCServer) Register(method string, fn RPCHandler) error` + `(s *RPCServer) ServeConn(conn net.Conn)` + `(s *RPCServer) ServeAddr(addr string) (net.Listener, error)`, `type RPCClient struct...` with `func DialRPC(addr string) (*RPCClient, error)` + `(c *RPCClient) Call(method string, payload []byte, timeoutMs int) ([]byte, error)` + `(c *RPCClient) Close() error`, errors `ErrUnknownMethod`, `ErrTimeout`, `ErrCorrupt`, `ErrClosed`.
- **E4 demo_cli.go** — `func Demo() string`, `func DemoRESPCodec() string`, `func DemoAtomicIncr() string`, `func DemoRESPBench(n, concurrency int) string`, `func DemoRPC() string`, `func DemoRaftShapedRPC() string`, `func DemoCLI(args []string) int`.

## Edge cases (will be spot-checked)

- `$-1` / `*-1` nulls parse; bulk with embedded `\r\n` round-trips; truncated frame → `ErrTruncated` (never hang); oversize bulk → `ErrFrameTooLarge` mid-stream without OOM.
- Well-framed unknown RESP command → `-ERR unknown command`, conn stays open. Unknown RPC method → `ErrUnknownMethod`, conn stays open. RPC timeout → `ErrTimeout`, conn reusable.
- `INCR` on non-integer → `-ERR value is not an integer`; 20×500 concurrent INCRs land exact.
- No `fmt.Print*` in library code; exported symbols have `//` doc comments.

## Suggested files

- `resp_codec.go`, `resp_server.go`, `rpc.go`, `demo_cli.go`. Any layout OK as long as symbols exist.
- Package clause MUST be `package resprpc`.

## Build-chain (reuse + banned imports)

- MUST use `44_kv_store`, `25_taskpool`, `18_logger`, `21_config`. SHOULD cite `23_jsonlite`, `15_queue`/`17_ringbuffer`, `41_cli`.
- Banned in impl: `go-redis/miniredis/redcon`, `net/rpc`, `grpc`, `encoding/gob`, `encoding/json` (wire path). Demos may use raw `net.Conn` + (for proof only) compare against stdlib in isolation.
- 47-seam: keep Envelope + RPCServer/RPCClient signatures stable; `DemoRaftShapedRPC` proves `raft.requestVote`/`raft.appendEntries` round-trip.

## Stretch (optional, credited)

- `SUBSCRIBE`-lite pub/sub fan-out via your `19_event` bus.
- RPC streaming call (server→client pushes on one call ID).
- RESP3 attribute type passthrough (parse + re-emit, documented subset).
