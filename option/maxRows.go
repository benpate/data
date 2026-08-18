package option

// TypeMaxRows is the token that designates the maximum number of records to be returned
const TypeMaxRows = "MAXROWS"

// MaxRowsOption is a query option that limits the number of rows to be included in a dataset.
// A value of zero means "no limit", which makes the zero value safe to use.
type MaxRowsOption int64

// MaxRows returns a query option that will limit the query results to a certain number of rows
func MaxRows(maxRows int64) Option {
	return MaxRowsOption(clampMaxRows(maxRows))
}

// OptionType identifies this object as a query option
func (option MaxRowsOption) OptionType() string {
	return TypeMaxRows
}

// MaxRows returns the maximum number of rows to include in a dataset, where zero means "no limit"
func (option MaxRowsOption) MaxRows() int64 {
	// Clamp on the way out as well, so a value built by direct conversion
	// (MaxRowsOption(-1)) obeys the same invariant as one from MaxRows().
	return clampMaxRows(int64(option))
}

// clampMaxRows folds every meaningless row limit onto zero, the "no limit" value
func clampMaxRows(maxRows int64) int64 {

	// RULE: A negative limit has no meaning, and adapters disagree on it
	// (mongodb reads -1 as "one row, then close the cursor"), so it is not
	// passed through. Zero is the one value every adapter reads as "no limit".
	if maxRows < 0 {
		return 0
	}

	return maxRows
}
