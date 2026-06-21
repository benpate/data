# Option

Query options for the [data](../README.md) library. Each option is a small value type implementing the `Option` interface; together they describe optional query behavior (sorting, row limits, field projection, case sensitivity) without binding to any specific database.

These types **only encapsulate intent**. They contain no query logic — it is each database adapter's job to inspect the options it receives and translate them into its own query engine. See the adapter packages ([data-mongo](https://github.com/benpate/data-mongo), [data-slice](https://github.com/benpate/data-slice)) for how each is interpreted.

## What matters here

- **`FirstRow()` means "return only the first matching row" — it is NOT a pagination offset.** `FirstRowOption` is an empty struct and carries no row number; adapters translate it to a limit of one row. The name misleads: use `MaxRows(n)` to cap results, and do not expect `FirstRow` to skip to an arbitrary row.
- **Each option carries a `TypeXxx` string token returned by `OptionType()`.** Adapters switch on the concrete type (not the token) to apply each option; the token exists mainly for debugging. When adding a new option, follow the existing trio: a `TypeXxx` constant, an `XxxOption` value type, and a lowercase constructor returning `Option`.
- **An adapter may silently ignore an option it doesn't support.** There is no central registry forcing every backend to honor every option, so a query relying on `CaseSensitive` or `Fields` can behave differently across backends. Confirm support in the specific adapter.
