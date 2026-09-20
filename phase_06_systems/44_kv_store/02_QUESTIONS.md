# kvstore — Durable KV — Questions (run code, then EDIT answers here)

> Type B (run-and-answer) unless marked Code-only. Fill each `> Your answer:` IN THIS FILE. Verifier checks this file + your `.go` code. `main.go` is ignored.

How to answer: implement the exercise, wire root `main.go` temporarily, `go run .`, paste output/numbers below.

---

### Q1: Reuse map — which own-packages + 43 did you use and where?

> Your answer:

TODO: e.g. 43_wal_seglog for log+Replay, 18_logger for compaction events, 20_cache hot path, 24_scheduler TTL sweeper, 25_taskpool background compact, 21_config LoadOptions.

### Q2: Reopen durability — paste `DemoReopen()` output (5k puts / 1k dels / survivors after reopen).

> Your answer:

```
TODO: paste program output here
```

### Q3: Compaction — paste `DemoCompact()` before/after (keys, live/dead bytes, segments) + survivors verified after reopen.

> Your answer:

```
TODO: paste program output here
```

### Q4: Bench — paste `DemoBench(20000)` output (puts/s, gets/s, p50/p99 sample, machine note).

> Your answer:

```
TODO: paste program output here
```

### Q5: Atomic batch + TTL — paste transcript proving (a) failed batch leaves store unchanged, (b) 50ms TTL expires.

> Your answer:

```
TODO: paste program output here
```

### Q6: Race run — command + result for the 8W×8R mixed workload.

> Your answer:

TODO: e.g. `go run -race .` calling DemoServeStyle-concurrent, no WARNING.

---

## Self-check before asking for verification

- [ ] All boxes above filled?
- [ ] `go vet ./phase_06_systems/44_kv_store/...` clean?
- [ ] `README.md` in this folder with arch diagram + durability contract + bench table?
