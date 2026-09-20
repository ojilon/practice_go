# Core Note — Structs

> Prerequisite for: `phase_01_core/02_structs`. Used by: `math2` (Vec/Mat), `list` (Node), `logger`, `cache`, `http_api` (Todo), everything with state.

## What a struct is
- A **composite value type**: groups named fields under one type. `type User struct { Name string; Age int }`.
- Literals: `User{Name: "a", Age: 3}` (field names — prefer this) vs `User{"a", 3}` (positional, brittle).
- `==` compares field-by-field **only if all fields are comparable** (no slices/maps/funcs inside). Copying a struct copies all fields (shallow).

## Embedding (not inheritance)
- `type Admin struct { User; Level int }` — `Admin.Name` is **promoted**. No polymorphism; just field/method promotion.
- Used to compose behavior cheaply (e.g. options structs, request contexts later).

## Tags + constructors
- Tags are metadata strings for reflection: ``Name string `json:"name"` ``. Meaningless until some code reads them via `reflect` (see `36_reflection`, `23_jsonlite`).
- **Constructors** (`func NewUser(...) (*User, error)`) enforce invariants at creation: validate, set defaults, return error instead of a half-built value. Return pointer when the value is large or must be mutated/shared; value when small + immutable.

## Building-block role
- Every package downstream centers on 1–3 structs: `Stack`, `Ring`, `Logger`, `Bus`, `LRU`, `Heap`, `Graph`, `Server`.
