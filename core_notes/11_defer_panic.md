# Core Note — Defer, Panic/Recover, Control Flow

> Prerequisite for: `phase_01_core/11_defer_panic_modules`. Used by: resource cleanup everywhere (files, mutexes, servers, tickers).

## Defer
- `defer f()` runs at function exit, LIFO order. Canonical uses: `defer file.Close()`, `defer mu.Unlock()`, `defer ticker.Stop()`.
- Args evaluate IMMEDIATELY: `defer fmt.Println(i)` prints i-at-defer-time. Loop gotcha: `defer` in a loop stacks until function end — usually refactor into helper or collect closers.
- With named returns, deferred funcs can modify results (`defer func(){ n += 3 }()` runs before actual return).

## Panic / recover
- `panic(v)` unwinds the stack running defers. `recover()` in a deferred func stops it and yields the value — use ONLY at strategic boundaries (worker wrapper, `SafeRun`, HTTP middleware).
- Rule restated: user errors → `error` return; panics → converted to errors at the boundary, then handled as errors.

## Control flow compact
- `if x, err := f(); err != nil { ... }` — init statement scoped to branches.
- `switch` (no fallthrough by default; `fallthrough` keyword opts in), `switch v := x.(type)` = type switch.
- `for` is the only loop (3 forms + `range`); labeled `break`/`continue` for nested loops (used in tokenizer/parser inner loops).

## Building-block role
- Every `Open/Close`, `Lock/Unlock`, `Start/Stop`, `Subscribe/Unsubscribe` pair downstream is a defer site. Missing defers = leaked files/goroutines/tickers — verifier spot-checks them.
