# paperreimpl — Paper Capstone — Exercises

> Work in THIS folder, `package paperreimpl`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)

- **E1 types.go** — `func Track() string` (`"dynamo"|"mapreduce"|"gfs"`), `type EvalRow{Config string; OpsPerSec float64; Notes string}`, `type FailureResult{Scenario, Outcome string; RecoveredMs int64; DataIntact bool}`, errors `ErrConflict` (dynamo siblings), `ErrNoWorker`/`ErrJobFailed` (mapreduce), `ErrNoLease`/`ErrCorruptChunk` (gfs), `ErrClosed`, `ErrBadInput`.
- **E2 track core** — ONE track fully (others may be absent, but the chosen track must be complete):
  - Dynamo: `type VectorClock map[string]uint64`, `func (v VectorClock) Compare(o VectorClock) int` (-1/0/1/2=concurrent), `func Merge(a, b VectorClock) VectorClock`, `type Dynamo struct...` + `func NewDynamo(...)` + `(d *Dynamo) Put(key string, val []byte) error` + `(d *Dynamo) Get(key string) ([][]byte, VectorClock, error)` (siblings on conflict) + membership/ring helpers + handoff/anti-entropy entry points.
  - MapReduce: `type MapFunc func(line string) [][2]string`, `type ReduceFunc func(key string, vals []string) string`, `type JobRunner struct...` + `func NewJobRunner(...)` + `(j *JobRunner) Run(inputs []string, mapFn MapFunc, redFn ReduceFunc) (map[string]string, error)` + split/shuffle status helpers.
  - GFS: `type GFSClient/GFSMaster/ChunkServer` minimal trio + `func NewTestCluster(...)` + `(c *GFSClient) Create(path string) error` + `(c *GFSClient) Write(path string, data []byte) error` + `(c *GFSClient) Read(path string) ([]byte, error)` + `(c *GFSClient) Delete(path string) error` + chunk-corrupt/re-replicate helper.
- **E3 eval.go** — `func DemoEval() string` (perf table print), `func DemoFailure() string` (failure matrix print with DataIntact asserts), `func ReuseMap() string` (≥6 own packages + one-line where-used each).
- **E4 demo_cli.go** — `func Demo() string` (5-minute live path), `func DemoCLI(args []string) int` (serve/bench/fail subcommands for the chosen track), track correctness demo: `DemoConflict()` XOR `DemoFaultTolerance()` XOR `DemoFailover()` (matching Track()).

## Edge cases (will be spot-checked)

- Dynamo: concurrent writers → siblings (not last-writer-wins silent drop); R+W>N vs R+W≤N behaviour difference stated + measured; down-node write + heal → handoff drains, later reads converge.
- MapReduce: worker-kill mid-job → byte-identical output + quantified slowdown; empty input → empty output (not error); same input twice → byte-identical outputs.
- GFS: corrupt chunk → `ErrCorruptChunk` on direct read, re-replicated after scrub; master restart → namespace intact; primary-kill mid-write → no half-chunk visible to readers.
- Common: empty key/path → `ErrBadInput`; ops after `Close` → `ErrClosed`; no `fmt.Print*` in library code; exported symbols have `//` doc comments.

## Suggested files

- `types.go`, `dynamo.go` XOR `mapreduce.go` XOR `gfs.go`, `eval.go`, `demo_cli.go` (+ required `PAPER_NOTES.md`, `README.md`, `RETRO.md` you author).
- Package clause MUST be `package paperreimpl`.

## Build-chain (reuse + banned imports)

- MUST document ≥6 reuses in `ReuseMap()` + Q1 (43 + 44 + 46 + logger + config + taskpool is the expected spine; scheduler/hashmap/jsonlite/cli fill the rest).
- Banned in impl: riak/hadoop/hdfs libs, etcd/raft + hashicorp/raft, `net/http` internal RPC, `encoding/gob` (wire/store). Edge serving via 45 + isolated comparison cites allowed.
- <6 reuses → PARTIAL. Ready-made paper system import → FAIL.

## Stretch (optional, credited)

- Second track's thin slice (M1-level path) for comparison notes.
- Dynamo Merkle-tree anti-entropy; MR straggler mitigation with measured win; GFS snapshot + lease-recovery timing.
- 5-minute screen-record script (commands + expected outputs) in README.
