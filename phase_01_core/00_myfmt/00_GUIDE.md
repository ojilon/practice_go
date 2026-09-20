# myfmt — Your Own Mini fmt — Guide (`myfmt`)

> Folder: `phase_01_core/00_myfmt` · Package: `myfmt` · Run via root `main.go` (verifier ignores `main.go`).
> Prerequisite reading: `core_notes/01_variables_functions.md`. Do this FIRST — it earns you the right to use `fmt` for debug everywhere else.

## Why this exists
`fmt` is the one stdlib package allowed for debug printing — but only because you prove here you understand what it does. Build a tiny, practice-grade subset from core only (no `fmt`/`strconv`/`strings` imports in impl files; demos may use `fmt` to compare outputs).

## Steps (simple → complex)
1. `Itoa`/`Atoi`-lite: int↔string with sign + error on bad input (see `01_EXERCISES.md` E1).
2. `Print`/`Println`-lite writing to any `io.Writer` (default: `os.Stdout`), space-separation like `fmt`.
3. `Sprintf`-lite supporting `%s %d %v %%` only. Unknown verbs → keep literally (e.g. `%q` stays as-is) + document the limitation.
4. Concat/bytes discipline: build with `[]byte` append, single final `Write` (no `+` in a loop).

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/00_myfmt")
func main() { fmt.Println(myfmt.Sprintf("%s=%d", "n", 42)) }
```
```powershell
go run .
go vet ./phase_01_core/00_myfmt/...
```

## Done when
- All `01_EXERCISES.md` symbols exist; `Sprintf("%s=%d", "n", 42)` → `"n=42"`.
- `02_QUESTIONS.md` boxes filled.
- Downstream you may use `fmt` for debug — but `18_logger` formatting logic must be YOURS (see `BUILD_CHAIN.md`).
