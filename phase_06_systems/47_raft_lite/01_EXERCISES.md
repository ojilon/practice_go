# raftlite — Raft-lite — Exercises

> Work in THIS folder, `package raftlite`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)

- **E1 types.go** — `type State int` (`Follower, Candidate, Leader`), `type Entry{Term, Index uint64; Command []byte}`, `type Committed{Index uint64; Command []byte}`, `type NodeOptions{ID string; Peers []string; ElectionTimeoutMs, HeartbeatMs int; Dir string}`, `func DefaultNodeOptions(id string, peers []string) NodeOptions`, errors `ErrNotLeader`, `ErrClosed`, `ErrNoLeader`, `type ErrNotLeader struct{ LeaderHint string }` with `Error() string`.
- **E2 node.go** — `type Node struct...` with `func NewNode(opt NodeOptions) (*Node, error)`, `(n *Node) Start() error`, `(n *Node) Stop() error`, `(n *Node) State() State`, `(n *Node) Term() uint64`, `(n *Node) Leader() string`, `(n *Node) CommitIndex() uint64`, `(n *Node) Submit(cmd []byte) (<-chan Committed, error)`, `(n *Node) ApplyCh() <-chan Committed`.
- **E3 rpc_transport.go** — RequestVote/AppendEntries over 46 (`raft.requestVote`, `raft.appendEntries` registered on `resprpc.RPCServer`), `type VoteArgs/Reply`, `AppendArgs/Reply` structs + encode/decode helpers, conflict-hint repair (nextIndex decrement on rejection).
- **E4 cluster.go** — in-process `Network` simulator (drop/partition/kill hooks) + TCP cluster helper `func StartTCPCluster(n int, baseDir string) ([]*Node, func(), error)`, chaos helpers `(Kill(id)/Partition(id)/HealAll)` used by demos.
- **E5 demo_cli.go** — `func Demo() string`, `func DemoElection() string` (20 runs, exactly-1-leader each), `func DemoStepDown() string`, `func DemoReplicate(n int) string` (identical maps + p50/p99), `func DemoCatchup() string`, `func DemoTCPCluster() string`, `func DemoChaos() string` (4-scenario table), `func DemoBenchRaft(n int) string`, `func DemoCLI(args []string) int`.

## Edge cases (will be spot-checked)

- Two candidates, one term → exactly one winner; double-vote in same term impossible (restarted node honours votedFor).
- Higher-term RPC forces step-down before any log mutation. Follower rejects AppendEntries with mismatched prevLogTerm (no silent overwrite of committed prefix).
- `Submit` on follower → `ErrNotLeader` with hint; `Submit` with no leader → `ErrNoLeader`; `Submit` after `Close` → `ErrClosed`.
- Apply delivers in index order, exactly once per index, across kill/restart (no gaps, no dups — assert in DemoCatchup).
- Committed entries never roll back when old leader rejoins (assert committed-prefix terms unchanged in chaos #4).
- No `fmt.Print*` in library code; exported symbols have `//` doc comments.

## Suggested files

- `types.go`, `node.go`, `rpc_transport.go`, `cluster.go`, `demo_cli.go` (+ your own `RAFT_NOTES.md`, `README.md` — required docs, not checked for symbols).
- Package clause MUST be `package raftlite`.

## Build-chain (reuse + banned imports)

- MUST import `43_wal_seglog` + `46_resprpc`, use `18_logger`, `21_config`, `24_scheduler`/timers, `25_taskpool`. SHOULD cite `44_kv_store`, `19_event`, `41_cli`.
- Banned in impl: `etcd/raft`, `hashicorp/raft`, any consensus lib, `net/http` for Raft RPC, `encoding/gob` (log entries). Comparison use isolated in demos + cited is allowed.
- Private log/socket re-implementation instead of 43/46 without justification → PARTIAL.

## Stretch (optional, credited)

- Pre-vote phase; leadership transfer (`TransferLeadership`); `InstallSnapshot`-lite compaction via 44 snapshot.
- Linearizable read (`ReadIndex`-lite) vs relaxed read demo with staleness numbers.
- Jepsen-lite history checker (record submit/commit/apply trace, verify linearizability offline).
