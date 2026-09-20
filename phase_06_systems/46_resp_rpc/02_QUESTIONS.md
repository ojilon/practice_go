# resprpc — RESP + Toy RPC — Questions (run code, then EDIT answers here)

> Type B (run-and-answer) unless marked Code-only. Fill each `> Your answer:` IN THIS FILE. Verifier checks this file + your `.go` code. `main.go` is ignored.

How to answer: implement the exercise, wire root `main.go` temporarily, `go run .`, paste output/numbers below.

---

### Q1: Reuse map — which own-packages did you use and where?

> Your answer:

TODO: e.g. 44_kv_store backend, 25_taskpool handlers, 18_logger, 21_config options, 23_jsonlite RPC payloads.

### Q2: Codec — paste `DemoRESPCodec()` output (1000 round-trips, mismatches must be 0; include tricky cases: embedded CRLF, empty, null, nested array).

> Your answer:

```
TODO: paste program output here
```

### Q3: Interop — paste raw-bytes transcript (request bytes → reply bytes) proving pipelining (≥3 commands in one write, 3 replies in order). State the client used (nc / Go net.Conn / other).

> Your answer:

```
TODO: paste bytes transcript here
```

### Q4: Atomic INCR + benches — paste `DemoAtomicIncr()` (20×500 → exact final) and `DemoRESPBench(20000,20)` ops/s + one-sentence HTTP(45)-vs-RESP comparison.

> Your answer:

```
TODO: paste program output here
```

### Q5: RPC — paste `DemoRPC()` (50×200 multiplex, 0 mixups, timeout case reuses conn) and `DemoRaftShapedRPC()` (requestVote/appendEntries round-trip) outputs.

> Your answer:

```
TODO: paste program output here
```

### Q6: Race run — command + result for the concurrent INCR/RPC demos.

> Your answer:

TODO: e.g. `go run -race .` concurrent path, no WARNING.

---

## Self-check before asking for verification

- [ ] All boxes above filled (real transcripts)?
- [ ] `go vet ./phase_06_systems/46_resp_rpc/...` clean?
- [ ] `README.md` in folder with wire diagrams + error-vs-close table + limits?
