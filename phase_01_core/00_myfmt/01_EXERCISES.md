# myfmt — Your Own Mini fmt — Exercises

> Work in THIS folder, `package myfmt`. Run via root `main.go` (temporary import). Verifier checks THESE files, not `main.go`.

## Must implement (checked by verifier)
- **E1 ints** — `e01.go`: `func Itoa(n int) string` (sign-aware, no `strconv`), `func Atoi(s string) (int, error)` (`ErrInvalid` on empty/bad chars, sign-aware).
- **E2 print** — `e02.go`: `func Fprint(w io.Writer, args ...any) (int, error)`, `func Print(args ...any) (int, error)` (stdout), `func Println(args ...any) (int, error)` (spaces + trailing newline). One `Write` call per invocation; return bytes written.
- **E3 sprintf** — `e03.go`: `func Sprintf(format string, args ...any) string` supporting `%s %d %v %%` (+ `%t` stretch). Missing arg → `%!verb(MISSING)`; extra args ignored; unknown verbs kept literally.
- **E4 widths (stretch, credited)** — `%5d` / `%-5s` padding.

## Banned in impl (verifier greps)
- `fmt`, `strconv`, `strings`, `bytes`, `encoding/*` inside `e01.go`/`e02.go`/`e03.go`. (`doc.go`/demo funcs may import `fmt` to compare.)

## Edge cases (will be spot-checked)
- `Atoi("")`, `Atoi("12x")`, `Atoi("-")` → error, not panic, not partial number.
- `Itoa(0)` → `"0"`; round-trip `Atoi(Itoa(n)) == n` for negative/large values.
- `Sprintf("100%%")` → `"100%"`; `Sprintf("%d")` (no args) contains `MISSING`.

## Suggested files
- `e01.go`, `e02.go`, `e03.go`, package clause MUST be `package myfmt`.

## Stretch (optional, credited)
- `Demo()` printing side-by-side `fmt` vs `myfmt` outputs for manual `go run .` check.
