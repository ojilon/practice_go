# walseglog — Segmented WAL — Exercises

> Work in THIS folder, `package walseglog`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)

- **E1 types.go** — `Record{Offset uint64; Key, Value []byte}`, `Options{Dir string; SegmentMaxBytes int64; SyncOnWrite bool; MaxKeyBytes, MaxValueBytes int; Logger any-or-your-logger}`, `Stats{Segments, Records int; Bytes int64; Lowest, Highest uint64; TruncatedTails int}`, sentinel errors `ErrClosed`, `ErrCorrupt`, `ErrBadOffset`, `ErrBadHeader`, `ErrTooLarge`.
- **E2 log.go** — `func Open(dir string, opt Options) (*Log, error)`, methods `(l *Log) Append(key, value []byte) (uint64, error)`, `(l *Log) Read(offset uint64) (Record, error)`, `(l *Log) Sync() error`, `(l *Log) Close() error`, `(l *Log) LowestOffset() uint64`, `(l *Log) HighestOffset() uint64`, `(l *Log) Stats() Stats`.
- **E3 segments.go** — rotation on `SegmentMaxBytes`, `func (l *Log) TruncateBefore(offset uint64) error`, `func (l *Log) Replay(from uint64, fn func(Record) error) error` (streaming, bounded RAM).
- **E4 consumer.go** — ONE of: `Subscribe(from uint64) (<-chan Record, cancel func(), err error)` OR `WaitFor(offset uint64, timeoutMs int) (Record, error)`. Document choice in a doc comment.
- **E5 demo.go** — `func Demo() string` (happy-path summary), `func DemoReuse44Style() string` (write→reopen→Replay into map, print count), `func DemoBenchSync() string` (sync-on-write vs batched numbers), `func DemoCLI(args []string) int` (append/read/stats/truncate/bench).

## Edge cases (will be spot-checked)

- Empty key returns documented error; nil/empty value is legal (tombstone convention for 44). Oversize key/value → `ErrTooLarge`, never partial append.
- `Append` after `Close` → `ErrClosed`. `Read` past highest → `ErrBadOffset`. `TruncateBefore` on active segment / beyond highest → `ErrBadOffset`.
- Torn tail on `Open` truncates and counts in `Stats().TruncatedTails`; mid-file corruption → `Open` or `Read` returns `ErrCorrupt`.
- No `fmt.Print*` inside library code (demos may print; logger writes to injected `io.Writer`).
- Exported funcs/types have `//` doc comments (1 line min).
- Concurrency: `Append` from 8 goroutines × 500 appends each loses nothing and `go vet` is clean.

## Suggested files

- `types.go`, `log.go`, `segments.go`, `consumer.go`, `demo.go` (any layout OK as long as symbols exist).
- Package clause MUST be `package walseglog`.

## Build-chain (reuse + banned imports)

- MUST import or mirror-cite your `18_logger` and your `21_config`/`22_ini` (Options defaults). MUST cite your `34_io` exact-read pattern.
- Banned in impl: `encoding/gob`, `encoding/json` (record path), `go.etcd.io/bbolt`, `cockroachdb/pebble`, `mattn/go-sqlite3`, any `kafka` client. Demo-only comparison use is allowed if isolated in `demo.go`.
- Downstream seam: 44/47 import `practice_go/phase_06_systems/43_wal_seglog`. Keep `Open/Append/Read/Replay/TruncateBefore/Close` signatures stable.

## Stretch (optional, credited)

- Sparse in-memory segment index (`baseOffset → file`) + binary search for `Read`.
- `VerifyAll() error` full-CRC scrub command.
- Fsync-batch tuning bench table (batch 1/10/100/1000) in README.
