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

# Extended Journey (Months 3–18)

> Phase 6 is now **fully specified** (`phase_06_systems/43–48`, each with `00_GUIDE.md` / `01_EXERCISES.md` / `02_QUESTIONS.md`).
> Phases 7+ remain **goals + resource pointers** until a future agent expands them with `SCAFFOLD_FOR_FUTURE_AGENTS.md`.

---

## Phase 6 — Systems, Protocols & Papers (≈ Months 3–5) — FULLY SPECIFIED

**Goal:** Learn how real systems are designed by implementing simplified-but-operational versions of papers and classic projects: your own log, store, wire protocols, consensus, and one full paper capstone.

| Module | Folder / package | Weeks | What ships |
|--------|------------------|-------|------------|
| 43 Segmented WAL | `phase_06_systems/43_wal_seglog` (`walseglog`) | 2.5–3.5 | Crash-safe segmented log: CRC framing, rotation, truncation, Replay, kill-test + sync bench |
| 44 Durable KV | `phase_06_systems/44_kv_store` (`kvstore`) | 3–4 | Bitcask/LSM-lite on 43: batches, TTL, LRU, compaction with before/after proof, CLI + bench |
| 45 Raw HTTP | `phase_06_systems/45_http_server_from_tcp` (`httpraw`) | 2–3 | HTTP/1.1 from `net.Conn` (no `net/http`): keep-alive, chunked, router, timeouts, KV service, attack + slow-client demos |
| 46 RESP + RPC | `phase_06_systems/46_resp_rpc` (`resprpc`) | 1.5–2.5 | RESP2 subset fronting 44 + length-prefixed multiplexed RPC (the 47 transport), pipelining + atomic-INCR proofs |
| 47 Raft-lite | `phase_06_systems/47_raft_lite` (`raftlite`) | 3.5–4.5 | Election + replication + commit/apply over 43+46, TCP cluster, 4-scenario chaos table passing 3× |
| 48 Paper capstone | `phase_06_systems/48_paper_reimpl` (`paperreimpl`) | 3–4 | ONE of Dynamo-lite / MapReduce-lite / GFS-lite with PAPER_NOTES + eval + failure matrix + RETRO |

**Order matters:** 43 → 44 → 45 → 46 → 47 → 48. Each module's GUIDE names the exact seam the next one consumes (`Replay`, `Open/Put/Get`, `ServeKV`, `RPCServer`, cluster helpers).

**Shippable definition:** every module has a runnable CLI/bench demo + README with real numbers from your machine; 47's chaos table and 48's eval+failure tables are the Phase 6 diploma.

**Reuse spine:** logger, config/ini, taskpool, scheduler, cache, jsonlite, channels — plus 43→44→45/46→47→48 chaining. Every module documents ≥ (growing) reuses; 48 requires ≥6.

**Resources (primary sources, per-module GUIDEs add more):**
- "In Search of an Understandable Consensus Algorithm" (Raft paper)
- Dynamo (SOSP'07), MapReduce (OSDI'04), GFS (SOSP'03) — one fully re-implemented in 48
- Redis RESP spec; Bitcask paper; SQLite architecture overview
- "The Log" (Kreps); RFC 9110/9112 framing chapters; Kafka segment design notes

---

## Phase 7 — Version Control From First Principles (≈ Months 5–7) — SCAFFOLD

**Goal:** Understand Git by rebuilding a usable subset, then add a few modern ideas. Builds directly on Phase 6: your 43 log becomes the object store journal, your 44 compaction thinking becomes GC, your 46 framing becomes the pack format.

**Staged journey (future agent: expand into ~6 modules, each with GUIDE/EXERCISES/QUESTIONS):**
1. `49_cas_store` (2 wks) — blobs/trees/commits, SHA-256 addressing, zlib-lite or stored-raw objects, `hash-object`/`cat-file` equivalents + corruption scrub.
2. `50_index_status` (2 wks) — index/staging (mtime+size cache), `.ignore` via your tokenizer, `status`/`diff` (reuse Myers-lite or simple LCS + your sort/hashmap).
3. `51_branches_porcelain` (2–3 wks) — refs/HEAD, `commit/log/checkout/branch`, CLI with your 41 patterns, file-locking for ref updates.
4. `52_merge_gc` (2 wks) — 3-way merge + conflict markers, pack/ delta-lite + `gc` (explicit nod to 44 compaction), hooks-lite + sparse-checkout-lite stretch.
5. `53_modern_capstone` (2 wks) — ONE modern idea shipped: partial clone / virtual FS overlay / GUI status view (pre-arms Phase 8) / editor integration stub (pre-arms Phase 9).

**Resources:**
- Official Git source (especially the early parts of `git/object.c`, `read-cache.c`)
- "Git Internals" / Pro Git book chapters on internals
- libgit2 documentation (concepts, not the C API)
- "Git from the Bottom Up"
- Papers / talks on content-addressable storage
- Your Phase 6 READMEs (storage + protocol lessons to reuse, not relearn)

**Shippable definition:** A CLI tool that can initialize a repo, commit files, show history, and switch branches on a small project. Put it in its own GitHub repo. Must survive kill-mid-commit and prove it with a recovery demo (same standard as 43/44).

**Reuse:** Your previous CLI, logger, config, tokenizer (for ignore patterns), set/hashmap, 43 WAL ideas, 44 compaction ideas, 46 framing for packs.

---

## Phase 8 — GUI Framework Journey (≈ Months 7–9, 3–6 focused weeks) — SCAFFOLD

**Goal:** Build a small retained-mode or immediate-mode GUI toolkit that is good enough to host real applications (first client: the Phase 9 editor; second client: a Phase 7 status viewer).

**Staged path (future agent: expand into ~5 modules):**
1. `54_window_loop` (1 wk) — window + event loop. Choose ONE backend and document why: platform (Win32/X11) raw, or pure-Go on top of Gio/Fyne/GLFW for the hard parts. Frame budget + input-event queue + shutdown discipline.
2. `55_widgets_draw` (1–2 wks) — button, label, text input, list, scroll; drawing primitives + text rendering (font atlas or system text — document choice).
3. `56_layout_input` (1 wk) — row/column (or constraint-lite) layout, focus/keyboard/mouse routing, clipboard stub, DPI/scale note.
4. `57_theming_perf` (1 wk) — theme/style separation, dirty-rect/redraw budget, 60fps list-scroll bench with 10k rows, accessibility-lite (keyboard-only operation proof).
5. `58_gui_capstone` (1 wk) — library + 3 demos (counter, form with validation, virtualized list viewer over your 44 KV). Own repo. "What I simplified" doc required.

**Resources (study designs, do not paste):**
- Gio (Go immediate-mode) design notes & source
- Fyne architecture
- Dear ImGui / "Immediate Mode GUIs" talks and papers
- "The GUI Toolkit" design discussions in various open-source projects
- Platform docs (Win32, X11/Wayland, or cross-platform abstractions)

**Shippable definition:** A small library + demo apps (button counter, form, list viewer). Own repository. Document what you deliberately simplified. List viewer must page through 10k KV rows without UI freeze (taskpool background load — reuse from 25/44).

**Reuse:** Event system, logger, config, taskpool (for background work), your earlier packages where they fit. Phase 7 status data as a demo data source.

---

## Phase 9 — Editor Built on Your Stack (≈ Months 9–11) — SCAFFOLD

**Goal:** A usable text editor that *uses* the GUI framework + previous tools. You must dogfood it: from the first multi-buffer milestone, all Phase 10+ notes and code get edited in it at least once a week.

**Staged path (future agent: expand into ~5 modules):**
1. `59_buffer_core` (2 wks) — gap buffer vs piece table vs rope: learner researches, benchmarks insert/delete on 1MB file, CHOOSES with numbers. Undo/redo (grouped), save with atomic rename + fsync (43/44 durability lesson reused).
2. `60_view_highlight` (2 wks) — multi-buffer, viewport + soft-wrap, syntax highlight for Go via your 32 tokenizer (extended), search/replace (literal + regex-lite), config via your 21 system.
3. `61_git_integration` (1–2 wks) — mini-git status/stage/commit/diff pane against Phase 7 CLI-as-library, conflict-marker navigation (from 52), large-file guard.
4. `62_lsp_perf` (1–2 wks) — diagnostics via `go vet` subprocess or LSP-stdio skim client; 100k-line file open/scroll bench; startup-time budget + plugin-hook stub.
5. `63_editor_ship` (1 wk) — keymap docs, crash-recovery (swap-file/undo-persist proof by kill test — same bar as 43), README with arch diagram.

**Resources:**
- Source of small editors (e.g. micro, helix concepts, early vim, or pure-Go editors)
- Text buffer data structures papers (piece table, rope, gap buffer)
- LSP specification (skim)

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

**Shippable definition:** You can open, edit, save, and commit a real project with it. Own repository. README shows architecture diagram and which of *your* packages it depends on. Kill-mid-edit recovery demo required.

---

## Phase 10 — Productization & Feedback Loop (≈ Months 11–14) — SCAFFOLD

**Goal:** Turn the best previous projects into proper, shippable repositories and improve them by *using* them. Output is public artifacts + a maintained best-practices file, not just more code.

**Staged activities (future agent: expand into ~4 modules with checklists, not just prose):**
1. `64_extract_harden` (2–3 wks) — pick 3 packages (expected: 44 KV, 46 RPC, 43 log OR 47 cluster) → independent Go modules with semver tags, examples/, `go vet`+race CI stub, fuzz seed corpus for protocol parsers (45/46).
2. `65_docs_release` (1–2 wks) — READMEs a stranger can follow (install → run → bench → failure demo), CHANGELOG, build/release scripts or documented steps, issue-template + roadmap per repo.
3. `66_dogfood_loop` (ongoing) — use editor + mini-git daily on this repo; file ≥10 friction notes; fix ≥5 with measured before/after (startup ms, scroll fps, commit time, bench ops/s). Keep living `BEST_PRACTICES.md` (pattern → where seen → why kept/rejected → where applied).
4. `67_perf_pass` (2 wks) — profile-guided pass on KV + GUI list + editor buffer: `pprof` CPU + alloc, bench tables before/after, compaction/GC pause note, slow-client/load re-runs proving no regression.

**Shippable definition:** At least 3 independent repositories that a stranger could clone and run, plus a short "lessons learned" document. Each repo's README must contain the same proof trio Phase 6 taught: bench numbers + failure demo + reuse map.

---

## Phase 11+ — Capstones & Depth (Months 14–18) — SCAFFOLD (expanded menu)

Pick **2–3** larger projects. Each must reuse ≥5 earlier packages, require new external reading, and reach show-in-public quality with bench + failure + reuse proofs (the 48 template applies unchanged).

**Menu (future agent: expand each chosen capstone into a 48-style GUIDE with Track/Notes/Eval/Failure/CLI):**

- **C1 Distributed toy system** — Raft-KV cluster (47+44) behind your raw HTTP (45) with Dynamo-style read-repair stretch; chaos suite extended to 6 scenarios; 3-node localhost + Docker-compose-lite deploy note.
- **C2 Build system / static site generator** — content hashing (7 CAS ideas), incremental rebuild graph (31 graph + 25 pool), file watcher (42 patterns), RESP/HTTP status server (45/46); must rebuild this repo's docs incrementally as the demo.
- **C3 Multiplayer / realtime game on your GUI** — authoritative tick (24 scheduler), snapshot interpolation, toy netcode over 46 RPC (loss/drop simulator from 47 chaos), 60fps GUI (58) stress with 500 entities.
- **C4 Time-series / OLAP-lite engine** — columnar segment files (43 segment ideas + 44 compaction), Gorilla-style float compression (paper re-impl #2), SQL-subset parser (38 + 32), bench vs SQLite on 10M points.
- **C5 Classic paper #2 at higher fidelity** — Bitcask→full LSM (Bloom + levels), Raft→snapshots + joint-consensus membership change, or GFS→HDFS-compatible mini-namenode; must include the evaluation section the original paper had, rerun at your scale.
- **C6 Contribution-ready OSS project** — take one Phase 10 repo to "accepts external PRs" bar: API docs, compatibility promise, fuzz + race CI, 3 good-first-issues, one external-user test note.

Each capstone ships in its own repo with `PAPER_OR_DESIGN_NOTES.md`, eval tables, failure matrix, and a 5-minute live-demo script (the 48 `Demo()` discipline, reused).

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
