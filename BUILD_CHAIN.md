# Build Chain — Use What You Built

Principle: **core → your mini-stdlib → your tooling → real apps → GUI → editor → more apps**.  
Practice, not industry-standard. Deliberately rebuild simplified versions, then *depend on your own versions* downstream.

---

## Layer 0 — Core only (+ `fmt` exception)

- Concept notes live in `core_notes/`.
- Allowed everywhere: language builtins + `fmt` **for debug printing / Demo() output only** (after completing `00_myfmt`).
- First rite of passage: `phase_01_core/00_myfmt`.

---

## Layer 1 — Your mini-stdlib (practice-grade)

Built from core only. Intentionally simpler than the real stdlib.

| You build | Instead of | Used later by |
|-----------|------------|---------------|
| `00_myfmt` | `fmt` internals | logger, TUI, CLI, editor help |
| list / stack / queue / set / ringbuffer | `container/*` | cache, taskpool, tree, graph, editor data structures |
| strings builder + ops | `strings.Builder` | parsers, scraper, editor buffer helpers |
| hashmap / sort | map internals / `sort` | many later structures |
| logger | `log/slog` | almost everything after Phase 2 |
| event bus | pub/sub | HTTP API, editor events, GUI |
| config + ini | flag + env patterns | every app |
| jsonlite | `encoding/json` | HTTP handlers, config, editor settings |
| scheduler + taskpool | time.Ticker + errgroup | scraper, background jobs, editor workers |
| tokenizer | scanners | parsers, syntax highlighting, ignore files |
| io rebuilds | parts of `io` | everything that reads/writes |

---

## Layer 2 — Tooling built ON your stdlib

- Cache must use your list + map ideas.
- Scheduler and taskpool must use your logger (and preferably your queue/ringbuffer).
- Tree level-order and graph BFS/DFS must use your queue/stack.
- Parsers and scraper must use your io, strings, tokenizer, taskpool, set.
- HTTP API, CLI, final boss must wire multiple of the above.

Banned imports are listed in each folder’s `01_EXERCISES.md`. Thin wrappers = fail.

---

## Layer 3 — Real apps (Phase 5)

HTTP API, parsers, scraper, TUI, CLI, final boss — all on *your* stack.

---

## Layer 4 — Systems & papers (Phase 6)

KV store, consensus lite, protocol parsers, paper re-implementations.  
Must reuse logger, config, channels, taskpool, jsonlite, etc.

---

## Layer 5 — Version control (Phase 7)

Mini-git (and modern extensions) lives in its own repository.  
Reuse: CLI patterns, logger, config, set/hashmap, tokenizer (for .gitignore), previous data structures.

---

## Layer 6 — GUI framework (Phase 8)

Your GUI toolkit becomes a foundation library (own repo).  
Reuse: event bus, logger, config, taskpool (background work), any drawing helpers you built.

---

## Layer 7 — Editor (Phase 9)

The editor is the first major application of the GUI.  
It **must** integrate:
- the GUI framework
- mini-git (status / stage / commit)
- your tokenizer or highlighter
- config + logger

This is the strongest demonstration of the whole build-chain idea.

---

## Layer 8 — Productization & feedback (Phase 10+)

- Extract strong packages into independent modules/repos.
- Use the editor + mini-git on your own work daily.
- Improve earlier packages based on real friction.
- Collect best practices from the projects you studied.

---

## Practice rules that never change

1. Simplify ruthlessly. Flat before nested, polling before fancy watchers, ANSI or simple retained-mode before full native polish.
2. Document reuse: either an import of `practice_go/...` or a clear code comment citing the package you mirrored.
3. Verifier (and future you) checks banned imports and documented reuse.
4. Experience and a shippable artifact beat a perfect unfinished one.
