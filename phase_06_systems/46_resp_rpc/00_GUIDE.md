# 46 — Binary & Text Protocols: RESP + Toy RPC — Guide (`resprpc`)

> Folder: `phase_06_systems/46_resp_rpc` · Package: `resprpc` · Run via root `main.go` (verifier ignores `main.go`).
> Time budget: **1.5–2.5 weeks** (~15–25 hours). The transport 47's Raft RPC will ride on.

## Why it exists

HTTP is readable but wasteful. Real systems speak framed binary or tight text protocols: Redis RESP, Postgres wire, gRPC/HTTP2 frames, Raft RPCs. This module teaches the two skills every protocol engineer needs: **(1) streaming parse with hard frame limits** (never trust the peer's length prefix), and **(2) request/response multiplexing with timeouts** (never let one slow peer stall the world). You prove both by fronting your 44 KV with a Redis-flavoured protocol *and* a tiny typed RPC that 47 reuses for `RequestVote`/`AppendEntries`.

Real-life payoff: you can add a wire protocol to anything (KV, queue, scheduler), attack it with truncated/garbage frames, and explain pipelining vs multiplexing from measurements, not blog posts.

## What you will learn

- RESP2 subset: `+ - : $ *` types, bulk-string streaming with `MaxFrameBytes` guard, inline-command fallback, pipelined request handling.
- Length-prefixed toy RPC: `Envelope{ID, Method, Payload}` over any `net.Conn`, little-endian header you document, handler registry, client with in-flight map + per-call deadline.
- Server discipline: per-conn read loop, handler execution on your `25_taskpool` (slow handler never blocks other conns), write serialization per conn, graceful close with in-flight drain.
- Interop proof: real `redis-cli`-style bytes (or `nc`/`socat` transcript) against your RESP front; Go client pipelining 10k commands and reporting throughput.
- Failure behaviour: truncated frame → error + close (never hang); oversize frame → error reply + close; unknown method → typed error, server stays up.

## Study first

1. Redis RESP specification (redis.io/docs/reference/protocol-spec) — read the full short spec; implement only the subset below.
2. Raft paper §5 (leader election RPCs) + §7 (joint consensus skim) — note the *exact* RPC fields Raft needs; your `Envelope` must be able to carry them in 47.
3. Go `bufio.Reader`/`Scanner` split-function docs + `encoding/binary` — framing patterns.
4. Your 45 `ParseRequest` — compare text-protocol error handling (431/413/400+close) with what you do here (error reply vs close — document the rule).

## Milestones (do in order)

### M1 — RESP codec + streaming reader (Week 1, days 1–4)

- `Value{Type, Str, Int, Bulk []byte, Array []Value}`, `Marshal(v) []byte`, `Reader{...}` with `ReadValue() (Value, error)` streaming from `bufio.Reader` (no `io.ReadAll` on the conn — unbounded read = FAIL).
- `MaxFrameBytes` (default 1 MiB bulk, 16k args) enforced *during* read; oversize → `ErrFrameTooLarge`, never OOM on hostile `$-1`-then-gigabytes input.
- Commands: `PING`, `ECHO`, `SET k v`, `GET k`, `DEL k [k...]`, `INCR k`, `EXISTS`, `COMMAND DOCS`-lite (`COMMAND` returns your supported list). Errors are RESP `-ERR ...`, never a dropped connection for a well-framed unknown command.
- `DemoRESPCodec()` round-trips 1k random values (bulk with `\r\n` inside, empty bulk, null bulk `$-1`, nested arrays) and prints mismatches (must be 0). This is the verifier's spot-check — keep it deterministic with a fixed seed.

### M2 — RESP server fronting the KV (Week 1–2)

- `RESPServer{Addr, Store *kvstore.Store, ...}`, `Serve(addr, dir)`, per-conn loop supporting **pipelining** (5 commands in one write get 5 replies in order). Slow `SET` handler never blocks other connections (taskpool or documented worker model).
- `INCR` is atomic under concurrency (prove with 20 clients × 500 incrs → exact final value in `DemoAtomicIncr()`).
- Interop transcript: drive the server with raw bytes (`printf 'SET ...\r\n' | nc` or a Go `net.Conn` client — either is fine) and paste request/response bytes into QUESTIONS. At least one transcript must show pipelined commands.
- Bench: `DemoRESPBench(n=20000, c=20)` pipelined SET/GET mix → ops/s. Compare against your 45 HTTP KV numbers in one sentence (which is faster on your machine and why you think so).

### M3 — Toy RPC for Raft (Week 2)

- `Envelope{ID uint64; Method string; Payload []byte}`, `MarshalEnvelope/UnmarshalEnvelope` with `magic(2) | version(1) | id(8) | methodLen(2) | method | payloadLen(4) | payload | crc32(4)`. CRC covers method+payload; mismatch → `ErrCorrupt` + close.
- `RPCServer`: `Register(method string, fn func(ctx, payload []byte) ([]byte, error))`, `ServeConn(conn)`, `ServeAddr(addr)`; `RPCClient`: `Call(method, payload, timeoutMs) ([]byte, error)` with concurrent in-flight calls multiplexed over one conn (ID map + response dispatcher goroutine). Unknown method → `ErrUnknownMethod` typed error, conn stays open. Timeout → `ErrTimeout`, conn stays open, late reply is dropped safely.
- `DemoRPC()` runs server + 50 concurrent clients × 200 calls (echo + kv-put-through-RPC) with 0 mixups (ID routing proof) + one timeout case (handler sleeps 200ms, client timeout 50ms → `ErrTimeout`, next call on same conn still works).
- 47-seam proof: `DemoRaftShapedRPC()` registers `raft.requestVote` and `raft.appendEntries` methods carrying `encoding/binary`-or-jsonlite payloads of your 47 struct shape (even stub bytes are fine) and round-trips them — proves 47 can ride this transport without forking it.

### M4 — Ship it

- `DemoCLI(args) int`: `resp-serve --addr --dir`, `resp-bench`, `rpc-demo`, `rpc-bench`. Same CLI shape as 41/44/45.
- `README.md` (your file): wire diagrams (RESP example bytes + RPC header layout), error-vs-close policy table, pipelining vs multiplexing note with your bench numbers, limits (`MaxFrameBytes`, max args, max method length), what `redis-cli` can/can't do against you.
- Keep `Value/Marshal/Reader/RESPServer/RPCServer/RPCClient/Envelope` API stable — 47 imports this package.

## Reuse (required)

- MUST use `44_kv_store` (RESP backend + RPC kv-passthrough), `25_taskpool` (handler execution), `18_logger`, `21_config` (server options). SHOULD use `23_jsonlite` (RPC payload encoding option), `15_queue`/`17_ringbuffer` (per-conn write queue cite), `41_cli` patterns.
- Document in Q1.

## Banned imports (in impl)

Any Redis client/server lib (`go-redis`, `miniredis`, `redcon`), any RPC framework (`net/rpc`, `grpc`, `gob`-RPC). `net`, `bufio`, `sync`, `encoding/binary`, `hash/crc32` expected. `encoding/gob` and `encoding/json` banned on the wire path (demos may compare with them in isolation).

Wrapping `net/rpc` or embedding a Redis server = FAIL.

## How to run

```go
// root main.go (TEMPORARY, verifier ignores it)
package main

import (
	"fmt"
	"practice_go/phase_06_systems/46_resp_rpc"
)

func main() { fmt.Println(resprpc.Demo()) }
```

```powershell
go run .
go vet ./phase_06_systems/46_resp_rpc/...
```

## Done when

- Codec round-trip 0 mismatches; RESP server passes pipelining + atomic-INCR + raw-bytes interop; RPC passes 50×200 multiplex + timeout-reuse + Raft-shaped methods.
- Bench numbers (RESP ops/s, RPC calls/s) pasted from your machine + one-sentence HTTP-vs-RESP comparison.
- `go vet` clean; `-race` on the concurrent demos clean (state command in Q6).
