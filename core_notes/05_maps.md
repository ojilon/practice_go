# Core Note — Maps

> Prerequisite for: `phase_01_core/05_maps`. Used by: `set`, `cache`, `config`, `ini`, `graph`, `jsonlite` (objects), `parsers` (summaries).

## Essentials
- `m := make(map[string]int)` or literal. Read missing key → zero value (can't distinguish absent vs stored-zero without comma-ok).
- **Comma-ok:** `v, ok := m[k]` — THE lookup idiom. `ok==false` means absent.
- `delete(m, k)` is safe on absent keys. Assignment to `nil` map PANICS; reading nil map is fine.
- **Iteration order is RANDOMIZED** per run (deliberate). For deterministic output: collect keys, `sort`, then iterate.
- Keys must be COMPARABLE (`==`/`!=`): no slices/maps/funcs as keys. Struct keys OK if all fields comparable.
- Maps are REFERENCES: passing to a func shares it; concurrent read+write without sync = fatal (needs `Mutex` or channels — see `09/10`).

## Patterns you will reuse
- Frequency table: `freq[w]++` (zero value does the work).
- Nested: `map[string]map[string]int` (grades), adjacency: `map[string][]string` (graph).
- Set-as-map: `map[T]struct{}` — zero-size value, presence = membership (see `16_set`).

## Building-block role
- `Set`, `LRU index`, `Config values`, `INI sections`, `Graph adjacency`, JSON objects — all maps with different value types.
