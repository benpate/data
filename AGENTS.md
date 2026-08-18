# Data — Notes for AI Agents

This package contains **interfaces and data only** — there is no database logic. CRUD behavior lives entirely in the adapter packages ([data-mongo](https://github.com/benpate/data-mongo), [data-mock](https://github.com/benpate/data-mock), [data-slice](https://github.com/benpate/data-slice)). When changing this package, remember that every adapter must continue to satisfy the interfaces.

- **`option.FirstRow()` means "return only the first matching row" — it is not a pagination offset.** Despite the name, it carries no row number (`FirstRowOption` is an empty struct). Adapters translate it to a limit of one row (Mongo `SetLimit(1)`, slice `value[:1]`). Reaching for it to "skip to row N" will silently do the wrong thing.

- **Options only encapsulate intent; adapters decide what they mean.** An adapter is free to ignore an option it doesn't support, so a query that relies on `CaseSensitive` or `Fields` may behave differently across backends. Verify support in the specific adapter, not here.

- **`Delete` is a soft delete; `HardDelete` is permanent.** `Delete(object, note)` stamps the journal's `DeleteDate` and leaves the record in place; `HardDelete(criteria)` removes matching rows outright and takes no object or note.

- **The journal is the source of truth for object lifecycle.** Embedding `journal.Journal` satisfies all of `Object` except `ID()`. `SetUpdated`/`SetDeleted` bump `Revision` (the ETag); `SetCreated` does not. The `Revision` field serializes as `"signature"` for backward compatibility — don't rename the tag.

- **`Collection.Load` and `Collection.Query` take their arguments in OPPOSITE order.** `Load(criteria, target, options...)` is criteria-first; `Query(target, criteria, options...)` is target-first. Both accept a target and a criteria, so a transposed call still reads plausibly. `Load`'s target is a `data.Object`; `Query`'s is a bare `any` (it's usually a slice), so the compiler will not always catch the swap. Check the order at every call site.

- **Only POINTERS satisfy `data.Object`.** `Journal`'s accessors use value receivers but its `SetCreated`/`SetUpdated`/`SetDeleted` mutators use pointer receivers, so an embedding type satisfies `Object` only in its pointer form. `var _ data.Object = Person{}` does not compile; `&Person{}` does. `journal/contract_test.go` pins this — if you add a method to `Object`, that file stops compiling, which is the intended alarm.

- **`option.MaxRows(0)` means "no limit", not "no rows".** Negative limits are clamped to zero by both the `MaxRows()` constructor and the `MaxRows()` accessor, so a meaningless limit can never reach an adapter (mongodb reads a raw `-1` as "one row, then close the cursor"). Zero is also the zero value of `MaxRowsOption`, which makes an uninitialized option harmless. To ask for a single row, use `FirstRow()`.

- **Adapters must branch on `SortOption.IsDescending()`, never on the raw `Direction` field.** `Direction` is exported and unvalidated, so it can hold a typo or the empty string from a zero-value `SortOption`. `IsDescending()` returns TRUE only for the exact `SortDirectionDescending` token, so everything unrecognized sorts ASCENDING — matching data-mongo and data-mock. data-slice previously defaulted the other way, which is why the accessor exists.

- **`option.Fields` copies the caller's slice in both directions.** `Fields(names...)` clones on the way in and the `Fields()` accessor clones on the way out, so neither the caller nor an adapter can reach through and mutate a live option. Don't "optimize" either clone away — a variadic spread passes the caller's own backing array.
