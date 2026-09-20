# tokenizer — Lexer & Mini-Regex — Guide (`tokenizer`)

> Folder: `phase_03_dsa/32_tokenizer` · Package: `tokenizer` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — Token{Kind,Lit}; kinds: Number/Ident/Plus/Star/LParen/...
- Step 2 — Lex(s string) []Token handling spaces, strings with escapes, numbers.
- Step 3 — mini-match: Match(pat like `a*b?`) stretch; CSV-line lexer variant.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package tokenizer` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_03_dsa/32_tokenizer"` (package name `tokenizer`), call your funcs, `go run .`.
3. `go vet ./phase_03_dsa/32_tokenizer/...` must be clean before moving on.

### Details
- Step 1 — Token{Kind,Lit}; kinds: Number/Ident/Plus/Star/LParen/...
- Step 2 — Lex(s string) []Token handling spaces, strings with escapes, numbers.
- Step 3 — mini-match: Match(pat like `a*b?`) stretch; CSV-line lexer variant.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_03_dsa/32_tokenizer")
func main() { fmt.Println(tokenizer.Demo()) }
```
```powershell
go run .
go vet ./phase_03_dsa/32_tokenizer/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
