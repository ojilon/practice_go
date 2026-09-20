# queue — FIFO + Deque — Guide (`queue`)

> Folder: `phase_02_packages/15_queue` · Package: `queue` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — slice+head index queue (avoid O(n) shifts) OR list-backed.
- Step 2 — Enqueue/Dequeue(error)/Peek/Len/IsEmpty.
- Step 3 — Deque: PushFront/PushBack/PopFront/PopBack.
- Step 4 — channel-backed blocking queue stretch.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package queue` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/15_queue"` (package name `queue`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/15_queue/...` must be clean before moving on.

### Details
- Step 1 — slice+head index queue (avoid O(n) shifts) OR list-backed.
- Step 2 — Enqueue/Dequeue(error)/Peek/Len/IsEmpty.
- Step 3 — Deque: PushFront/PushBack/PopFront/PopBack.
- Step 4 — channel-backed blocking queue stretch.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/15_queue")
func main() { fmt.Println(queue.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/15_queue/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
