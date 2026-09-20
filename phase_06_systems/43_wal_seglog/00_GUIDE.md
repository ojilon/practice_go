# 43 — Segmented WAL / Unified Log — Guide (`walseglog`)

> Folder: `phase_06_systems/43_wal_seglog` · Package: `walseglog` · Run via root `main.go` (verifier ignores `main.go`).
> Time budget: **2.5–3.5 weeks** (~25–35 hours). This is the foundation for 44 (KV) and 47 (Raft). Do it properly — every later system trusts this log.

## Why it exists

Every serious storage system (SQLite WAL, Bitcask append-log, Kafka segments, etcd Raft log, Postgres WAL) is the same idea: **append-only records with offsets, checksums, segments, and crash recovery**. Jay Kreps' essay "The Log" argues the log *is* the database. This module makes that concrete: you build a small production-flavoured segmented log that survives kills, detects torn writes, rotates segments, and supports consumers and truncation.

Real-life payoff: after this you can explain — from experience, not theory — why databases fsync, why CRCs matter, why Kafka uses segments instead of one giant file, and what "replay on restart" actually costs.

## What you will learn

- On-disk record framing: `len | offset | key-len | key | value | crc32` (your exact layout is your choice, but it must be documented).
- Segment rotation (`seg-000001.log`, index of base offsets), retention/truncation.
- Crash safety: torn-tail detection on open, `ErrCorrupt` semantics, `SyncOnWrite` vs batched `Sync()` trade-off you can *measure*.
- Reader semantics: `Read(offset)`, blocking `Subscribe`/tail, consumer checkpoints.
- Operational reality: throughput bench, recovery-time bench, kill -9 test, disk-full handling.

## Study first (do not copy — extract ideas)

1. Jay Kreps, "The Log: What every software engineer should know about real-time data's unifying abstraction" — read fully, take 10-bullet notes.
2. Bitcask paper (Riak docs, ~6 pages) — focus on append + in-memory index + hint files. You only need §2–3.
3. Kafka log segment design docs (log segments, sparse index, retention) — skim `kafka/log/LogSegment` concepts, not the Scala.
4. SQLite WAL mode overview (sqlite.org/wal.html) — checkpoints vs truncation.
5. Go stdlib: `os`, `bufio`, `hash/crc32`, `encoding/binary` source — how they frame + flush + fsync.

Write your notes as comments in `DESIGN_NOTES.md` (your own file, 20–40 lines): record layout diagram, fsync policy, what happens on torn write, segment naming. Verifier does not check this file, but 44/47 reviewers will read it.

## Milestones (simple → complex — do in order)

### M1 — Single-segment append + read + reopen (Week 1, days 1–5)

- `Open`, `Append(key, value []byte) (uint64, error)`, `Read(offset)`, `Close`.
- Fixed header with magic + version so a garbage file fails fast with `ErrBadHeader`, not a panic.
- CRC32 (Castagnoli) over key+value; `Read` verifies and returns `ErrCorrupt` on mismatch.
- Reopen test: append 1k records, close, reopen, all readable with identical offsets. Offsets start at 0 or resume from existing tail — document which.
- Banned: no `encoding/gob`, no embedded DB (bbolt/pebble/sqlite). Frame bytes yourself with `encoding/binary`.

### M2 — Crash safety + fsync policy you can measure (Week 1–2)

- `Options{Dir, SegmentMaxBytes, SyncOnWrite bool, MaxKeyBytes, MaxValueBytes}`.
- `Sync() error` explicit; when `SyncOnWrite=false`, appends sit in `bufio.Writer` until `Sync` or segment rotation or `Close`.
- Torn-tail handling: on `Open`, scan forward; on short read / CRC fail **at the tail only**, truncate the torn suffix, record `Stats().TruncatedTails`, keep serving. CRC fail in the *middle* → `ErrCorrupt` and refuse to open past it (never silently skip).
- Kill test (real life): write a `cmd`-style `DemoKill()` that appends 50k records without sync, kills the process mid-way (or simulates by closing the file descriptor without `Close`), reopens, and reports how many survived vs truncated. Paste the numbers into 02_QUESTIONS.
- Bench: `Benchmark`-style `DemoBenchSync()` comparing `SyncOnWrite=true` vs batched `Sync` every N appends. Expect ≥5× difference on most disks — record what you actually get.

### M3 — Segments, rotation, truncation, consumers (Week 2–3)

- When active segment exceeds `SegmentMaxBytes`, seal it (rename/freeze, flush+sync) and open `seg-<baseoffset>.log`.
- `LowestOffset()`, `HighestOffset()`, `Stats() Stats{Segments, Records, Bytes, Lowest, Highest}`.
- `TruncateBefore(offset)` deletes whole sealed segments strictly below offset; refuses to truncate the active segment or an offset beyond highest (`ErrBadOffset`); fsyncs directory state where platform allows.
- `Replay(from uint64, fn func(Record) error) error` — the exact hook 44 and 47 will use to rebuild indexes on startup. Must stream (no loading whole log into RAM).
- Consumer helper: `Subscribe(from uint64) (<-chan Record, func(), error)` or a blocking `WaitFor(offset, timeout)` — pick one, document it. Must not busy-spin (use cond/channel + `context`-style timeout param, not `time.Sleep` polling faster than 5ms).
- Concurrency: concurrent `Append` from ≥8 goroutines is safe (mutex or single writer goroutine — your choice); concurrent `Append` + `Read` never races (`go vet` + `-race` clean by your own run).

### M4 — Operations + README (Week 3, shippable)

- Small CLI in `cmd`-style demo func `DemoCLI(args []string) int`: `append <k> <v>`, `read <offset>`, `stats`, `truncate <offset>`, `bench`. Wire it from root `main.go` temporarily — this is the same CLI-shape skill 41 taught.
- `README.md` (your own file, ~40 lines): layout diagram, fsync guidance ("when to turn SyncOnWrite on"), recovery story, benchmark table from your machine, known simplifications vs Kafka (no sparse index, no compression, single writer).
- Leave a seam for 44/47: they must be able to `import "practice_go/phase_06_systems/43_wal_seglog"` and call `Open/Append/Read/Replay/TruncateBefore`. Do not break this API later without updating them.

## Reuse (required — verifier checks)

- MUST use your `18_logger` (inject `*logger.Logger` or `io.Writer`; structured fields for rotation/truncation events) and your `21_config` (or `22_ini`) for `Options` defaults.
- SHOULD mirror your `34_io` helpers (`ReadFull`-style exact reads, `LimitReader` guard on value sizes) — cite with a comment if you re-implement inline.
- Document reuse in 02_QUESTIONS Q1 (import path or `// mirrors …` comment counts).

## Banned imports (in impl files)

`encoding/gob`, `encoding/json` (for the record path — Demo comparison may use it), any embedded store (`go.etcd.io/bbolt`, `cockroachdb/pebble`, `mattn/go-sqlite3`, `kafka` clients). `hash/crc32`, `encoding/binary`, `os`, `bufio`, `sync`, `io` are fine and expected.

Thin wrapper over someone else's WAL = FAIL.

## How to run (example — adapt)

```go
// root main.go (TEMPORARY, verifier ignores it)
package main

import (
	"fmt"
	"practice_go/phase_06_systems/43_wal_seglog"
)

func main() {
	fmt.Println(walseglog.Demo())
}
```

```powershell
go run .
go vet ./phase_06_systems/43_wal_seglog/...
```

## Done when

- All `01_EXERCISES.md` symbols exist with the stated edge-case behaviour.
- `02_QUESTIONS.md` boxes filled, including real kill-test + sync-bench numbers from your machine.
- `go vet ./phase_06_systems/43_wal_seglog/...` clean; `-race` run of your concurrent-append demo shows no data race (state the command you ran in Q5).
- 44 can import this package and rebuild state via `Replay` — prove it with one `DemoReuse44Style()` func that writes 100 records, reopens, replays into a `map[string]string`, and prints the count.
