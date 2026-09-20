# Verifier Guide — For You or Another Model (DO NOT DELETE)

Purpose: let any reviewer (human, this model, or a different model) verify **folder-by-folder** whether the learner followed instructions, **independent of root `main.go`**.

## 0. Golden rules

1. **IGNORE root `main.go`, `practice_go.exe`, and `.git/.exe` artifacts.** The learner's `main.go` is a throwaway scratch runner that changes from module to module (different imports/calls each time). It may import anything, be broken, empty, or stale — even at the END of the project. NEVER pass/fail a folder based on `main.go`, NEVER ask the learner to "restore" it, NEVER diff it. If `go vet ./...` fails ONLY because of `main.go`, vet the target folder alone (`go vet ./<folder>/...`) and proceed.
2. **Scope = one folder at a time.** Example: verifying `phase_01_core/01_variables_functions` means ONLY reading files under that folder + its three `.md` guides. Do not require imports from other folders (except where an exercise explicitly says "reuse your X package" — then check import string contains that path, but still don't check `main.go`).
3. **Two evidence types:**
   - (a) `.go` source files the learner created in the folder (package clause, funcs/types, behavior).
   - (b) `02_QUESTIONS.md` answer boxes (`> Your answer:`) filled by the learner for Type B tasks.
4. **Pass criteria per folder:** ALL `MUST` checks below pass + `go vet ./<folder>/...` has no errors for that folder's packages (if no `.go` files yet → FAIL as NOT_STARTED, not as wrong).
5. **Statuses:** `PASS` / `FAIL` / `NOT_STARTED` / `PARTIAL` (list what's missing). Always give file:line references.

## 1. How to verify a folder (step-by-step for AI)

Given target folder e.g. `phase_01_core/06_interfaces`:

1. `Read` the folder dir listing. Expect `00_GUIDE.md`, `01_EXERCISES.md`, `02_QUESTIONS.md` (guides, pre-existing) + learner `.go` files.
2. `Read` `01_EXERCISES.md` → extract required symbols (funcs/types/methods listed under "Must implement" / "Must exist").
3. `Read` `02_QUESTIONS.md` → find every `> Your answer:` line. If the line after it is empty / `TODO` / placeholder → unanswered. Code-only questions say `Code-only — no answer required` → skip answer check.
4. `Glob` `*.go` in folder. If none → `NOT_STARTED`.
5. For each required symbol, `Grep` for `func <Name>` / `type <Name>` in folder. If missing → FAIL, name it.
6. Spot-check semantics: read the main `.go` file(s), check 2–3 edge cases the EXERCISES call out (e.g. Pop on empty stack returns error, LRU evicts oldest, INI ignores `;` comments). No need to run unless uncertain — but you MAY suggest `go vet` / `go run` commands for the learner to run.
7. Check `02_QUESTIONS.md` numeric answers for plausibility (don't need exact re-run; flag obvious nonsense, e.g. negative cap, empty where number required).
8. Emit report (format below).

Do NOT `Edit`/`Write` learner files. Do NOT touch `main.go`.

## 2. Per-phase MUST checklists (quick reference)

### Phase 1 core
- `00_myfmt`: package `myfmt`; `Itoa/Atoi` (sign + errors), `Fprint/Print/Println` (single Write, byte count), `Sprintf` (`%s %d %v %%`, MISSING marker); banned imports `fmt/strconv/strings/bytes/encoding/*` absent from impl files (grep); QUESTIONS answered.
- `core_notes/`: reference ONLY — never verify, never fail. Learner does not edit these.
- `01_variables_functions`: package `variablesfunctions`; funcs with multiple/named returns, variadic, error return on defined case.
- `02_structs`: `type User/Product/...` structs, composite literals, embedding, constructor `NewX` returning value or pointer consistently.
- `03_methods`: methods with value AND pointer receivers; `String() string`; chaining (return *T).
- `04_slices_arrays_strings`: demonstrates len/cap, append growth, copy, slice expressions, rune vs byte handling.
- `05_maps`: make/map literal, comma-ok, delete, iteration, struct-key or nested map.
- `06_interfaces`: interface def + at least 2 impls, type switch or assertion, `io.Writer` or `Stringer` usage.
- `07_errors`: sentinel error + `fmt.Errorf("...: %w", err)`, custom error type with `Error()`, no bare `panic` for expected errors.
- `08_pointers`: `*T` params that mutate, `new`/`&`, nil checks, pointer vs value demo.
- `09_goroutines`: `go` + `sync.WaitGroup`, no `time.Sleep` as sync (except demos), race-safe (mutex or channels).
- `10_channels`: buffered/unbuffered, `close`, `select`, pipeline or fan-in/fan-out; no goroutine leak in happy path.
- `11_defer_panic_modules`: `defer` (incl. loop/closure gotcha), `panic/recover`, multi-file package.

### Phase 2 packages
- `12_math2`: `Vec2/Vec3`, `Add/Sub/Dot/Norm`, `Mat2` Mul; QUESTIONS numeric answers filled.
- `13_list`: `Node`, `PushFront/Back`, `Remove`, `Len`, traversal; generics `List[T]` preferred.
- `14_stack`: `Push/Pop/Peek/Len/IsEmpty`, empty-pop returns error.
- `15_queue`: `Enqueue/Dequeue/Peek/Len`, deque or ring variant.
- `16_set`: `Add/Has/Remove/Size`, `Union/Intersect/Diff`, generic with `comparable`.
- `17_ringbuffer`: fixed cap, `Write/Read`, full/empty distinction, overwrite-or-error policy documented.
- `18_logger`: `Level` type, `New(w io.Writer)`, `Info/Error/Debug`, fields/structured, no `fmt.Print` in lib (only in demo via main).
- `19_event`: `Bus` with `Subscribe/Publish/Unsubscribe`, concurrent-safe.
- `20_cache`: `Get/Set`, eviction (LRU), optional TTL; uses map+list.
- `21_config`: `Load()` precedence env>flag>file>default, `GetString/GetInt` with defaults.
- `22_ini`: parses `[section]`, `key=value`, `;`/`#` comments, quoted values; returns typed getters.
- `23_jsonlite`: `Marshal`/`Unmarshal` for flat structs/maps/slices WITHOUT importing `encoding/json` in impl (tests may use it to compare).
- `24_scheduler`: `Every(d)`, `OnceAfter(d)`, `Stop()`, uses `time.Ticker` + `context`.
- `25_taskpool`: fixed workers, `Submit(func() ...)` returning Future with `Get()`, `Shutdown`, error propagation.

### Phase 3 DSA
- `26_strings_builder`: custom `Builder` (WriteString/WriteRune/String/Reset), rune-correct reverse/count.
- `27_sort_search`: ≥3 sorts + binary search, table test or benchmark file (any `*.go` showing runs).
- `28_hashmap`: buckets, hash func, insert/get/delete, resize at load factor.
- `29_heap`: `Push/Pop` maintaining heap invariant, `Top`, usable as priority queue.
- `30_tree`: BST `Insert/Search/Delete`, in/pre/post/level traversals (level uses queue).
- `31_graph`: `AddVertex/AddEdge`, BFS/DFS, cycle detect or topo sort, shortest path (BFS).
- `32_tokenizer`: token types, lexer for arithmetic/CSV/INI-subset, handles strings/escapes.

### Phase 4 systems
- `33_unsafe_memory`: `Sizeof/Alignof` answers filled, `unsafe.Pointer` + `StringHeader`/`SliceHeader` demos, explicit SAFETY comments.
- `34_io`: `Reader/Writer` interfaces + `ReadFull/Copy/LimitReader/TeeReader` rebuilds, error handling (`io.EOF`).
- `35_mini_os`: args/env/files/signals, tiny `cat`/`ls`-lite, signal handling with context.
- `36_reflection`: `reflect.TypeOf/ValueOf`, iterate struct fields, read tags, mini marshal via reflection.

### Phase 5 apps
- `37_http_api`: `net/http` server, JSON handlers, uses ≥2 of learner's packages (logger/cache/config/event).
- `38_parsers`: CSV + log parser producing structs, malformed-row errors.
- `39_scraper`: concurrent fetch with taskpool or goroutines, rate limit, context timeout, respects delay.
- `40_tui`: ANSI rendering, non-blocking input or tick loop with channels; no mandatory deps.
- `41_cli`: subcommands + flags, help text, wires other packages.
- `42_final_boss`: single app reusing ≥5 learner packages, README section describing architecture, runs via `go run .` wiring (but verify app files, not main).

## 3. QUESTIONS.md answer checking

- Pattern to find: lines starting with `> Your answer:`.
- PASS if: every such box under a question that does NOT say `Code-only` has non-empty, plausible content (numbers where numbers asked, 1+ sentences where explanation asked, code fence where code asked).
- FAIL if: any box empty / `TODO` / `...` / copy of question.
- Math/numeric (esp. `12_math2`, `27_sort_search`, `33_unsafe_memory`): flag if physically impossible (e.g. `cap` shrinking on append, `Sizeof(int)` = 1 on 64-bit) but don't demand bit-exact without running.

## 4. Report template (copy/paste per folder)

```
## Verify: <folder path>
Status: PASS|PARTIAL|FAIL|NOT_STARTED
Checked:
- files: <list .go files found>
- symbols: <required → found/missing>
- questions: <answered X/Y>
Issues:
- <file:line — what's wrong/missing>
Next:
- <concrete fix, e.g. "add func (s *Stack) Pop() (int,error) returning ErrEmpty in stack.go:12">
```

Verify ONE folder per report unless user asks for sweep. For sweep, emit one block per folder + summary table.

## 5. Common false-positives to avoid

- Do not fail for missing `package main` in exercise folders — they are libraries (`package xxx`), run via root `main.go` scratch file.
- Do not fail for unused funcs — learner may implement extra helpers.
- Do not require tests, but credit them if present.
- `go vet` failures in OTHER folders (including root `main.go`) don't fail this folder.

## 6. Build-chain rules (per BUILD_CHAIN.md)

- **Banned-import check:** each folder's `01_EXERCISES.md` names imports forbidden in impl files (e.g. no `encoding/json` in `23_jsonlite`, no `container/list` in `13_list`/`20_cache`, no `fmt/strconv` in `00_myfmt`). `Grep` the folder's non-demo `.go` files for those import strings. Thin wrapper over stdlib = FAIL with note "reimplement, don't wrap".
- Demo/comparison funcs MAY import the banned package to cross-check outputs — allow if isolated in a `demo*.go`/`compare` func and the real logic is hand-rolled.
- **Reuse check (Layer 2+):** folders `20_cache`, `24_scheduler`, `25_taskpool`, `30_tree`, `31_graph`, `38_parsers`, `39_scraper`, `37_http_api`, `41_cli`, `42_final_boss` must document reuse: either an import of `practice_go/...` own package OR a code comment citing the mirrored package (e.g. `// mirrors 15_queue ring logic`). Missing both = PARTIAL with note "state which own-package you reused".
- `core_notes/*.md` are assistant-written references: never verify, never require edits.
