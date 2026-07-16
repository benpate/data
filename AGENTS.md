# Data — Notes for AI Agents

This package contains **interfaces and data only** — there is no database logic. CRUD behavior lives entirely in the adapter packages ([data-mongo](https://github.com/benpate/data-mongo), [data-mock](https://github.com/benpate/data-mock), [data-slice](https://github.com/benpate/data-slice)). When changing this package, remember that every adapter must continue to satisfy the interfaces.

- **`option.FirstRow()` means "return only the first matching row" — it is not a pagination offset.** Despite the name, it carries no row number (`FirstRowOption` is an empty struct). Adapters translate it to a limit of one row (Mongo `SetLimit(1)`, slice `value[:1]`). Reaching for it to "skip to row N" will silently do the wrong thing.

- **Options only encapsulate intent; adapters decide what they mean.** An adapter is free to ignore an option it doesn't support, so a query that relies on `CaseSensitive` or `Fields` may behave differently across backends. Verify support in the specific adapter, not here.

- **`Delete` is a soft delete; `HardDelete` is permanent.** `Delete(object, note)` stamps the journal's `DeleteDate` and leaves the record in place; `HardDelete(criteria)` removes matching rows outright and takes no object or note.

- **The journal is the source of truth for object lifecycle.** Embedding `journal.Journal` satisfies all of `Object` except `ID()`. `SetUpdated`/`SetDeleted` bump `Revision` (the ETag); `SetCreated` does not. The `Revision` field serializes as `"signature"` for backward compatibility — don't rename the tag.
