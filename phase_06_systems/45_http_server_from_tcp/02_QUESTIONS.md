# httpraw — HTTP/1.1 From TCP — Questions (run code, then EDIT answers here)

> Type B (run-and-answer) unless marked Code-only. Fill each `> Your answer:` IN THIS FILE. Verifier checks this file + your `.go` code. `main.go` is ignored.

How to answer: implement the exercise, wire root `main.go` temporarily, `go run .`, paste output/numbers below. Client checks REQUIRE a real client (curl or Go net/http client in demo).

---

### Q1: Reuse map — which own-packages did you use and where?

> Your answer:

TODO: e.g. 18_logger access log, 21_config options, 23_jsonlite /stats, 25_taskpool workers, 44_kv_store backend, 19_event request events.

### Q2: Interop proof — paste (a) `curl -v` GET /health tail (status + headers), (b) Go-client PUT then GET /kv/hello transcript.

> Your answer:

```
TODO: paste both transcripts here
```

### Q3: Attack demo — paste `DemoAttack()` output (garbage/huge-header/bad-chunk/CL+TE → your status codes). Confirm no panic/hang.

> Your answer:

```
TODO: paste program output here
```

### Q4: Slow client + shutdown — paste `DemoSlowClient()` (408 + close timing) and `DemoShutdown()` (in-flight drained) outputs.

> Your answer:

```
TODO: paste both outputs here
```

### Q5: Load — paste `DemoLoad(2000, 50)` output (req/s, failures, p50/p99 if measured).

> Your answer:

```
TODO: paste program output here
```

### Q6: Race run — command + result for the concurrent load.

> Your answer:

TODO: e.g. `go run -race .` load path, no WARNING.

---

## Self-check before asking for verification

- [ ] All boxes above filled (real transcripts, not invented)?
- [ ] `go vet ./phase_06_systems/45_http_server_from_tcp/...` clean?
- [ ] `README.md` in folder with subset table + timeout guidance + threat model?
