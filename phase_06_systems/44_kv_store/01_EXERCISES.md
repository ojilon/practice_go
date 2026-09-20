# kvstore — Durable KV — Exercises

> Work in THIS folder, `package kvstore`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)

- **E1 types.go** — `StoreOptions{Dir string; SegmentMaxBytes int64; SyncOnWrite bool; CacheSize int; DeadBytesThreshold int64; SweepIntervalMs int; MaxValueBytes int}`, `func DefaultOptions(dir string) StoreOptions`, `func LoadOptions(path string) (StoreOptions, error)`, `Stats{Keys int; LiveBytes, DeadBytes int64; Segments int; CacheHits, CacheMisses uint64}`, `CompactionStats{SegmentsBefore, SegmentsAfter int; BytesBefore, BytesAfter int64}`, errors `ErrNotFound`, `ErrLocked`, `ErrClosed`, `ErrTooLarge`, `ErrBadKey`.
- **E2 store.go** — `func Open(dir string, opt StoreOptions) (*Store, error)`, `(s *Store) Put(key string, val []byte) error`, `(s *Store) Get(key string) ([]byte, error)`, `(s *Store) Delete(key string) error`, `(s *Store) Close() error`, `(s *Store) Stats() Stats`.
- **E3 batch_ttl.go** — `(s *Store) Batch(puts map[string][]byte, deletes []string) error` (atomic), `(s *Store) PutWithTTL(key string, val []byte, ttlMs int64) error`, `(s *Store) GetWithMeta(key string) (val []byte, expiresAtMs int64, err error)`.
- **E4 compact.go** — `(s *Store) Compact() (CompactionStats, error)` (crash-safe swap, tombstone GC rule documented).
- **E5 demo_cli.go** — `func Demo() string`, `func DemoReopen() string`, `func DemoCompact() string`, `func DemoBench(n int) string`, `func DemoServeStyle() string`, `func DemoCLI(args []string) int`.

## Edge cases (will be spot-checked)

- Empty key → `ErrBadKey`; oversize value → `ErrTooLarge`; `Get` missing/expired/deleted → `ErrNotFound` (never nil,nil).
- `Batch` with empty maps is a no-op success; failed batch leaves store unchanged (verify by Get-after-fail in demo).
- Second `Open` on same live dir → `ErrLocked`; `Put` after `Close` → `ErrClosed`.
- TTL expiry is observable: `PutWithTTL(k,v,50)` + sleep 100ms → `Get` returns `ErrNotFound`.
- Compaction preserves survivors across reopen; bytes/segments shrink when dead ratio was high.
- No `fmt.Print*` in library code; exported symbols have `//` doc comments.
- Concurrent 8W×8R×1k mixed ops: no lost committed writes, no race.

## Suggested files

- `types.go`, `store.go`, `batch_ttl.go`, `compact.go`, `demo_cli.go`. Any layout OK as long as symbols exist.
- Package clause MUST be `package kvstore`.

## Build-chain (reuse + banned imports)

- MUST import `practice_go/phase_06_systems/43_wal_seglog` (the log), your `18_logger`, `21_config`/`22_ini`, `20_cache`, `24_scheduler`, `25_taskpool`. Cite `28_hashmap`/`23_jsonlite`/`41_cli` where mirrored.
- Banned in impl: `encoding/gob` (storage path), `bbolt/pebble/sqlite` libs, redis/kafka clients, `net/http` server code.
- Missing reuse import/cite → PARTIAL ("state which own-package you reused").

## Stretch (optional, credited)

- SSTable-lite sealed tables (sorted keys + binary search `Get`) instead of plain log compaction.
- Prefix scan `Scan(prefix string, limit int) ([]KV, error)` + range demo.
- `Verify() error` scrub (re-read all live values through 43 CRCs) + repair report.
