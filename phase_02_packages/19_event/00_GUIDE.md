# event — EventBus — Guide (`event`)

> Folder: `phase_02_packages/19_event` · Package: `event` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type Bus struct{mu; subs map[string][]chan Event}.
- Step 2 — Subscribe(topic) <-chan Event, Publish(topic,payload), Unsubscribe.
- Step 3 — concurrency-safe with RWMutex; non-blocking publish (drop or goroutine choice documented).
- Step 4 — generic payload any; typed-wrapper stretch.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package event` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_02_packages/19_event"` (package name `event`), call your funcs, `go run .`.
3. `go vet ./phase_02_packages/19_event/...` must be clean before moving on.

### Details
- Step 1 — type Bus struct{mu; subs map[string][]chan Event}.
- Step 2 — Subscribe(topic) <-chan Event, Publish(topic,payload), Unsubscribe.
- Step 3 — concurrency-safe with RWMutex; non-blocking publish (drop or goroutine choice documented).
- Step 4 — generic payload any; typed-wrapper stretch.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_02_packages/19_event")
func main() { fmt.Println(event.Demo()) }
```
```powershell
go run .
go vet ./phase_02_packages/19_event/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
