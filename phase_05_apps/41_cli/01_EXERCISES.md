# CLI — Subcommands (flag only) — Exercises

> Work in THIS folder, `package clix`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 cli.go** — Parse(args []string)(cmd string, err) + Run(args) int exit code.
- **E2 help text + 3 subcommands working.** — E2 help text + 3 subcommands working.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package clix`.

## Build-chain (reuse + banned imports)
- Subcommands MUST delegate to your packages: `kv` → your `28_hashmap` or map logic + `35_mini_os` file persist; `serve` → your `37_http_api`; `parse` → your `38_parsers`; all logging via `18_logger`.
- Banned: `cobra`/`urfave/cli` — `flag` + `os.Args` only (you're learning the seam frameworks hide).

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
