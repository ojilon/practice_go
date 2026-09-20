# Core Note — Methods (value vs pointer receivers)

> Prerequisite for: `phase_01_core/03_methods`. Used by: every package API you will design.

## Receiver = the type the method belongs to
- `func (u User) Greet() string` — **value receiver**: works on a copy. Use for read-only ops on small values.
- `func (u *User) Birthday()` — **pointer receiver**: can mutate; avoids copying large structs.
- **Consistency rule:** if ANY method needs a pointer receiver, use pointer receivers for ALL methods of that type (keeps method sets predictable).

## Method sets (why it matters for interfaces)
- `User`'s method set = value-receiver methods only. `*User`'s method set = both.
- So a `*User` satisfies MORE interfaces than `User`. When you store things in interface-typed collections, be deliberate (pointer vs value).

## Two idioms you will reuse constantly
1. **`String() string`** satisfies `fmt.Stringer` — controls how your type prints everywhere (logger lines, TUI cells, debug).
2. **Fluent builder:** methods returning the receiver (`func (b *B) Add(x int) *B`) enable chaining: `b.Add(1).Add(2).Build()`.

## Building-block role
- Push/Pop/Enqueue/Set/Get/Publish — every downstream API is methods on a struct with a chosen receiver. Wrong choice = mutation silently lost (value receiver on Deposit) — the classic exam question in `02_QUESTIONS.md`.
