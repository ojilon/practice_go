# How To Practice (READ THIS FIRST)

This repo is your long-term Go dojo. You learn by **writing code yourself**, reading real systems, and shipping small but real artifacts.

The journey is designed for **12–18 months** of steady work (1–2 hours most days is enough). The first ~60 days give you a solid foundation; everything after that builds on the same reuse principle.

---

## 1. Core rules (never change)

### 1a. One rule about `main.go`
There is exactly **one** `main.go` in the root (`package main`, module `practice_go`).  
It is YOUR scratch playground. Use it to temporarily import and call whatever you are practicing. The verifier **always ignores** `main.go`.

### 1b. Core-first + build-chain rule
- Start from **core language only**. `fmt` is allowed **only for debug printing / Demo() output** after you complete `00_myfmt`.
- After that: **core → your mini-stdlib → your tooling → real apps → GUI → editor → more apps**.  
  See `BUILD_CHAIN.md` for the full dependency map and banned-import rules.
- If an exercise says "build X", do **not** just wrap the stdlib version. Thin wrappers fail verification.

### 1c. Who writes what
- The assistant (or future agents) write **guiding `.md` files**.
- **You** write **all `.go` code**.
- `doc.go` stubs are empty package declarations only.

---

## 2. Folder layout (current + future)

```
ROADMAP.md                      # Master timeline (12–18 months)
HOW_TO_PRACTICE.md              # This file
BUILD_CHAIN.md                  # Reuse + banned imports (full year)
EXPLORATION.md                  # How to read papers & real projects
SCAFFOLD_FOR_FUTURE_AGENTS.md   # Instructions for expanding later phases
VERIFIER_GUIDE.md               # How any reviewer should check work
PROGRESS.md                     # Your personal tracking (you maintain this)
core_notes/                     # Reference sheets for language core
phase_01_core/ … phase_05_apps/ # Fully specified (GUIDE + EXERCISES + QUESTIONS)
phase_06_systems/ …            # Scaffold only (goals + resources) until expanded
```

Each fully-specified exercise folder contains:
- `00_GUIDE.md` — concept + steps (simple → complex). Read first.
- `01_EXERCISES.md` — concrete code tasks.
- `02_QUESTIONS.md` — some require running code and writing answers back into the file.

Package naming rule: folder `01_variables_functions` → package `variablesfunctions` (strip leading numbers and underscores).

---

## 3. Two types of tasks

**Type A — Code-only**  
Implement the symbols. Verifier checks existence + edge-case behaviour.

**Type B — Run-and-answer**  
Write code, run it, fill the `> Your answer:` boxes in `02_QUESTIONS.md`.

---

## 4. Daily / weekly rhythm

- Most days: one focused folder or one clear sub-goal.
- Weekly: one exploration activity from `EXPLORATION.md` (read a paper, skim a real repo, take notes).
- Every 2–3 months: produce something you can put in its own repository and show someone.

---

## 5. What to create vs what NOT to create

**DO**
- Create `*.go` files inside the exercise folder.
- Edit `02_QUESTIONS.md` answer boxes.
- Temporarily edit root `main.go` to run things.
- Later: create new GitHub repositories for the stronger projects (mini-git, GUI, editor, etc.).

**DO NOT**
- Create new `go.mod` files inside this practice repo (stay in one module until you deliberately extract).
- Delete any guiding `.md` files.
- Put exercise implementation code in the root.

---

## 6. How to run

From the repo root:

```bash
go run .
go vet ./path/to/current/folder/...
```

If you get import or package-name errors, check that the package clause matches the naming rule and the import path is `practice_go/<folder>/...`.

---

## 7. Done = verifiable + usable

A folder is done when:
1. Required `.go` symbols exist and behave correctly on the stated edge cases.
2. Type-B answer boxes are filled.
3. `go vet` on that folder is clean.

A larger project (Phase 7+) is done when it reaches the **shippable definition** listed in its scaffold (runs, has a short README, demonstrates the intended capability).

---

## 8. Long-term mindset

- **Experience over perfection.** A working simplified version that you understand deeply is worth more than a half-finished perfect one.
- **Reuse is the point.** Later projects should import or clearly cite earlier ones.
- **Read real systems.** The `EXPLORATION.md` file exists so you deliberately study good (and sometimes bad) open-source code and papers.
- **Ship.** At least a few times a year put something on GitHub that a stranger could clone and run.
- **Future agents will help expand the later phases.** The scaffolds already contain the goals and resource pointers they need.

When in doubt, open `ROADMAP.md` and do the next concrete step.
