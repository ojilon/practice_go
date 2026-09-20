# raftlite — Raft-lite — Questions (run code, then EDIT answers here)

> Type B (run-and-answer) unless marked Code-only. Fill each `> Your answer:` IN THIS FILE. Verifier checks this file + your `.go` code. `main.go` is ignored.

How to answer: implement the exercise, wire root `main.go` temporarily, `go run .`, paste output/numbers below. Chaos MUST pass 3 consecutive runs — paste all three or state variance.

---

### Q1: Reuse map — which own-packages + 43/46 did you use and where?

> Your answer:

TODO: e.g. 43_wal_seglog entry persistence+Replay, 46_resprpc vote/append transport, 18_logger term transitions, 21_config node opts, 25_taskpool apply loop, 44 shapes for demo SM.

### Q2: Election — paste `DemoElection()` (20 runs, leaders/terms) + `DemoStepDown()` output. Exactly one leader every run?

> Your answer:

```
TODO: paste program output here
```

### Q3: Replication — paste `DemoReplicate(500)` output (3 maps identical? commit p50/p99?) + `DemoCatchup()` (kill follower → restart → gap-free?) outputs.

> Your answer:

```
TODO: paste program output here
```

### Q4: TCP + bench — paste `DemoTCPCluster()` (200 cmds over sockets) and `DemoBenchRaft(2000)` (submits/s, p50/p99) + one-sentence cost-vs-single-node-44 comparison.

> Your answer:

```
TODO: paste program output here
```

### Q5: Chaos table — paste `DemoChaos()` 4-scenario table (kill-leader recovery time, minority-partition availability, majority-partition no-phantom-commit, stale-leader rejoin prefix intact). Confirm 3 consecutive passes.

> Your answer:

```
TODO: paste chaos table × 3 runs here
```

### Q6: Paper in 3 sentences — state Raft's election-safety argument (why two leaders can't be elected in the same term) + your timeout ratio choice and why.

> Your answer:

TODO: 3-6 sentences quoting your RAFT_NOTES.md.

### Q7: Race run — command + result for the chaos/replication path.

> Your answer:

TODO: e.g. `go run -race .` chaos path, no WARNING.

---

## Self-check before asking for verification

- [ ] All boxes above filled (real numbers, 3× chaos)?
- [ ] `go vet ./phase_06_systems/47_raft_lite/...` clean?
- [ ] `RAFT_NOTES.md` + `README.md` in folder (subset table, tuning guide, chaos table)?
