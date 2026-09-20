# Core Note — Arrays, Slices, Strings & Runes

> Prerequisite for: `phase_01_core/04_slices_arrays_strings`. Used by: `stack`, `ringbuffer`, `strings_builder`, `tokenizer`, `parsers`, `scraper`.

## Arrays vs slices
- `[3]int` is a VALUE (copy on assign, fixed length). `[]int` is a 3-word HEADER `{ptr, len, cap}` over a shared backing array.
- `len` = visible elements, `cap` = backing capacity. `make([]int, 0, 2)` pre-sizes to avoid regrowth.

## Append / copy / slicing (the three operations that cause 90% of bugs)
- `append` may REALLOCATE: always use result (`s = append(s, x)`). Growth ~doubling; `cap` history is a QUESTIONS item.
- `copy(dst, src)` copies `min(len)` elements — use `Clone` for independence.
- `s[1:3]` SHARES the backing array: appending through one alias can clobber the other. Full slice expr `s[1:3:3]` caps capacity to prevent that.
- `nil` slice (`var s []int`) reads fine (`len==0`, `range` ok), `append` works; only explicit index writes panic.

## Strings: bytes vs runes
- `string` = read-only `[]byte`, UTF-8 encoded. Indexing gives BYTES; `range` decodes RUNES.
- `len("é") == 2` bytes, `1` rune. `[]rune(s)` for correct reversal/counting; `[]byte(s)` for I/O.
- Build strings with `strings.Builder` (or YOUR builder in `26_strings_builder`) — `+` in a loop is O(n²).

## Building-block role
- Slice-backed `Stack`/`Queue`/`Heap`/`Ring` are all `{buf []T, head/len}` + append/index discipline.
- Rune-correctness gates `Reverse`, `WordCount`, CSV quoting, `<title>` extraction.
