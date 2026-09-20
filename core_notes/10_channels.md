# Core Note — Channels

> Prerequisite for: `phase_01_core/10_channels`. Used by: pipelines, `event` bus, `taskpool` job queue + futures, `scheduler` stop, `tui` loop.

## Essentials
- `ch := make(chan int)` unbuffered: send blocks until a receiver takes it (rendezvous). `make(chan int, 2)` buffered: blocks only when full/empty.
- Sender closes when done: `close(ch)`; receivers `for v := range ch`. **Never send on closed; never close twice; only the sender closes.**
- Receiving from closed (after drain) yields zero value + `ok==false`: `v, ok := <-ch`.
- `select` multiplexes: waits on whichever case is ready; `default` = non-blocking; `time.After`/`ctx.Done()` = timeouts/cancel.

## Patterns (memorize these four)
1. **Pipeline:** `gen → square → filter` — each stage `go func{ for v := range in { out <- f(v) } ; close(out) }`.
2. **Fan-out:** N workers reading one jobs channel. **Fan-in:** multiplex many channels into one (WaitGroup + close when all done).
3. **Done/cancel:** `done := make(chan struct{})`, `close(done)` broadcasts to all `<-done` listeners; `context.Context` is the production version.
4. **Future:** task returns a one-shot channel; `Get()` blocks for the result — the heart of `25_taskpool`.

## Building-block role
- Blocking `Ring` variant, `EventBus` subscriber channels, pool `jobs chan Task`, scraper semaphore (`sem := make(chan struct{}, 8)`), TUI `tick` + `input` select loop.
