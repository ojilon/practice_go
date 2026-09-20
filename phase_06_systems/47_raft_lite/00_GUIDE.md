# 47 — Raft-lite: Election + Log Replication — Guide (`raftlite`)

> Folder: `phase_06_systems/47_raft_lite` · Package: `raftlite` · Run via root `main.go` (verifier ignores `main.go`).
> Time budget: **3.5–4.5 weeks** (~35–50 hours). The hardest module in Phase 6. Expect to stare at timers and stare back. That staring *is* the learning.

## Why it exists

Raft ("In Search of an Understandable Consensus Algorithm", Ongaro & Ousterhout) is the paper every distributed engineer has implemented once badly and then understood forever. Your simplified version — leader election, log replication, persistence, commit/apply, and chaos testing — turns "majority quorum" from vocabulary into muscle memory. It rides directly on your 43 log and your 46 RPC, and it replicates commands into a 44-style state machine. Phase 48's Dynamo track then shows the *other* philosophy (leaderless quorum) so you can argue both sides.

Real-life payoff: a 3-node cluster on your laptop that elects a leader, survives a kill, survives a partition, and keeps a replicated KV consistent — with the election/replication traces to prove it. That demo gets you hired; the paper notes get you respected.

## What you will learn

- Raft core: terms, votedFor, leader/follower/candidate states, randomized election timeouts, RequestVote + AppendEntries semantics (including the *log-consistency check* that makes Raft correct).
- Persistence that matters: currentTerm + votedFor + log must survive restart (via your 43 log or a `meta` file you fsync — document which); committed entries must never roll back.
- Commit & apply: leader commitIndex advancement on majority matchIndex, in-order apply to a state-machine channel, linearizable-submit path with futures.
- Liveness engineering: heartbeat interval vs election timeout ratio (the 10× rule you verify), pre-vote optional stretch, step-down on higher term, single-vote-per-term guarantee.
- Chaos operation: kill leader → re-election ≤ ~2 election timeouts; minority partition → service continues; majority partition → leader steps down, no phantom commits; old-leader rejoin → log repair without data loss.

## Study first (papers before code — no exceptions)

1. Ongaro & Ousterhout, "In Search of an Understandable Consensus Algorithm" — read §1–8 fully (14 pages). Re-draw Figure 2 (state + RPCs + rules) by hand or ASCII in your `RAFT_NOTES.md`. If you can't recite the election-safety argument in 3 sentences, re-read §5.2.
2. Raft lecture/thesis § on log replication + Figure 7 (the consistency-check example) — implement exactly that check, not a weaker one.
3. etcd/raft *concepts* docs OR hashicorp/raft *concepts* (not the code) — snapshotting rationale + leadership transfer idea (stretch only).
4. Your 43 `Replay` + 46 `RPCServer/RPCClient` APIs — these are your transport and disk; Raft logic must not bypass them with ad-hoc sockets/files (except the small `meta` file for term/vote, which you document).
5. Go: `time.Timer` reset patterns (the classic leak), `context`, `sync.Mutex` discipline for `term/state/log` (one mutex, documented lock order — no lock-cycle deadlocks).

Write `RAFT_NOTES.md` (your file, ~60 lines): Figure-2 summary, your timeout choices with justification, your persistence scheme, what you simplified (no joint consensus? no snapshots? single-writer submits?) and what breaks because of it. Verifier doesn't grade it, but your Q6 answer quotes from it.

## Milestones (do in order — do not parallelize M1/M2, the bugs compound)

### M1 — Single-process cluster: election only (Week 1)

- `Node{ID, Peers, State, Term, VotedFor, ...}`, `NewNode(id, peers, opt)`, `Start()/Stop()`, in-process `Network` simulator (message-passing with drop/partition hooks — deterministic, no real sockets yet).
- Randomized election timeout (`ElectionTimeoutMs` base + jitter, default 150–300ms), heartbeat 30ms (document the ratio). One leader at a time; randomized timeouts make split-vote rare; single vote per term enforced (prove with a `DemoElection()` electing exactly 1 leader in 20 runs, printing terms).
- Higher-term step-down: partitioned old leader seeing higher term reverts to follower immediately (unit-demo `DemoStepDown()`).
- Persistence of `term+votedFor` across `Stop→Start` (restarted node never double-votes a term it already voted in). Log persistence can wait for M2, but term/vote cannot.

### M2 — Log replication + commit + apply (Week 2)

- `Entry{Term, Index uint64; Command []byte}`, `Submit(cmd []byte) (Future/chan, error)` (leader only; followers return `ErrNotLeader{LeaderHint}`), AppendEntries with prevLogIndex/prevLogTerm consistency check (reject + return conflict hint; leader decrements nextIndex and retries — the Figure-7 repair).
- Leader `commitIndex` advances on majority `matchIndex`; followers learn commit via `leaderCommit`; `ApplyCh() <-chan Committed{Index, Command}` delivers **in order, exactly once per index** (buffer out-of-order, never skip).
- State-machine demo: 3 nodes applying `kv:set k=v` commands into `map[string]string` (your 44 shapes, in-memory is fine) — `DemoReplicate(500 commands)` proves all three maps identical + commit latency p50/p99 printed.
- Log persisted via 43 (`walseglog` imported, entries encoded by you) so kill-and-restart of a follower repairs via leader retry, not from scratch. Prove: kill follower mid-replication, restart, it catches up with no gaps (`DemoCatchup()`).

### M3 — Real transport + chaos suite (Week 3)

- Swap the simulator for your 46 `RPCServer/RPCClient` (`raft.requestVote`, `raft.appendEntries` methods — the exact methods your 46 `DemoRaftShapedRPC` stubbed). Same Node logic, real TCP on localhost. `DemoTCPCluster()` elects + replicates 200 commands over sockets.
- Chaos matrix (`DemoChaos()`, the shippable proof — all four required):
  1. Kill leader → new leader ≤ ~2 election timeouts, zero committed entries lost, submits during blackout either commit after re-election or return documented error (never silent loss).
  2. Minority partition (1 of 3 isolated) → cluster keeps committing; healed node catches up.
  3. Majority partition (leader + 0 followers connected) → old leader steps down on seeing higher term / stops committing; no phantom commit on minority side.
  4. Old-leader rejoin with stale conflicting entries → leader overwrites only uncommitted suffix (committed prefix untouched — assert with per-index terms log).
- Print a chaos table (scenario × outcome × time-to-recover) into QUESTIONS. Flaky chaos (passes 1 in 3 runs) = FAIL — fix your timers, then paste 3 consecutive passes.

### M4 — Operate it: KV service + README (Week 3–4)

- `ServeRaftKV(...)`: replicated KV (Put/Get/Delete routed to leader, reads via leader by default + documented `ReadRelaxed()` follower path with staleness warning). Show `curl`-style transcript (through 45 HTTP or 46 RESP — your choice) of write → kill leader → write still works.
- `DemoCLI(args) int`: `cluster --nodes 3`, `submit k v`, `status` (leader/term/commit per node), `kill <id>`, `partition <id>`, `heal`, `bench --n`. Same CLI shape as 41/44/45/46.
- `README.md` (your file): Raft subset table (what's in: election/replication/commit/apply/persist; what's out: snapshots/membership change/joint consensus — and what breaks without them), timeout tuning guide, chaos table from your machine, 3-run stability note, next steps (snapshot + `InstallSnapshot`? pre-vote?).
- Performance note: `DemoBenchRaft(n=2000)` — submits/s + commit p50/p99 on localhost TCP. Compare one sentence vs single-node 44 puts/s (replication costs — quantify yours).

## Reuse (required)

- MUST import `43_wal_seglog` (log persistence) and `46_resprpc` (RequestVote/AppendEntries transport), MUST use `18_logger` (term/state transitions logged), `21_config` (node options), `24_scheduler` or timers for heartbeats, `25_taskpool` (apply loop / RPC handling).
- SHOULD reuse `44_kv_store` shapes for the demo state machine, `19_event` (leader-change events), `41_cli` patterns.
- Document in Q1. Re-implementing a private log or private sockets instead of 43/46 without written justification = PARTIAL.

## Banned imports (in impl)

`etcd/raft`, `hashicorp/raft`, any consensus lib, `net/http` for Raft RPC (use 46), `encoding/gob` for log entries. `net`, `sync`, `time`, `context`, `encoding/binary`, `math/rand` (jitter) fine.

Importing a Raft library for the core = FAIL (using one *to compare traces* in a demo, isolated + cited, is allowed and even encouraged).

## How to run

```go
// root main.go (TEMPORARY, verifier ignores it)
package main

import (
	"fmt"
	"practice_go/phase_06_systems/47_raft_lite"
)

func main() { fmt.Println(raftlite.Demo()) }
```

```powershell
go run .
go vet ./phase_06_systems/47_raft_lite/...
```

## Done when

- 20/20 elections yield exactly one leader; step-down, replication-identical-maps, catch-up-after-kill all print PASS.
- TCP cluster replicates; 4-scenario chaos table passes 3 consecutive runs with timings.
- Replicated-KV transcript shows write → kill leader → write still works.
- `go vet` clean; `-race` chaos run clean (state command in Q7); `RAFT_NOTES.md` + `README.md` exist with the subset/tuning/chaos content above.
