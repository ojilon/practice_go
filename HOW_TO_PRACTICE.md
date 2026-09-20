# How To Practice (READ THIS FIRST)

This repo is your Go dojo. You learn by **writing code yourself**, not by reading.

## 1. The one rule about `main.go`

## 1b. Core-first rule + `fmt` exception + build chain

- Start from **core only**. `fmt` is allowed **for debug printing / `Demo()` output only** — and you must also prove you can live without it in `phase_01_core/00_myfmt` (build `myprint`: Print/Println/Sprintf-lite, Itoa/Atoi-lite).
- After that: **core → your mini-stdlib → your tooling → real-world apps**. See `BUILD_CHAIN.md` for the full dependency map (what reuses what, banned imports per folder).
- Rule per folder: if the exercise says "build X", don't just wrap `std.X`. Thin wrappers FAIL verification. Simplified, practice-grade rebuilds are the point.
- I (the assistant) write **only guiding `.md` files**. You write **all `.go` code** yourself. The `doc.go` stubs in each folder are empty package declarations only — no solutions — so `go vet` passes before you start.

There is exactly **one** `main.go` in the root (`package main`, `module practice_go`).

It is YOUR scratch playground. Use it interchangeably to execute code from different subfolders:

```go
package main

import (
  "fmt"
  "practice_go/basics"
  "practice_go/phase_01_core/01_variables_functions" // example — change this line as you practice
)

func main() {
  fmt.Println("scratch run")
  // call whatever you are currently practicing:
  // variablesfunctions.Demo()
}
```

Workflow per exercise:
1. Create your `.go` files **inside the exercise folder** (e.g. `phase_01_core/01_variables_functions/`). Never put exercise code in root.
2. Temporarily import that package in `main.go`, add calls in `func main()`, run `go run .` from root.
3. Verify output by eye, iterate.
4. **Leave your exercise `.go` files in place.** Revert `main.go` to a clean/minimal state before asking for verification (or leave it — verifier IGNORES `main.go`).

> Verifier rule: marking happens **independent of `main.go`**. The verifier (human, you, or another model) must NEVER fail you because of `main.go` contents. It only checks the exercise folder.

## 2. Folder layout

```
ROADMAP.md                  # 60-day plan, what order to do
HOW_TO_PRACTICE.md          # this file
VERIFIER_GUIDE.md           # instructions for an AI verifier (don't edit, just follow)
go_core_guide.txt           # original seed list
basics/                     # your existing work — keep it
phase_01_core/              # Variables → Channels (foundations)
phase_02_packages/          # math2, list, stack, ... taskpool (rebuild stdlib pieces)
phase_03_dsa/               # strings, sort/search, hashmap, heap, trees, graphs
phase_04_systems/           # unsafe, io, reflection, mini-os
phase_05_apps/              # http api, parsers, scraper, TUI, CLI, final boss
```

Each exercise folder (e.g. `phase_01_core/01_variables_functions/`) contains:
- `00_GUIDE.md` — learn this concept, simple → complex steps. Read it first.
- `01_EXERCISES.md` — code tasks E1..En. You create `*.go` files (package name = folder short name) and run via `main.go`.
- `02_QUESTIONS.md` — some require you to **run code and write the answer back into this .md** (look for `> Your answer:`). Others are code-only (`Code-only — no answer to fill`).

Package naming: folder `01_variables_functions` → package `variablesfunctions` (no digits/underscores in package name). Same rule everywhere: strip leading `NN_` and underscores. Example: `12_math2` → package `math2`, `13_list` → package `list`, etc. Each GUIDE states the exact package name to use.

## 3. Two types of tasks

**Type A — Code-only:** "Write the code, run it, leave it."
- Example: implement `Stack.Push/Pop`. Verifier checks the `.go` file exists, has right funcs, handles edge cases. No md editing needed.

**Type B — Run-and-answer:** "Write code, run it, come back and edit the md."
- Example: "What is `cap` after 3 appends? Run your program and paste the numbers after `> Your answer:`."
- You MUST edit `02_QUESTIONS.md` in place. Do not create a separate answers file. Verifier checks both code AND filled answers.

## 4. Daily rhythm (for 1–2 months)

- 1 folder per day is plenty (some big folders = 2–3 days, see ROADMAP.md).
- Read `00_GUIDE.md` (15 min) → Do `01_EXERCISES.md` (60–90 min) → Do `02_QUESTIONS.md` (20 min) → Run via `main.go`.
- Weekly: one QUEST (read docs / build API / parser / scraper / TUI).
- Do NOT use external deps except where explicitly allowed (`phase_05_apps` may use stdlib only unless noted; TUI exercise is stdlib-only with `os/exec` + ANSI, no bubbletea needed — stretch goal allows it).

## 5. What to create vs what NOT to create

- DO create: `*.go` files inside the exercise folder. One file per exercise is fine (`e01.go`, `stack.go`, etc.) or one combined file. Your choice, unless EXERCISES says otherwise.
- DO edit: `02_QUESTIONS.md` answer boxes.
- DO temporarily edit: root `main.go` to run things.
- DO NOT create: new `go.mod` files (stay in one module `practice_go`), new root folders, or `*_test.go` unless the exercise asks (you can, but verifier doesn't require them).
- DO NOT delete: any `*.md` guide files.

## 6. How to run

From root `D:\projects\go_practice`:

```powershell
go run .
# or for a subpackage directly (if it has its own main — most don't):
go vet ./...
go run .
```

If you get `import cycle` or `package not found`, check your import path is `practice_go/<folder>/<subfolder>` and package clause matches.

## 7. Done = verifiable

You are done with a folder when:
1. All `01_EXERCISES.md` files exist as `.go` code and run without panic (via your own `main.go` wiring).
2. All `02_QUESTIONS.md` answer boxes are filled (if any).
3. `go vet ./<your-folder>/...` passes.

Then ask verifier (AI) to check that folder per `VERIFIER_GUIDE.md`, folder-by-folder.
