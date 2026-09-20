# list — Linked List (generics) — Guide (`list`)

> Folder: `phase_02_packages/13_list` · Package: `list` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Node[T any]{Val T; Next *Node[T]}; type List[T any]{head,tail,len}.
- Step 2 — PushFront/PushBack/PopFront/Len/ToSlice.
- Step 3 — Remove(v T) bool, Reverse(), doubly-linked stretch.
- Step 4 — why generics: same code for int/string.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package list` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/13_list"` (package name `list`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/13_list/...` must be clean before moving on.

### Details
- Step 1 — type Node[T any]{Val T; Next *Node[T]}; type List[T any]{head,tail,len}.
- Step 2 — PushFront/PushBack/PopFront/Len/ToSlice.
- Step 3 — Remove(v T) bool, Reverse(), doubly-linked stretch.
- Step 4 — why generics: same code for int/string.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/13_list")
func main() { fmt.Println(list.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/13_list/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
