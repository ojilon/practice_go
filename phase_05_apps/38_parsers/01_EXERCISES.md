# parsers — CSV + Logs — Exercises

> Work in THIS folder, `package parsers`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 csv.go** — ParseCSV with quoted commas, ErrLine struct{Line int; Msg string}.
- **E2 logs.go** — ParseLogLine + Summarize([]Log) report.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package parsers`.

## Build-chain (reuse + banned imports)
- MUST read via your `34_iox` Reader abstraction and assemble strings with your `26_stringsx` Builder logic (import or mirror + cite).
- Banned in impl: `encoding/csv`, `encoding/json` for the parsing itself (demo comparison allowed in a separate func).
- Real-world output: summary report (status counts, top IP) using your `27_sortsearch` ordering.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
