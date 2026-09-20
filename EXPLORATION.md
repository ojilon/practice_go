# Exploration — How to Learn From Real Systems

This dojo is not only about writing your own code. A large part of the growth comes from **studying** real projects, papers, and design notes, then deliberately borrowing (or rejecting) ideas.

Do one exploration activity every week or two. Keep short notes (even 5–10 bullets) in `PROGRESS.md` or a personal journal.

---

## 1. How to read a real codebase

1. Start with the README and any architecture / design document.
2. Find the entry points (main, server start, command registration).
3. Trace one happy-path feature end-to-end.
4. Look for the data structures that hold the important state.
5. Notice error handling, logging, configuration, and concurrency patterns.
6. Ask: "What would I steal for my own projects? What would I avoid?"

You do **not** need to understand every line.

---

## 2. Suggested starting points (pick according to the phase you are in)

**Language & packages**
- Go standard library source (especially `fmt`, `io`, `net/http`, `sync`, `encoding/json`)
- Small, high-quality pure-Go libraries

**Systems & storage**
- Redis design notes / RESP protocol
- SQLite architecture overview
- "The Log…" essay by Jay Kreps
- Raft paper + any simplified educational implementation

**Version control**
- Pro Git book (internals chapters)
- Early Git source or well-commented educational Git reimplementations
- libgit2 concepts (not the C API itself)

**GUI**
- Gio design and source (Go)
- Fyne architecture
- Dear ImGui talks and "Immediate Mode GUI" literature
- Any small retained-mode toolkit design notes

**Editors & text**
- Data structures for text buffers (gap buffer, piece table, rope papers/articles)
- Small editors written in Go or other systems languages
- LSP specification (skim only)

**General craft**
- "The Practice of Programming", "A Philosophy of Software Design", or similar short books
- Post-mortems and design retrospectives of tools you already use

---

## 3. Papers & primary sources

Prefer the original paper or the original design document over secondary blog posts when possible.  
When an exercise points you at a paper, read the abstract + introduction + the core algorithm section, then try to implement a tiny version.

---

## 4. Collecting best practices

Maintain a living file (suggested name `BEST_PRACTICES.md` in this repo or a personal notes repo) with short entries such as:

- Pattern observed
- Where you saw it
- Why it is useful (or why you rejected it)
- Where you applied (or plan to apply) it in your own code

This becomes invaluable in Phase 10 when you productize your work.

---

## 5. Rules of exploration

- Never paste large chunks of external code into your practice projects.
- Always attribute ideas in comments or README when you deliberately copy a design.
- Prefer understanding + re-implementation over dependency on the original library (until Phase 10+ where you may consciously choose to depend on mature libraries).
- Exploration is mandatory for the later phases; the scaffolds already list the key resources.

The goal is not to become a copy of any existing project. The goal is to absorb the thinking behind good systems so that your own simplified versions are solid.
