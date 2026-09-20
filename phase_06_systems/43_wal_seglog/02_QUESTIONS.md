# walseglog — Segmented WAL — Questions (run code, then EDIT answers here)

> Type B (run-and-answer) unless marked Code-only. Fill each `> Your answer:` IN THIS FILE. Verifier checks this file + your `.go` code. `main.go` is ignored.

How to answer: implement the exercise, wire root `main.go` temporarily, `go run .`, paste output/numbers below.

---

### Q1: Which of YOUR packages did you reuse? (import path or `// mirrors` cite)

> Your answer:

TODO: e.g. 18_logger for rotation events, 21_config for Options defaults, 34_io exact-read pattern.

### Q2: Paste `Demo()` output (must show append→read→reopen→replay count).

> Your answer:

```
TODO: paste program output here
```

### Q3: Kill / torn-tail test — how many records survived, how many truncated? Paste `DemoKill`-style output or describe the procedure + `Stats().TruncatedTails`.

> Your answer:

```
TODO: e.g. appended 50000 unsynced, reopened, survived=31200 truncated-tails=1
```

### Q4: Sync policy bench — `DemoBenchSync()` numbers (sync-on-write vs batched). What slowdown did YOU measure?

> Your answer:

```
TODO: e.g. sync-on-write 410 ops/s, batch-100 12800 ops/s (~31x)
```

### Q5: Race check — what command did you run for concurrent appends and what was the result? (Code-only if you show the demo func name; else paste `go run -race` tail)

> Your answer:

TODO: e.g. `go run -race .` with 8×500 appends, no WARNING, final count 4000.

### Q6: Record layout — draw your on-disk framing in 5 lines (magic/version/len/crc order). Why did you put CRC last?

> Your answer:

TODO: write 3-5 sentences here.

---

## Self-check before asking for verification

- [ ] All boxes above filled (or marked Code-only)?
- [ ] `go vet ./phase_06_systems/43_wal_seglog/...` clean?
- [ ] `DemoReuse44Style()` proves 44 can rebuild a map via `Replay`?
