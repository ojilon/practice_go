# math2 — Vectors & Matrices — Guide (`math2`)

> Folder: `phase_02_packages/12_math2` · Package: `math2` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Vec2{X,Y float64} + Add/Sub/Scale.
- Step 2 — Dot, Len/Norm (normalize with zero-vector error), Cross (2D scalar).
- Step 3 — Vec3 + Cross3, type Mat2 [2][2]float64 + MulVec/MulMat.
- Step 4 — float tolerance: approxEq with 1e-9.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package math2` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/12_math2"` (package name `math2`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/12_math2/...` must be clean before moving on.

### Details
- Step 1 — type Vec2{X,Y float64} + Add/Sub/Scale.
- Step 2 — Dot, Len/Norm (normalize with zero-vector error), Cross (2D scalar).
- Step 3 — Vec3 + Cross3, type Mat2 [2][2]float64 + MulVec/MulMat.
- Step 4 — float tolerance: approxEq with 1e-9.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/12_math2")
func main() { fmt.Println(math2.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/12_math2/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
