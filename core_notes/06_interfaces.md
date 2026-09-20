# Core Note — Interfaces

> Prerequisite for: `phase_01_core/06_interfaces`. Used by: `logger` (Writer), `iox` (Reader), `event` payloads, `taskpool` tasks, `34_io` chain.

## What an interface is
- A SET OF METHODS: `type Shape interface { Area() float64 }`. Any type with those methods satisfies it **implicitly** — no `implements` keyword.
- Small interfaces compose: `io.ReadWriter` = `Reader` + `Writer`. Prefer MANY SMALL interfaces over one fat one.
- `any` (= `interface{}`) accepts everything; recover concrete type with **type assertion** (`v.(int)`) or **type switch**.

## The three stdlib interfaces to know cold
- `error { Error() string }`, `fmt.Stringer { String() string }`, `io.Writer { Write([]byte) (int, error) }` (+ `io.Reader`).
- Accept interfaces, return structs: `func WriteTo(w io.Writer, s string)` works with files, buffers, network conns, YOUR logger — that's the power.

## Traps
- **Nil interface vs typed-nil:** interface holding `(*Circle)(nil)` is NON-nil (has type info). Check with type assertion or design constructors to return nil interfaces explicitly.
- Keep interface values comparable-safe; document whether impls are pointer or value receivers (method sets!).

## Building-block role
- `Logger{out io.Writer}` (swap stdout/buffer/file), `Parse(r io.Reader)` (string/file/network all work), `Task func() (any,error)` executed by pool workers — all interface-shaped seams.
