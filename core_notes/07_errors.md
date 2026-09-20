# Core Note — Errors

> Prerequisite for: `phase_01_core/07_errors`. Used by: every API that can fail (all parsers, queue/stack pops, config getters, HTTP handlers).

## Go error philosophy
- Errors are VALUES, not exceptions. Return them, check them: `if err != nil`.
- Sentinel: `var ErrEmpty = errors.New("empty")` — stable identity, compare with `errors.Is`.
- **Wrap with context at each layer:** `fmt.Errorf("open %s: %w", path, err)` — `%w` preserves the chain; `%v` breaks it.
- Inspect: `errors.Is(err, ErrEmpty)` (identity through wraps), `errors.As(err, &vErr)` (extract typed error).
- Custom type: `type ValidationError struct{ Field, Msg string }` + `Error() string` → carries structured data, still an `error`.

## Panic rule
- `panic` = programmer bug / truly unrecoverable (index out of range on YOUR bug, not user input).
- Expected failures (empty pop, bad config, malformed row) → return `error`, NEVER panic.
- `recover` only in a deferred top-level guard (`SafeRun`, HTTP middleware, worker wrapper turning task panics into errors).

## Building-block role
- Uniform style downstream: `ErrFull`/`ErrEmpty`/`ErrNoSection`/`ErrLine{Line,Msg}`; wrap at boundaries; handlers map errors → status codes/messages.
