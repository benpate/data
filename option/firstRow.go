package option

// TypeFirstRow is the token that designates the "row number" to begin results.
const TypeFirstRow = "FIRSTROW"

// FirstRowOption is a query option that sets the first row to be returned in a dataset
type FirstRowOption struct{}

// FirstRow returns a query option that will limit the query results to a certain number of rows
func FirstRow() Option {
	return FirstRowOption{}
}

// OptionType identifies this object as a query option
func (option FirstRowOption) OptionType() string {
	return TypeFirstRow
}
