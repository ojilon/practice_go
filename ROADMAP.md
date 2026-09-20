# Roadmap — ~60 Days, Simple → Complex → Insane

Do in order. Times are estimates. Slow is fine — this is designed for weeks of practice.

## Phase 0 — Setup (Day 0)
- [x] Read `HOW_TO_PRACTICE.md`
- [ ] Read `go_core_guide.txt` (seed)
- [ ] Skim `core_notes/00_INDEX.md` (reference sheets for every core block — revisit per folder)
- [ ] Read `BUILD_CHAIN.md` (core → your mini-stdlib → your tooling → real apps; banned-import rules)
- [ ] Make `go run .` work from root (your scratch `main.go` — changes every module, never verified)

## Phase 1 — Core Go (Days 1–14)
| Day | Folder | Focus |
|-----|--------|-------|
| 0 | `phase_01_core/00_myfmt` + `core_notes/01_variables_functions.md` | earn `fmt` debug rights: Itoa/Atoi, Print, Sprintf-lite (core only, no `fmt` in impl) |
| 1 | `phase_01_core/01_variables_functions` | vars, consts, funcs, multiple returns, named returns |
| 2 | `phase_01_core/02_structs` | structs, embedding, tags, constructors |
| 3 | `phase_01_core/03_methods` | value vs pointer receivers, String(), fluent builders |
| 4 | `phase_01_core/04_slices_arrays_strings` | arrays, slices, cap/len, copy, strings/runes |
| 5 | `phase_01_core/05_maps` | maps, zero values, iteration order, struct keys |
| 6 | `phase_01_core/06_interfaces` | interfaces, type switch, io.Reader, empty any |
| 7 | `phase_01_core/07_errors` | errors, wrap (%w), custom types, panic/recover |
| 8 | `phase_01_core/08_pointers` | pointers, new/make, lifetimes, pointer vs value |
| 9 | `phase_01_core/09_goroutines` | go, WaitGroup, race, GOMAXPROCS |
| 10–11 | `phase_01_core/10_channels` | unbuffered/buffered, select, close, pipelines, fan-in/out |
| 12 | `phase_01_core/11_defer_panic_modules` | defer, panic/recover, packages, go.mod, control flow deep-dive |
| 13–14 | Buffer / catch-up + redo weakest folder | — |

## Phase 2 — Rebuild Packages From Core (Days 15–32)
These are the `go_core_guide.txt` packages + extras. Each forces structs/interfaces/errors/pointers/slices/maps in real designs.

| Day | Folder | What you build |
|-----|--------|----------------|
| 15–16 | `phase_02_packages/12_math2` | Vec2/Vec3, Mat2x2, dot/cross/norm, QUESTIONS need numeric answers |
| 17 | `phase_02_packages/13_list` | Singly + doubly linked list with generics |
| 18 | `phase_02_packages/14_stack` | Slice-backed + linked stack, LIFO, min-stack stretch |
| 19 | `phase_02_packages/15_queue` | Queue + deque (ring), FIFO |
| 20 | `phase_02_packages/16_set` | Generic Set[T comparable], union/intersect/diff |
| 21 | `phase_02_packages/17_ringbuffer` | Fixed ring, overwrite vs error, blocking variant with channels |
| 22 | `phase_02_packages/18_logger` | Levels, structured fields, io.Writer injection, hooks |
| 23–24 | `phase_02_packages/19_event` | EventBus: Subscribe/Publish, generics or string topics |
| 25 | `phase_02_packages/20_cache` | LRU cache with map+list, TTL variant |
| 26 | `phase_02_packages/21_config` | Env + flags + defaults, precedence chain |
| 27 | `phase_02_packages/22_ini` | INI parser (sections/keys/comments) |
| 28–29 | `phase_02_packages/23_jsonlite` | Simplified JSON serializer/deserializer (reflection-lite, no encoding/json in impl) |
| 30 | `phase_02_packages/24_scheduler` | Ticker scheduler, one-shot + recurring, stop contexts |
| 31–32 | `phase_02_packages/25_taskpool` | Worker pool, futures, errgroup-style |

## Phase 3 — DSA + Rebuild Stdlib (Days 33–44)
| Day | Folder |
|-----|--------|
| 33 | `phase_03_dsa/26_strings_builder` — rune handling, Builder from scratch, KMP stretch |
| 34–35 | `phase_03_dsa/27_sort_search` — bubble/insert/merge/quicksort, binary search, benchmarks |
| 36 | `phase_03_dsa/28_hashmap` — hash table from slices+lists, load factor, resizing |
| 37 | `phase_03_dsa/29_heap` — heap/priority queue (container/heap-free impl + interface comparison) |
| 38–39 | `phase_03_dsa/30_tree` — BST, traversal, level-order with your queue |
| 40–41 | `phase_03_dsa/31_graph` — adjacency list, BFS/DFS, topo sort, shortest path |
| 42 | `phase_03_dsa/32_tokenizer` — lexer/tokenizer/mini-regex (*, +, ?) |
| 43–44 | Buffer — benchmark everything, compare to stdlib |

## Phase 4 — Systems Low-Level (Days 45–50)
| Day | Folder |
|-----|--------|
| 45 | `phase_04_systems/33_unsafe_memory` — unsafe.Pointer, Sizeof/Alignof, string↔[]byte zero-copy (read docs quest) |
| 46 | `phase_04_systems/34_io` — Reader/Writer/Closer from scratch, TeeReader, LimitReader rebuilds |
| 47 | `phase_04_systems/35_mini_os` — env, args, files, signals, exec, tiny shell pieces (small/safe subset) |
| 48 | `phase_04_systems/36_reflection` — reflect basics, struct-tag decoder, mini-marshaler |
| 49–50 | Stretch: arena allocator / slab / object pool with sync.Pool comparison |

## Phase 5 — Real Apps From Your Packages (Days 51–60+)
| Day | Folder |
|-----|--------|
| 51–52 | `phase_05_apps/37_http_api` — net/http CRUD API using your logger/cache/config/event |
| 53 | `phase_05_apps/38_parsers` — CSV + log parser (uses your tokenizer, ini, jsonlite) |
| 54 | `phase_05_apps/39_scraper` — concurrent fetcher with taskpool + rate limit + robots respect |
| 55–56 | `phase_05_apps/40_tui` — ANSI TUI (list viewer, progress bars) with goroutines/channels; bubbletea stretch |
| 57 | `phase_05_apps/41_cli` — CLI tool (flags/subcommands) wiring all packages |
| 58–60+ | `phase_05_apps/42_final_boss` — pick one: KV server, markdown-to-html, cron daemon, chat bus, file watcher. Must reuse ≥5 of your packages. |

## Quests (interleaved, see each `02_QUESTIONS.md`)
- Q-DOCS: go read `go doc io.Reader`, `sync.WaitGroup`, `encoding/json`, `net/http`, `unsafe`, `reflect` then come back and implement a mini version.
- Q-API: build a JSON API with only `net/http` + your packages.
- Q-PARSER: parse INI/CSV/logs without `encoding/*` in impl.
- Q-SCRAPER: polite concurrent scraper (rate-limited, context-cancelled).
- Q-TUI: raw ANSI TUI, no deps.
- Q-UNSAFE: answer Sizeof puzzles by running code.
- Q-BOSS: final integration.

## Completion checklist
- Each folder: `.go` files present + `02_QUESTIONS.md` answers filled (if Type B) + `go vet` clean.
- Verify folder-by-folder with `VERIFIER_GUIDE.md`. Do not skip Phase 1 — everything compounds.
