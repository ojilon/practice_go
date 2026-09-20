# Goroutines — Guide (`goroutinesx`)

> Folder: `phase_01_core/09_goroutines` · Package: `goroutinesx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — go f(), sync.WaitGroup Add/Done/Wait pattern.
- Step 2 — closure capture bug: go func(i int) with loop var (Go 1.22+ semantics note).
- Step 3 — sync.Mutex for shared counter; go run -race.
- Step 4 — GOMAXPROCS intuition; time deterministic vs nondeterministic output.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package goroutinesx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/09_goroutines"` (package name `goroutinesx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/09_goroutines/...` must be clean before moving on.

### Details
- Step 1 — go f(), sync.WaitGroup Add/Done/Wait pattern.
- Step 2 — closure capture bug: go func(i int) with loop var (Go 1.22+ semantics note).
- Step 3 — sync.Mutex for shared counter; go run -race.
- Step 4 — GOMAXPROCS intuition; time deterministic vs nondeterministic output.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/09_goroutines")
func main() { fmt.Println(goroutinesx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/09_goroutines/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
