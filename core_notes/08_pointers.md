# Core Note — Pointers

> Prerequisite for: `phase_01_core/08_pointers`. Used by: linked structures (`list`, `tree`, `graph` nodes), mutating methods, constructors.

## Essentials
- `&x` = address, `*p` = dereference. `var p *int` is `nil` until set — dereferencing nil PANICS, so nil-guard public APIs.
- `new(T)` allocates zeroed `T`, returns `*T`. `&T{...}` same with literal.
- Pass `*T` when func MUST mutate caller state (`Inc`, `Swap`, `Deposit`) or value is big. Otherwise pass `T`.
- Returning `&local` is SAFE in Go (escapes to heap) — enables `NewX(...) *X` constructors.

## Pointer-shaped thinking
- Linked nodes ARE pointers: `type Node struct { Val int; Next *Node }` — `nil` means end. Walking = pointer chasing.
- Receivers (see methods note): pointer receiver = method mutates shared struct.
- Slices/maps already share backing storage — don't add `*[]T` unless you must replace the header itself.

## Building-block role
- `List/Tree/Graph` nodes, `LRU` elements (`*list.Element`), `Heap` sift swaps by index — pointer discipline decides correctness.
