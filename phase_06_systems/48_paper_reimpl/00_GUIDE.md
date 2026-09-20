# 48 — Paper Re-implementation Capstone — Guide (`paperreimpl`)

> Folder: `phase_06_systems/48_paper_reimpl` · Package: `paperreimpl` · Run via root `main.go` (verifier ignores `main.go`).
> Time budget: **3–4 weeks** (~30–45 hours). Pick ONE track and implement it at evaluation quality. This closes Phase 6 and pre-arms Phase 7 (your store/transport lessons) and Phase 11 (this is your first capstone template).

## Why it exists

Reading a paper is tourism; re-implementing its core and *measuring where it breaks* is residency. Real distributed careers are built on 3–4 papers understood this deeply (Dynamo, MapReduce, GFS, Raft — you already did Raft in 47). This capstone forces the full loop: read primary source → write structured notes → build the minimal faithful core on YOUR stack → evaluate with failure + scale numbers → document what the paper glossed over.

Pick the track that excites you — all three are real-life systems, all reuse ≥6 of your packages, all take ~a month done properly.

## Tracks (pick EXACTLY ONE — state it in Q1)

### Track A — Dynamo-lite (recommended if you loved 47's contrast)

Dynamo (DeCandia et al., SOSP'07): consistent hashing ring + virtual nodes, vector clocks, sloppy quorum (R/W/N), hinted handoff, gossip membership + anti-entropy (Merkle-lite).
Build: N=5 in-process nodes fronted by your 46 RPC, coordinator routing, `Put(k,v)`/`Get(k)` with clock reconciliation (siblings returned on conflict — client resolves, exactly like Dynamo), 1-node-down write still succeeds via handoff, healed node catches up via anti-entropy.
Evaluate: quorum sweep (R+W vs N curves you measure), down-node availability proof, sibling-conflict demo with concurrent writers.

### Track B — MapReduce-lite (recommended if you loved 25/39)

MapReduce (Dean & Ghemawat, OSDI'04): master/worker over your 46 RPC + 25 taskpool, map → shuffle (partition by hash) → reduce, fault tolerance by re-execution.
Build: `WordCount` + `InvertedIndex` jobs over real files (your 38 parsers feed splits), worker crash mid-job re-runs without wrong output, straggler re-execution stretch.
Evaluate: 1-worker vs 4-worker speedup on ≥20MB text, kill-1-worker-mid-job still-correct proof, shuffle-size vs time note.

### Track C — GFS-chunk-lite (recommended if you loved 43/44)

GFS (Ghemawat et al., SOSP'03) §2–3: single master (namespace + chunk map + leases), chunkservers storing 64KB-lite chunks (your 43 log or flat files), client write with primary-lease ordering, master failover from persisted state.
Build: `Create/Write/Read/Delete` files striped into chunks across 3 chunkservers, primary lease per chunk, checksum per chunk (corrupt one → detected + re-replicated), master restart recovers namespace.
Evaluate: 50MB file write/read throughput, kill-primary-mid-write recovery, 1-chunkserver-down read still works.

## Study first (primary sources — secondary blogs don't count)

- Your track's paper, FULL read: Dynamo (all 9 pages) / MapReduce (§1–5) / GFS (§1–4). No skimming the evaluation section — the evaluation is the point.
- The paper's *related work* section: one paragraph in your notes on what the authors rejected and why (this becomes your Q6 answer).
- Two implementations' design notes (not code): e.g. Riak's Dynamo notes / Hadoop's MapReduce docs / HDFS architecture guide. Steal operational wisdom, not code.
- Your 43/44/46/47 READMEs — the seams you promised must hold under this capstone's load.

Write `PAPER_NOTES.md` (your file, ~80 lines, REQUIRED): problem → key idea → core algorithm (your own words + one diagram) → consistency/fault model → evaluation claims → what you implement vs deliberately cut → what surprised you. Quote it in Q6.

## Milestones (all tracks share the shape — 4 weeks)

### M1 — Notes + thin slice (Week 1)

- `PAPER_NOTES.md` done + `Track() string` returns `"dynamo" | "mapreduce" | "gfs"` (verifier checks you committed to one).
- Thin end-to-end slice on localhost: Dynamo `put→get` 1 key through coordinator; MapReduce wordcount on 5 lines with 1 worker; GFS create+write+read 1 chunk. Ugly is fine — the path must exist.
- Interface freeze: your public API for the track (`Dynamo.Put/Get`, `JobRunner.Run`, `GFSClient.Write/Read`) is declared in `types.go` and unchanged after M1 without a note.

### M2 — Faithful core (Weeks 1–2)

- Dynamo: ring + vnodes + preference list, vector-clock compare (`Before/Concurrent/ Merge`), R/W quorum paths, sibling set return, hinted handoff store + forward on heal, gossip liveness (even crude round-robin heartbeats count if measured).
- MapReduce: split → map (taskpool) → partitioned shuffle files → reduce, master state machine (idle/in-progress/completed per task), worker heartbeat + timeout re-queue, deterministic output (same input → byte-identical output across runs).
- GFS: namespace + chunk-map + lease manager on master (persisted via 43 or snapshot file), chunkservers with per-chunk CRC, client library with lease-ordered writes, master-restart recovery demo.
- Correctness demos per track (verifier spot-checks the matching one): `DemoConflict()` (siblings) / `DemoFaultTolerance()` (worker-kill) / `DemoFailover()` (master-restart) — each prints PASS/FAIL + numbers, not vibes.

### M3 — Evaluation + failure injection (Week 3)

- Numbers, not adjectives. Minimum per track:
  - Dynamo: quorum sweep table (3 configs × puts/s + gets/s + conflict rate), 1-down availability (`DemoDownNode()` — writes succeed, reads succeed, handoff backlog drains on heal with counts).
  - MapReduce: speedup table (1/2/4 workers × seconds on same corpus + shuffle bytes), worker-kill-mid-job proof (job still byte-correct + extra seconds quantified).
  - GFS: 50MB write/read MB/s, primary-kill-mid-write recovery (no half-chunk visible), 1-chunkserver-down availability.
- Failure injection uses YOUR chaos helpers' style from 47 (kill/partition/heal funcs — reuse the pattern, cite it). Every failure demo ends with a consistency assert (sibling check / byte-compare / checksum scrub), never just "it didn't crash".
- Bench + failure outputs pasted into QUESTIONS. Invented numbers are obvious (they won't match your machine's scale) — run them.

### M4 — Ship the capstone (Week 4)

- `DemoCLI(args) int`: per-track serve/bench/fail commands (`dynamo-serve`, `dynamo-bench`, `dynamo-kill-test` / `mr-run`, `mr-bench`, `mr-kill-worker` / `gfs-serve`, `gfs-bench`, `gfs-kill-primary`). Same CLI shape as 41/44/45/46/47.
- `README.md` (your file, the Phase-6 crown doc): paper → what you built (arch diagram) → subset table (paper feature × in/out × why) → evaluation tables from YOUR machine → failure table → reuse map (≥6 own packages named) → "what I'd do with 4 more weeks".
- Phase-6 retrospective `RETRO.md` (short, 20 lines): which of 43–47 paid off most, what API seam broke, what you'd redesign before Phase 7's mini-git storage layer. This file is the bridge to Phase 7 — write it like your future self will read it (they will).

## Reuse (required — the heaviest in Phase 6)

- MUST reuse ≥6 own packages. Natural map: `43_wal_seglog` (Dynamo handoff log / MR task log / GFS master log) + `44_kv_store` (Dynamo storage engine / GFS chunk store) + `46_resprpc` (all RPC) + `18_logger` + `21_config` + `25_taskpool` + (`24_scheduler` heartbeats/gossip) + (`28_hashmap` ring index / `23_jsonlite` payloads / `41_cli` CLI). Name all six in Q1 with one line each on *where*.
- SHOULD run the demo service behind 45 HTTP or 46 RESP (one transcript proves it).
- <6 documented reuses → PARTIAL regardless of code quality.

## Banned imports (in impl)

The paper's reference systems (`riak`, `hadoop`, `hdfs` clients), consensus/store libs from 47's ban list, `net/http` as the *internal* RPC (edge serving via 45 is fine), `encoding/gob` on wire/store paths. `net`, `sync`, `time`, `context`, `os`, `sort`, `hash/crc32`, `math/rand` fine.

Importing a ready-made Dynamo/MR/GFS = FAIL.

## How to run

```go
// root main.go (TEMPORARY, verifier ignores it)
package main

import (
	"fmt"
	"practice_go/phase_06_systems/48_paper_reimpl"
)

func main() { fmt.Println(paperreimpl.Demo()) }
```

```powershell
go run .
go vet ./phase_06_systems/48_paper_reimpl/...
```

## Done when (Phase 6 done when THIS is done)

- One track committed (`Track()`), `PAPER_NOTES.md` + `README.md` + `RETRO.md` exist with the content above.
- Track's correctness demo prints PASS with numbers; evaluation tables (perf + failure) pasted from your machine; consistency assert holds after every failure injection.
- ≥6 reuses documented; `go vet` clean; `-race` on the failure demo clean (state command in Q7).
- You can demo the whole thing live in ≤5 minutes (`Demo()` + one CLI command) — practice that run once and note the commands in README.
