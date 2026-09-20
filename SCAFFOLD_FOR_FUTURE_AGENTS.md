# Scaffold for Future Agents

This file tells any future agent (or human) how to expand the later phases of the journey.

The first five phases already have full `00_GUIDE.md`, `01_EXERCISES.md`, and `02_QUESTIONS.md` files.  
Phases 6+ currently exist only as **goals + resource pointers** in `ROADMAP.md`. Your job is to turn those goals into the same three-file structure used by the early phases.

---

## What a complete exercise folder looks like

For every new module you create under `phase_06_…`, `phase_07_…`, etc.:

```
phase_XX_name/NN_shortname/
  00_GUIDE.md          # Why it exists, simple→complex steps, how to run, done-when criteria
  01_EXERCISES.md      # Concrete "Must implement" symbols + edge cases + banned imports
  02_QUESTIONS.md      # Mix of code-only and run-and-answer questions
  doc.go               # package declaration only (empty)
```

Keep the same tone as the existing guides:
- Short, direct, practical.
- No hand-holding "how to write the algorithm" — point to resources and let the learner figure it out.
- Explicit banned imports.
- Explicit reuse requirements ("must import or clearly cite your X package").
- Clear "Done when" checklist.

---

## How to expand a phase

1. Read the phase description in `ROADMAP.md`.
2. Read `BUILD_CHAIN.md` so you know what must be reused.
3. Read `EXPLORATION.md` so you know the spirit of external study.
4. Create the folder structure and the three markdown files.
5. In the GUIDE, list concrete resources (papers, repos, RFCs, books) the learner should look at. Prefer primary sources.
6. In the EXERCISES, define a minimal viable set of types/functions that prove the concept, plus 1–2 stretch goals.
7. Keep the scope shippable: a working demo is more important than feature completeness.

---

## Phase-by-phase guidance for the agent

### Phase 6 — Systems & Papers
- Create 4–6 modules covering storage, consensus, protocols, and at least one paper re-implementation.
- Emphasize "read the paper / design notes first, then implement a tiny version".
- Require reuse of logger, config, taskpool, channels, jsonlite.

### Phase 7 — Version Control
- Structure the journey as successive capabilities (content-addressable store → index → branches → basic porcelain → one modern idea).
- Point heavily at Git internals resources and early Git source.
- The final artifact should be a separate repository the learner creates.

### Phase 8 — GUI Framework
- 3–6 week focused path: window/event loop → widgets → layout → input → theming.
- Point at Gio, Fyne, Dear ImGui design notes, and "Immediate Mode GUI" literature.
- Explicitly allow a platform backend or a pure-Go approach; the learner chooses.
- End product = library + demos in its own repo.

### Phase 9 — Editor
- Must consume the GUI framework and the mini-git.
- Core editing model (gap buffer / piece table / rope — learner researches and chooses).
- Syntax highlighting via the earlier tokenizer or a new simple one.
- Shippable = can edit and commit real code with it.

### Phase 10 — Productization
- Less "new code", more "extract, document, improve, ship".
- Guide the learner to create proper READMEs, examples, and independent modules.
- Include a living best-practices collection activity.

### Phase 11+
- Capstone projects chosen by the learner from a menu of suggestions.
- Each must force significant reuse and external reading.

---

## Style rules for all new guides

- Speak directly to the learner.
- Prefer "build a simplified X that demonstrates Y" over "implement the full specification".
- Always state the banned imports and the required reuse.
- Always give a concrete "Done when" definition that includes a runnable demo.
- Never write the solution code yourself in the guides.

---

## When you finish expanding a phase

Update `ROADMAP.md` to mark the phase as "fully specified" and update `VERIFIER_GUIDE.md` with the new MUST checks (same style as the existing ones).

Keep the overall spirit: long-term deliberate practice, heavy reuse, external study, and regular shipping of real (if simplified) artifacts.
