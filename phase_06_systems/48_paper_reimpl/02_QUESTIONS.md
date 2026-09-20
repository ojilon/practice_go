# paperreimpl — Paper Capstone — Questions (run code, then EDIT answers here)

> Type B (run-and-answer) unless marked Code-only. Fill each `> Your answer:` IN THIS FILE. Verifier checks this file + your `.go` code. `main.go` is ignored.

How to answer: implement the exercise, wire root `main.go` temporarily, `go run .`, paste output/numbers below. Numbers must come from YOUR machine.

---

### Q1: Track + reuse map — which track did you pick and which ≥6 own-packages did you reuse (one line each: where)?

> Your answer:

TODO: e.g. Track=dynamo. 43 handoff log, 44 storage engine, 46 RPC, 18_logger, 21_config, 25_taskpool map workers, 24_scheduler gossip ticks.

### Q2: Paper in your words — what problem, what key idea, what did you cut and why? (Quote PAPER_NOTES.md, 5–8 sentences.)

> Your answer:

TODO: write 5-8 sentences here.

### Q3: Correctness proof — paste the matching demo output: `DemoConflict()` (siblings shown) XOR `DemoFaultTolerance()` (byte-identical after worker kill) XOR `DemoFailover()` (namespace intact after master restart).

> Your answer:

```
TODO: paste program output here
```

### Q4: Performance evaluation — paste `DemoEval()` table (your machine numbers: throughput/speedup/MB/s + corpus size + worker/node counts).

> Your answer:

```
TODO: paste eval table here
```

### Q5: Failure matrix — paste `DemoFailure()` table (scenario × outcome × recovery-ms × DataIntact). Every row must assert data intactness, not just liveness.

> Your answer:

```
TODO: paste failure table here
```

### Q6: Related work — what did the paper's authors reject and why? What surprised you most on re-implementation? (2–4 sentences.)

> Your answer:

TODO: write 2-4 sentences here.

### Q7: Race + live demo — (a) race command + result on the failure path, (b) the exact 2–3 commands for your ≤5-minute live demo.

> Your answer:

TODO: e.g. `go run -race .` failure path clean; demo: `go run .` → DemoCLI bench → kill-test.

---

## Self-check before asking for verification

- [ ] All boxes above filled (real numbers)?
- [ ] `go vet ./phase_06_systems/48_paper_reimpl/...` clean?
- [ ] `PAPER_NOTES.md` + `README.md` + `RETRO.md` exist in folder with guide-stated content?
- [ ] `ReuseMap()` names ≥6 own packages?
