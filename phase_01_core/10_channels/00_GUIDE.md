# Channels — Guide (`channelsx`)

> Folder: `phase_01_core/10_channels` · Package: `channelsx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — unbuffered ch := make(chan int); send/recv; close + range.
- Step 2 — buffered, len/cap, blocking demo.
- Step 3 — select with time.After/default; done channel + context pattern.
- Step 4 — pipelines: gen→sq→filter; fan-in (2 chans→1), fan-out (N workers).

## Steps (simple → complex)
Do in order. Create `*.go` files with `package channelsx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/10_channels"` (package name `channelsx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/10_channels/...` must be clean before moving on.

### Details
- Step 1 — unbuffered ch := make(chan int); send/recv; close + range.
- Step 2 — buffered, len/cap, blocking demo.
- Step 3 — select with time.After/default; done channel + context pattern.
- Step 4 — pipelines: gen→sq→filter; fan-in (2 chans→1), fan-out (N workers).

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/10_channels")
func main() { fmt.Println(channelsx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/10_channels/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
