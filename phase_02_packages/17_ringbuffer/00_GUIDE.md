# ringbuffer — Fixed Ring — Guide (`ringbuffer`)

> Folder: `phase_02_packages/17_ringbuffer` · Package: `ringbuffer` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — buf []T fixed cap, head/tail/count.
- Step 2 — Write(v)(error if full) / Read()(error if empty) + Peek/Len/Cap.
- Step 3 — Overwrite variant WriteOverwrite; policy documented.
- Step 4 — blocking chan-backed stretch.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package ringbuffer` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/17_ringbuffer"` (package name `ringbuffer`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/17_ringbuffer/...` must be clean before moving on.

### Details
- Step 1 — buf []T fixed cap, head/tail/count.
- Step 2 — Write(v)(error if full) / Read()(error if empty) + Peek/Len/Cap.
- Step 3 — Overwrite variant WriteOverwrite; policy documented.
- Step 4 — blocking chan-backed stretch.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/17_ringbuffer")
func main() { fmt.Println(ringbuffer.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/17_ringbuffer/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
