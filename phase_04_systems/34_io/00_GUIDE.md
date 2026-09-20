# io — Reader/Writer Rebuilds — Guide (`iox`)

> Folder: `phase_04_systems/34_io` · Package: `iox` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- DOCS QUEST: go doc io.Reader/io.Writer first.
- Step 1 — type Reader interface{Read([]byte)(int,error)}; Writer; Closer.
- Step 2 — ReadFull/Copy/LimitReader/TeeReader mini versions.
- Step 3 — chaining: LimitReader(TeeReader(...)) demo.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package iox` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_04_systems/34_io"` (package name `iox`), call your funcs, `go run .`.
3. `go vet ./phase_04_systems/34_io/...` must be clean before moving on.

### Details
- DOCS QUEST: go doc io.Reader/io.Writer first.
- Step 1 — type Reader interface{Read([]byte)(int,error)}; Writer; Closer.
- Step 2 — ReadFull/Copy/LimitReader/TeeReader mini versions.
- Step 3 — chaining: LimitReader(TeeReader(...)) demo.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_04_systems/34_io")
func main() { fmt.Println(iox.Demo()) }
```
```powershell
go run .
go vet ./phase_04_systems/34_io/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
