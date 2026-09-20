# ringbuffer — Fixed Ring — Exercises

> Work in THIS folder, `package ringbuffer`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 ring.go** — generic Ring[T any] with Write/Read/Peek/Len/Cap/IsFull/IsEmpty + ErrFull/ErrEmpty.
- **E2 overwrite.go** — WriteOverwrite(v T) always succeeds.

## Edge cases (will be spot-checked)
- Empty/nil inputs return documented errors (never panic on expected failures).
- No `fmt.Print*` inside library code (demos/Example funcs may print; logger writes to injected `io.Writer`).
- Exported funcs have `//` doc comments (1 line min).

## Suggested files
- Put each E in its own file (`e01.go`, `e02.go`, ...) or group logically (`stack.go`, `balanced.go`). Any layout OK as long as symbols exist.
- Package clause MUST be `package ringbuffer`.

## Stretch (optional, credited)
- Table-driven demo func `Demo()` printing outputs for manual `go run .` check.
- Benchmarks with `time.Now` (no `testing` required).
