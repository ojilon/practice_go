# Roadmap — 12 to 18 Month Go Journey

**Mindset first:** This is a long practice dojo, not a sprint to perfection.  
Goal = **experience + shippable artifacts**. Slow is expected. Skipping Phase 1 is not.

Times are rough estimates assuming ~1–2 hours most days. Adjust freely.

---

## High-level arc

| Months | Phase | What you gain |
|--------|-------|---------------|
| 0–2    | 0–5   | Core language + your own mini-stdlib + first real apps |
| 3–5    | 6     | Systems thinking, papers, protocols |
| 5–7    | 7     | Version control from first principles (mini-git + modern ideas) |
| 7–9    | 8     | GUI framework (3–6 weeks of focused building) |
| 9–11   | 9     | Text editor built on *your* GUI + previous tools |
| 11–14  | 10    | Productization, multi-repo, feedback loops |
| 14–18  | 11+   | Capstones, deeper papers, domain projects |

Everything later **reuses** earlier work the same way Phase 5 reuses Phase 2.

---

## Phase 0 — Setup (Day 0)
- [ ] Read `HOW_TO_PRACTICE.md` (updated)
- [ ] Read `EXPLORATION.md` (how to read real projects & papers)
- [ ] Read `BUILD_CHAIN.md` (reuse rules — now covers the full year)
- [ ] Skim `core_notes/00_INDEX.md`
- [ ] Read `SCAFFOLD_FOR_FUTURE_AGENTS.md` (so you know what comes later)
- [ ] Make `go run .` work from root
- [ ] Add a proper `.gitignore` (template provided)

---

## Phase 1 — Core Go (Days 1–14)
| Day | Folder | Focus |
|-----|--------|-------|
| 0 | `phase_01_core/00_myfmt` + `core_notes/01_variables_functions.md` | earn `fmt` debug rights: Itoa/Atoi, Print, Sprintf-lite (core only) |
| 1 | `phase_01_core/01_variables_functions` | vars, consts, funcs, multiple/named returns |
| 2 | `phase_01_core/02_structs` | structs, embedding, tags, constructors |
| 3 | `phase_01_core/03_methods` | value vs pointer receivers, String(), fluent builders |
| 4 | `phase_01_core/04_slices_arrays_strings` | arrays, slices, cap/len, copy, strings/runes |
| 5 | `phase_01_core/05_maps` | maps, zero values, iteration order, struct keys |
| 6 | `phase_01_core/06_interfaces` | interfaces, type switch, io.Reader, empty any |
| 7 | `phase_01_core/07_errors` | errors, wrap (%w), custom types, panic/recover |
| 8 | `phase_01_core/08_pointers` | pointers, new/make, lifetimes |
| 9 | `phase_01_core/09_goroutines` | go, WaitGroup, race, GOMAXPROCS |
| 10–11 | `phase_01_core/10_channels` | unbuffered/buffered, select, close, pipelines |
| 12 | `phase_01_core/11_defer_panic_modules` | defer, panic/recover, packages, go.mod |
| 13–14 | Buffer / catch-up | redo weakest folder |

---

## Phase 2 — Rebuild Packages From Core (Days 15–32)
| Day | Folder | What you build |
|-----|--------|----------------|
| 15–16 | `12_math2` | Vec2/Vec3, Mat2x2, dot/cross/norm |
| 17 | `13_list` | Singly + doubly linked list (generics) |
| 18 | `14_stack` | Slice-backed + linked stack, min-stack stretch |
| 19 | `15_queue` | Queue + deque (ring) |
| 20 | `16_set` | Generic Set[T comparable] |
| 21 | `17_ringbuffer` | Fixed ring, overwrite vs error, channel variant |
| 22 | `18_logger` | Levels, structured fields, io.Writer injection |
| 23–24 | `19_event` | EventBus: Subscribe/Publish |
| 25 | `20_cache` | LRU (map+list), TTL variant |
| 26 | `21_config` | Env + flags + defaults, precedence |
| 27 | `22_ini` | INI parser |
| 28–29 | `23_jsonlite` | Simplified JSON (no encoding/json in impl) |
| 30 | `24_scheduler` | Ticker + one-shot + context |
| 31–32 | `25_taskpool` | Worker pool, futures |

---

## Phase 3 — DSA + Rebuild Stdlib Pieces (Days 33–44)
| Day | Folder |
|-----|--------|
| 33 | `26_strings_builder` |
| 34–35 | `27_sort_search` |
| 36 | `28_hashmap` |
| 37 | `29_heap` |
| 38–39 | `30_tree` |
| 40–41 | `31_graph` |
| 42 | `32_tokenizer` |
| 43–44 | Buffer + benchmarks |

---

## Phase 4 — Systems Low-Level (Days 45–50)
| Day | Folder |
|-----|--------|
| 45 | `33_unsafe_memory` |
| 46 | `34_io` |
| 47 | `35_mini_os` |
| 48 | `36_reflection` |
| 49–50 | Stretch: arena / slab / object pool |

---

## Phase 5 — Real Apps From Your Packages (Days 51–60+)
| Day | Folder |
|-----|--------|
| 51–52 | `37_http_api` |
| 53 | `38_parsers` |
| 54 | `39_scraper` |
| 55–56 | `40_tui` |
| 57 | `41_cli` |
| 58–60+ | `42_final_boss` (reuse ≥5 of your packages) |

**At the end of Phase 5 you should have several small but real tools that run.**

---

# Extended Journey (Months 3–18) — Scaffold Level

> The folders below do **not** yet contain full `00_GUIDE.md` / `01_EXERCISES.md` / `02_QUESTIONS.md`.  
> They exist as **goals + resource pointers**. A future agent (or you) will expand them using `SCAFFOLD_FOR_FUTURE_AGENTS.md`.

---

## Phase 6 — Systems, Protocols & Papers (≈ Months 3–5)

**Goal:** Learn how real systems are designed by implementing simplified versions of papers and classic projects.

**Suggested modules (to be fleshed out):**
- `43_kv_store` — simple key-value with WAL or append-only log
- `44_raft_lite` or `paxos_lite` — leader election + log replication (very simplified)
- `45_http_server_from_tcp` — build a tiny HTTP/1.1 server on net.Conn (no net/http)
- `46_protocol_parser` — parse a real binary or text protocol (e.g. Redis RESP subset, or a toy RPC)
- `47_paper_reimpl` — pick one paper from the reading list and implement its core algorithm

**Resources to start from (do not copy blindly — extract ideas):**
- "In Search of an Understandable Consensus Algorithm" (Raft paper)
- Redis design notes / RESP protocol
- SQLite architecture overview (for the idea of a simple storage engine)
- "The Log: What every software engineer should know about real-time data's unifying abstraction"

**Shippable definition:** At least two modules that can be demonstrated end-to-end and have their own short README.

**Reuse:** Your logger, taskpool, config, jsonlite, channels, etc.

---

## Phase 7 — Version Control From First Principles (≈ Months 5–7)

**Goal:** Understand Git by rebuilding a usable subset, then add a few modern ideas.

**Core journey:**
1. Content-addressable storage (blobs, trees, commits)
2. Index / staging area
3. Branches & HEAD
4. Basic `add`, `commit`, `log`, `checkout`, `status`, `diff`
5. Simple merge (or at least detect conflicts)
6. Stretch modern features: better status, partial ideas from sparse-checkout, or a simple hook system

**Resources:**
- Official Git source (especially the early parts of `git/object.c`, `read-cache.c`)
- "Git Internals" / Pro Git book chapters on internals
- libgit2 documentation (concepts, not the C API)
- "Git from the Bottom Up"
- Papers / talks on content-addressable storage

**Shippable definition:** A CLI tool that can initialize a repo, commit files, show history, and switch branches on a small project. Put it in its own GitHub repo.

**Reuse:** Your previous CLI, logger, config, tokenizer (for ignore patterns), set/hashmap, etc.

---

## Phase 8 — GUI Framework Journey (≈ Months 7–9, 3–6 focused weeks)

**Goal:** Build a small retained-mode or immediate-mode GUI toolkit that is good enough to host real applications.

**Suggested path:**
1. Window + event loop (platform backend or pure-Go with a known library for the hard parts)
2. Basic widgets: button, label, text input, list, scroll
3. Layout system (row/column or simple constraints)
4. Drawing primitives + text rendering
5. Focus, keyboard, mouse handling
6. Theme / style separation

**Resources (study designs, do not paste):**
- Gio (Go immediate-mode) design notes & source
- Fyne architecture
- Dear ImGui / "Immediate Mode GUIs" talks and papers
- "The GUI Toolkit" design discussions in various open-source projects
- Platform docs (Win32, X11/Wayland, or cross-platform abstractions)

**Shippable definition:** A small library + demo apps (button counter, form, list viewer). Own repository. Document what you deliberately simplified.

**Reuse:** Event system, logger, config, taskpool (for background work), your earlier packages where they fit.

---

## Phase 9 — Editor Built on Your Stack (≈ Months 9–11)

**Goal:** A usable text editor that *uses* the GUI framework + previous tools.

**Must have:**
- Multi-buffer / multi-file
- Basic editing (insert, delete, undo/redo)
- Syntax highlighting for at least one language (reuse or extend your tokenizer)
- Integration with your mini-git (status, stage, commit from the editor)
- Search / replace
- Configuration via your config system

**Nice-to-have (stretch):**
- Simple language server protocol client or built-in diagnostics
- Split panes
- Plugin-style extension points

**Resources:**
- Source of small editors (e.g. micro, helix concepts, early vim, or pure-Go editors)
- Text buffer data structures papers (piece table, rope, gap buffer)
- LSP specification (skim)

**Shippable definition:** You can open, edit, save, and commit a real project with it. Own repository. README shows architecture diagram and which of *your* packages it depends on.

---

## Phase 10 — Productization & Feedback Loop (≈ Months 11–14)

**Goal:** Turn the best previous projects into proper, shippable repositories and improve them by *using* them.

**Activities:**
- Extract the strongest packages into their own Go modules / repos
- Write proper READMEs, examples, and basic tests
- Use the editor + mini-git daily on the practice_go repo itself
- Collect best practices from the open-source projects you studied (keep a living `BEST_PRACTICES.md`)
- Improve performance / UX of the GUI and editor based on real use
- Add packaging / release scripts (or at least documented build steps)

**Shippable definition:** At least 3 independent repositories that a stranger could clone and run, plus a short "lessons learned" document.

---

## Phase 11+ — Capstones & Depth (Months 14–18)

Pick **2–3** larger projects. Examples (choose or invent):

- Distributed toy system (using your Raft/KV ideas)
- Domain-specific tool (e.g. static site generator, build system, or game with the GUI)
- Re-implementation of a classic paper at higher fidelity
- Contribution-ready open-source style project that reuses your stack heavily

Each capstone should:
- Force reuse of ≥5 earlier packages
- Require reading new external material
- Reach a state where you would be comfortable showing it publicly

---

## Quests that run the whole year

- **Q-READ:** Regularly read one real open-source project or paper and write 5–10 bullet notes of what you would steal / avoid.
- **Q-REUSE:** Every new project must list which of *your* previous packages it depends on.
- **Q-SHIP:** At least every 2–3 months produce something you can put on GitHub with a clear README.
- **Q-REFLECT:** Keep a short journal (or just notes in PROGRESS.md) of what felt hard and what clicked.

---

## Completion mindset

You are done with a phase when:
1. The artifacts run and demonstrate the intended capability.
2. You can explain the design decisions and the trade-offs you made.
3. Downstream work can import or copy the ideas without you having to rewrite everything.

Perfection is the enemy. Experience and shipping are the goals.
