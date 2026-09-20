# myfmt — Your Own Mini fmt — Guide (`myfmt`)

> Folder: `phase_01_core/00_myfmt`  
> Package name: `myfmt`  
> Run via root `main.go` (the verifier always ignores `main.go`)  
> Prerequisite: read `core_notes/01_variables_functions.md` first.

This is the **first real exercise**. Completing it earns you the right to use the real `fmt` package for debug printing everywhere else in the dojo.

---

## Why this exists

Almost every program needs to turn values into text. The real `fmt` package does a lot of sophisticated work. Here you rebuild a tiny, practice-grade subset so you understand the basic mechanics:

- turning integers into decimal strings (and back)
- writing text to an `io.Writer`
- a very limited form of formatted printing

You are **not** allowed to import `fmt`, `strconv`, `strings`, or `bytes` in your implementation files. Demos may import the real `fmt` only to compare your output side-by-side.

---

## What you will build (simple → complex)

1. **Itoa / Atoi lite**  
   Convert `int` ↔ decimal string. Handle the sign. Return a clear error on invalid input for Atoi.

2. **Print / Println / Fprint lite**  
   Accept any `io.Writer` (default to `os.Stdout`). Separate arguments with spaces the way `fmt` does. Println adds a final newline. Prefer a single `Write` call at the end.

3. **Sprintf lite**  
   Support only these verbs: `%s`, `%d`, `%v`, `%%`.  
   Unknown verbs should be left as-is (document this limitation).  
   Build the result with `[]byte` appends; avoid string `+` inside a loop.

4. **Discipline**  
   Keep the implementation small and readable. Edge cases matter more than fancy features.

Full task list and required symbols are in `01_EXERCISES.md`.

---

## How to run while developing

Temporarily change root `main.go` to something like:

```go
package main

import (
	"fmt"
	"practice_go/phase_01_core/00_myfmt"
)

func main() {
	fmt.Println(myfmt.Sprintf("%s=%d", "n", 42)) // should print n=42
}
```

Then from the repo root:

```bash
go run .
go vet ./phase_01_core/00_myfmt/...
```

When you are finished with the folder, you can leave `main.go` however you like — the verifier never looks at it.

---

## Done when

- Every symbol required by `01_EXERCISES.md` exists and behaves correctly on the stated edge cases.
- `Sprintf("%s=%d", "n", 42)` produces `"n=42"`.
- All answer boxes in `02_QUESTIONS.md` are filled.
- `go vet ./phase_01_core/00_myfmt/...` is clean.
- You understand why the real `fmt` package is more complex than what you just built.

After this folder you may freely use the real `fmt` for debug output.  
However, any formatting logic that belongs inside later packages you build yourself (especially the logger) must still be *your* code — see `BUILD_CHAIN.md`.
