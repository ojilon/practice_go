# Core Notes — Index (the building blocks)

These notes are your **reference sheets** for the Go core. Everything downstream
(`phase_02_packages` → `phase_05_apps`) is built from these blocks — when an
exercise says "uses structs + interfaces + errors", this is what it means.

Read the note **before** starting its matching exercise folder.

| Note | Read before folder | One-line summary |
|---|---|---|
| `01_variables_functions.md` | `phase_01_core/01_variables_functions` (+ `00_myfmt`) | declarations, zero values, func signatures, variadics, closures |
| `02_structs.md` | `phase_01_core/02_structs` | composite types, embedding, tags, constructors |
| `03_methods.md` | `phase_01_core/03_methods` | receivers (value vs pointer), `String()`, chaining |
| `04_slices_arrays_strings.md` | `phase_01_core/04_slices_arrays_strings` | len/cap, append growth, backing arrays, bytes vs runes |
| `05_maps.md` | `phase_01_core/05_maps` | hashing view, comma-ok, iteration order, key rules |
| `06_interfaces.md` | `phase_01_core/06_interfaces` | implicit satisfaction, type switch, `io.Writer`, nil traps |
| `07_errors.md` | `phase_01_core/07_errors` | sentinel/wrap (`%w`), `Is`/`As`, custom types, panic rule |
| `08_pointers.md` | `phase_01_core/08_pointers` | `&`/`*`, mutation, nil, escape-to-heap |
| `09_goroutines.md` | `phase_01_core/09_goroutines` | `go`, `WaitGroup`, `Mutex`, race detector |
| `10_channels.md` | `phase_01_core/10_channels` | buffered/unbuffered, close, `select`, pipelines, fan-in/out |
| `11_defer_panic.md` | `phase_01_core/11_defer_panic_modules` | LIFO defers, recover, control flow |
| `12_packages_modules.md` | `phase_01_core/11_defer_panic_modules` | exports, packages, `go.mod`, one-module rule of this dojo |

Reuse map (where each block gets used later) is in `BUILD_CHAIN.md`.
The matching hands-on work is in `phase_01_core/00_myfmt` + `01_…11_`.
