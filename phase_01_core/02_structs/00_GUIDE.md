# Structs — Guide (`structsx`)

> Folder: `phase_01_core/02_structs` · Package: `structsx` · Run via root `main.go` (verifier ignores `main.go`).

## What you will learn
- Step 1 — type User struct, composite literals, field access.
- Step 2 — embedding (Address in User), promoted fields.
- Step 3 — struct tags (`json:"name"`), constructors NewUser returning *User with validation.
- Step 4 — comparable vs non-comparable structs, == rules.

## Steps (simple → complex)
Do in order. Create `*.go` files with `package structsx` in THIS folder.

1. Read the steps below, then implement the matching item in `01_EXERCISES.md`.
2. Wire a temporary `main.go` import: `import "practice_go/phase_01_core/02_structs"` (package name `structsx`), call your funcs, `go run .`.
3. `go vet ./phase_01_core/02_structs/...` must be clean before moving on.

### Details
- Step 1 — type User struct, composite literals, field access.
- Step 2 — embedding (Address in User), promoted fields.
- Step 3 — struct tags (`json:"name"`), constructors NewUser returning *User with validation.
- Step 4 — comparable vs non-comparable structs, == rules.

## How to run (example — adapt)
```go
// root main.go (TEMPORARY, verifier ignores it)
package main
import ("fmt"; "practice_go/phase_01_core/02_structs")
func main() { fmt.Println(structsx.Demo()) }
```
```powershell
go run .
go vet ./phase_01_core/02_structs/...
```

## Done when
- All `01_EXERCISES.md` symbols exist and behave per edge cases.
- `02_QUESTIONS.md` boxes filled (if Type B).
