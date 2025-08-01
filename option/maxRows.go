package option

// TypeMaxRows is the token that designates the maximum number of records to be returned
const TypeMaxRows = "MAXROWS"

// MaxRowsOption is a query option that limits the number of rows to be included in a dataset
type MaxRowsOption int64

// MaxRows returns a query option that will limit the query results to a certain number of rows
func MaxRows(maxRows int64) Option {
	return MaxRowsOption(maxRows)
}

// OptionType identifies this object as a query option
func (option MaxRowsOption) OptionType() string {
	return TypeMaxRows
}

// MaxRows returns the maximum number of rows to include in a dataset
func (option MaxRowsOption) MaxRows() int64 {
	return int64(option)
}
