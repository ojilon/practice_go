# 44 — Durable KV Store (Bitcask + LSM-lite) — Guide (`kvstore`)

> Folder: `phase_06_systems/44_kv_store` · Package: `kvstore` · Run via root `main.go` (verifier ignores `main.go`).
> Time budget: **3–4 weeks** (~30–45 hours). This is your first real storage engine. It must survive kills, compact itself, and serve Benchmarks you can brag about.

## Why it exists

SQLite's architecture docs and the Bitcask paper teach the same lesson: a durable KV is a **log + index + compaction**, not magic. You already built the log (43). Now you turn it into a database people can actually use: `Put/Get/Delete`, restart recovery, background compaction, optional TTL + LRU hot cache, and batch writes. Phase 45 will serve this over raw TCP HTTP; Phase 46 over RESP; Phase 47 will replicate its operation log with Raft. If this store is flimsy, everything downstream wobbles.

Real-life payoff: a Go CLI + library you can point at ("my Bitcask clone, 40k puts/s, kill -9 safe, here are the bench numbers"), plus the lived experience of compaction debt, write amplification, and why Deletes are just tombstones.

## What you will learn

- Bitcask model: append values to 43's log, keep `map[key] → (offset, size)` in RAM, rebuild via `Replay` on open.
- LSM-lite extension: sealed segments become immutable sorted tables (simple sorted-string-table file you design) + background merge; or, simpler accepted path: single-active-log + periodic compaction rewrite (document which you chose).
- Correctness under concurrency + crash: single-writer + RWMutex readers, atomic `Batch`, fsync discipline inherited from 43.
- Operations: `Stats` (keys, live bytes, dead bytes, segments), `Compact()` (blocking + trigger threshold), TTL expiry sweeper via your `24_scheduler`, hot-read cache via your `20_cache`.
- Measurable performance: `DemoBench()` with puts/s, gets/s, p50/p99-ish (even crude `time.Since` histogram is fine), before/after compaction sizes.

## Study first

1. Bitcask paper (full, ~6 pages) + Riak's "hint files" idea.
2. SQLite architecture overview (sqlite.org/arch.html) — § on pager vs log-structured; steal the *mental model of recovery*, not the pager.
3. "LSM-tree" intro (any of: LevelDB docs, RocksDB wiki "LSM overview", or the original O'Neil paper abstract+§3) — just enough to decide Bitcask-vs-SSTable for your compaction story.
4. Your own 43 `DESIGN_NOTES.md` — the log API you promised (`Replay`, `TruncateBefore`) is now your client contract.
5. Go: `sync.RWMutex`, `os` file locking (how to refuse double-`Open` on the same dir — even a lockfile is enough), `time` for TTL.

## Milestones (do in order)

### M1 — Bitcask core on top of 43 (Week 1)

- `Open(dir string, opt StoreOptions) (*Store, error)` replays the log into an in-memory index; `Put`, `Get` (`ErrNotFound`), `Delete` (tombstone, not physical removal), `Close`.
- Double-open protection: second `Open` on same dir while first is alive → `ErrLocked` (lockfile or in-process registry — document it).
- Empty key → error; value up to `MaxValueBytes` (default 1 MiB); keys kept as strings, values as opaque bytes (so 45/46 can store JSON blobs from your `23_jsonlite`).
- Reopen test: put 5k keys, delete 1k, close, reopen, all 4k readable, deleted ones return `ErrNotFound`. This is the test the verifier spot-checks — make a `DemoReopen()` that prints the counts.

### M2 — Batches, TTL, cache, config (Week 2)

- `Batch(puts map[string][]byte, deletes []string) error` — atomic from the reader's view (all-or-nothing under one lock + one log sync boundary). Partial failure leaves no visible half-batch.
- TTL: `PutWithTTL(key, val, ttlMs)` + background sweeper every `SweepIntervalMs` using your `24_scheduler` (or a plain ticker you justify). Expired keys behave as `ErrNotFound` and are reclaimed by compaction. `GetWithMeta` returns `expiresAt` for at least one demo path.
- Hot cache: front reads with your `20_cache` LRU (capacity in options); `Stats().CacheHits/CacheMisses` proves it works. Document reuse.
- `StoreOptions` loads from your `21_config`/`22_ini` (dir, segment max, sync policy, cache size, compaction threshold). Provide `DefaultOptions()` + `LoadOptions(path)` — config round-trip demo required.

### M3 — Compaction that actually reclaims (Week 2–3)

- `Compact() (CompactionStats, error)` rewrites live keys into a fresh segment set (or sorted table if you chose SSTable-lite), swaps atomically, then `TruncateBefore`s dead segments. Crash *during* compaction must not lose committed data (write new files + fsync + rename-swap + only then delete old — document the order).
- Auto-trigger: when `deadBytes > DeadBytesThreshold` or `deadRatio > 0.3`, background compaction via your `25_taskpool` (one background worker, never blocks `Put` more than one batch window).
- Prove it: `DemoCompact()` — load 20k keys, overwrite/delete 70%, print `Stats` before/after (bytes + segments), reopen, verify survivors. Paste numbers into 02_QUESTIONS. A compaction that doesn't shrink anything = FAIL.
- Tombstone GC rule: a tombstone may only be dropped when compaction is sure no older segment still needs it — state your rule in one comment.

### M4 — CLI, bench, README — the shippable (Week 3–4)

- `DemoCLI(args []string) int`: `open --dir`, `put k v`, `get k`, `del k`, `batch`, `stats`, `compact`, `bench --n 50000`. Same subcommand shape as 41. Must run against a real dir on disk (use `t.TempDir()`-style temp dir in demos, real path in CLI).
- `DemoBench(n int) string`: N puts + N gets timed, prints puts/s, gets/s, crude p50/p99 from sampled latencies, plus `Stats`. Run with `--n 20000` minimum for your QUESTIONS numbers.
- `README.md` (your file): architecture diagram (ASCII ok: client → cache → index → 43 log / SSTables), durability contract ("what survives kill -9 and what doesn't"), compaction policy, bench table from your machine, what you'd do next (compression? bloom filters? — pick two and say why).
- Seam for 45/46/47: they import `practice_go/phase_06_systems/44_kv_store` and call `Open/Put/Get/Delete/Batch/Stats/Compact`. Add one `DemoServeStyle()` that does 500 mixed ops through only that public API (no internals) to prove the seam is clean.

## Reuse (required)

- MUST import `practice_go/phase_06_systems/43_wal_seglog` for the log (no second log implementation). MUST use `18_logger`, `21_config`/`22_ini`, `20_cache` (hot path), `24_scheduler` (TTL sweeper), `25_taskpool` (background compaction).
- SHOULD use `28_hashmap` ideas for the index (Go map is fine if you cite why), `23_jsonlite` for CLI value pretty-printing, `41_cli` command patterns.
- Document all reuse in Q1.

## Banned imports (in impl)

`encoding/gob` for storage path, any embedded DB (`bbolt`, `pebble`, `sqlite`), any Redis/Kafka client, `net/http` server code here (that belongs to 45 — keep this package transport-agnostic). `os`, `sync`, `time`, `hash/crc32` (via 43), `sort` (for SSTable-lite path) are fine.

## How to run

```go
// root main.go (TEMPORARY, verifier ignores it)
package main

import (
	"fmt"
	"practice_go/phase_06_systems/44_kv_store"
)

func main() { fmt.Println(kvstore.Demo()) }
```

```powershell
go run .
go vet ./phase_06_systems/44_kv_store/...
```

## Done when

- All `01_EXERCISES.md` symbols exist with edge-case behaviour (reopen survival, atomic batch, TTL expiry, compaction shrink).
- `02_QUESTIONS.md` filled with real reopen counts, compaction before/after bytes, and bench puts/s + gets/s from your machine.
- `go vet` clean; concurrent mixed workload (8 writers + 8 readers × 1k ops) loses no committed writes and shows no race in your own `-race` run (state command in Q6).
- 45/46/47 can drive the store through the public API only (`DemoServeStyle` proves it).
