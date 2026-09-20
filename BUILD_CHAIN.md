# Build Chain — Use What You Built (READ WITH HOW_TO_PRACTICE)

Principle: **core → your mini-stdlib → your tooling → real-world apps.** Practice, not industry-standard. Deliberately rebuild simplified versions, then *depend on your own versions* downstream.

## Layer 0 — Core only (+ `fmt` exception)
- Concept notes for every block live in `core_notes/` (start at `core_notes/00_INDEX.md`) — read the note before its exercise folder.
- Allowed everywhere: language builtins (vars, funcs, structs, slices, maps, pointers, goroutines, channels, errors) + `fmt` **for debug printing / Demo() output only**.
- NOT allowed as a crutch in library logic: if the exercise is "build X", don't just wrap `std.X` and call it done. Verifier will fail thin wrappers.
- First rite of passage: `phase_01_core/00_myfmt` — build `myprint` (Print/Println/Sprintf-lite, Itoa/Atoi-lite). After that, you *may* keep using `fmt` for debug, but you must prove you *could* live without it.

## Layer 1 — Your mini-stdlib (practice-grade, simplified)
Built from core only. Each is intentionally worse than stdlib — that's the point (you learn the internals):

| You build | Instead of | Used later by |
|---|---|---|
| `00_myfmt` (Print/Sprintf/Itoa) | `fmt` internals | `18_logger`, `40_tui` (rendering), `41_cli` (help text) |
| `04_slicesx` helpers, `05_mapsx` patterns | `slices`/`maps` pkg | everything |
| `13_list`, `14_stack`, `15_queue`, `16_set`, `17_ringbuffer` | `container/list`, std deque | `20_cache` (list+map), `25_taskpool` (queue), `29_heap` tests, `30_tree` (queue for level-order), `31_graph` (queue/stack for BFS/DFS) |
| `26_stringsx` Builder + ops | `strings.Builder` | `22_ini`, `38_parsers`, `39_scraper` (title extract) |
| `28_hashmap` | `map` internals study | `20_cache` variant, `31_graph` adjacency stretch |
| `27_sortsearch` | `sort` / `slices.Sort` | `29_heap` verification, `38_parsers` (top-N) |
| `34_iox` Reader/Writer/LimitReader/TeeReader | `io` | `22_ini` (Parse from Reader), `23_jsonlite` (lexer input), `38_parsers` (CSV/log from Reader) |
| `18_logger` | `log/slog` | `24_scheduler`, `25_taskpool`, `37_http_api`, `41_cli`, `42_final_boss` |
| `19_event` | pub/sub | `37_http_api` (request events), `42_final_boss` (chat bus option) |
| `21_config` + `22_ini` | `os.Getenv` patterns, `flag`, ini libs | `37_http_api` (port/config), `41_cli` |
| `23_jsonlite` | `encoding/json` | `37_http_api` (handlers), `38_parsers` comparison |
| `24_scheduler`, `25_taskpool` | `time.Ticker` patterns, `errgroup` | `39_scraper` (concurrent crawl), `37_http_api` (background jobs) |
| `32_tokenizer` | `text/scanner`, `encoding/csv` | `22_ini` (value lexing stretch), `38_parsers`, `39_scraper` |
| `33_unsafex`, `36_reflectx` | `unsafe`, `reflect` study | `23_jsonlite` (struct marshal), final boss |

## Layer 2 — Tooling built ON your stdlib
- `20_cache`: MUST use your `13_list` logic (copy the idea; import if you want — `import list "practice_go/phase_02_packages/13_list"`) + Go map. Document which.
- `24_scheduler`: MUST log via your `18_logger` (inject it).
- `25_taskpool`: MUST use your `15_queue` or `17_ringbuffer` as job queue + your `18_logger` for errors.
- `30_tree` LevelOrder: MUST use your queue (import or re-implement + cite).
- `31_graph` BFS/DFS: MUST use your queue/stack.
- `38_parsers`: MUST read from your `34_iox` Reader abstraction + use your `26_stringsx` Builder where strings are assembled.
- `39_scraper`: MUST use your `25_taskpool` (or its pattern) + your `16_set` for visited URLs + your `26_stringsx` for extraction.

> If you import stdlib for these (e.g. `container/list`, `encoding/json` inside impl), verifier FAILS that folder as "thin wrapper". Stdlib is allowed in *demo/test wiring* and in `main.go` scratch only — never as the implementation shortcut. Each EXERCISES file names the banned import for that folder.

## Layer 3 — Real-world apps (still on YOUR stack)
- `37_http_api`: JSON API with `net/http` + YOUR logger/cache/config/event/jsonlite. Real task: todo CRUD, curl transcript required.
- `38_parsers`: CSV + log parser producing reports. Real task: malformed-row errors with line numbers, summary stats.
- `39_scraper`: polite concurrent fetcher over `httptest` server (no internet needed). Real task: rate limit, timeout, dedup.
- `40_tui`: ANSI dashboard/progress. Real task: live rendering loop with channels.
- `41_cli`: multi-subcommand CLI wiring all of the above. Real task: `kv set/get`, `serve`, `parse`.
- `42_final_boss`: pick KV server / markdown→HTML / cron daemon / chat bus / file watcher. MUST reuse ≥5 of your packages and document the dependency list. This is the proof you can "use what you built to build complex tooling".

## Practice, not industry standard — what that means
- Simplify ruthlessly: flat JSON before nested, LRU before LFU, polling before inotify, ANSI before bubbletea, `httptest` before real web.
- Duplicate logic across folders is FINE if it teaches (e.g. re-implement queue inside graph with a comment "mirrors my 15_queue").
- Verifier checks: (1) banned-import rule respected, (2) reuse documented (import string or code comment citing source package), (3) behavior correct on edge cases.
