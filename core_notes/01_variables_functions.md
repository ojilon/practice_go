# Core Note — Variables & Functions

> Prerequisite for: `phase_01_core/00_myfmt`, `01_variables_functions`. Used by: everything.

## Variables
- `var x int` declares with **zero value** (`0`, `""`, `false`, `nil` for slices/maps/chans/pointers/interfaces). `:=` infers + declares (function scope only).
- `const` for compile-time values; untyped constants (`const N = 10`) adapt to context.
- **Shadowing trap:** `:=` inside a narrower scope creates a NEW variable; outer keeps old value. Classic bug: `err` shadowed inside `if`, checked outside.
- Blank identifier `_` discards values you must explicitly ignore.

## Functions
- Signature carries meaning: `func Div(a, b float64) (float64, error)` — Go convention returns result + error, caller checks `if err != nil`.
- **Named returns** (`func Split(s string) (first, last string)`) double as documentation; combined with `defer` they can be mutated before return.
- **Variadic:** `func Sum(nums ...int)`; call with slice via `Sum(s...)`.
- **First-class:** pass `func(int) int` as args, return closures that capture state (`Counter()` exercise). Loop-var capture: pre-1.22 `for i := ...; go func() { use i }` shared one variable — pass as param to be safe (know both semantics).

## Building-block role
- Multiple returns + `error` are the backbone of `07_errors`, `18_logger` signatures, `23_jsonlite` (`Marshal(v any) (string, error)`), every parser.
- Closures power `24_scheduler` (`Every(d, fn func())`) and `25_taskpool` (`Submit(task func() (any, error))`).
