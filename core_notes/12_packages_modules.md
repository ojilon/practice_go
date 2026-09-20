# Core Note — Packages & Modules

> Prerequisite for: `phase_01_core/11_defer_panic_modules`. Rule for the whole dojo.

## Packages
- One directory = one package. Exported (usable elsewhere) = Capitalized: `func New()`, `type Logger`. Unexported = lowercase, package-private.
- Import path in this dojo: `practice_go/<dir-path>`, e.g. `import "practice_go/phase_02_packages/14_stack"`. Package clause drops the `NN_` prefix and underscores (`14_stack` → `package stack`).
- Multi-file package: `ops_a.go` + `ops_b.go` with the same `package x` clause — one logical unit, split for readability.

## Modules (this dojo = ONE module)
- Root `go.mod` declares `module practice_go`. Do NOT create nested `go.mod` files — subfolders are packages, not modules.
- `go run .` (root scratch), `go vet ./...` (everything), `go vet ./<folder>/...` (one exercise).

## Design rules for your packages (used by verifier)
1. Accept interfaces, return structs (`New(w io.Writer) *Logger`, `Parse(r io.Reader)`).
2. Constructors validate (`New(cap int) (*Ring, error)` rejects `cap <= 0`).
3. Library code doesn't `fmt.Print` — it returns values/errors; printing happens in `Demo()`/callers.
4. One-line `//` doc comments on every exported symbol.

## Building-block role
- This is WHY the dojo works: each folder is a real importable package, so `37_http_api` can `import "practice_go/phase_02_packages/18_logger"` and run on YOUR code. See `BUILD_CHAIN.md` for the reuse map.
