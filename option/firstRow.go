package option

// TypeFirstRow is the token that designates the "first row only" query option.
const TypeFirstRow = "FIRSTROW"

// FirstRowOption is a query option that returns only the first matching row of a dataset.
type FirstRowOption struct{}

// FirstRow returns a query option that limits the results to only the first matching row.
func FirstRow() Option {
	return FirstRowOption{}
}

// OptionType identifies this object as a query option
func (option FirstRowOption) OptionType() string {
	return TypeFirstRow
}
