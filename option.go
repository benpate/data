package data

// OptionTypeStartRow is the token that designates the "row number" to begin results.
const OptionTypeStartRow = "STARTROW"

// OptionTypeMaxRows is the token that designates the maximum number of records to be returned
const OptionTypeMaxRows = "MAXROWS"

// OptionTypeSort is the token that designates a Sort order
const OptionTypeSort = "SORT"

// OptionSortDirectionAscending is the token that designates that records should be sorted lowest to highest
const OptionSortDirectionAscending = "ASC"

// OptionSortDirectionDescending is the token that designates that records should be sorted highest to lowest
const OptionSortDirectionDescending = "DESC"

// Option is a value that modifies a READ query result
type Option struct {
	Type  string
	Name  string
	Value interface{}
}

// OptionStartRow returns a query option that will limit the query results to a certain number of rows
func OptionStartRow(startRow int64) Option {
	return Option{
		Type:  OptionTypeStartRow,
		Value: startRow,
	}
}

// OptionMaxRows returns a query option that will limit the query results to a certain number of rows
func OptionMaxRows(maxRows int64) Option {
	return Option{
		Type:  OptionTypeMaxRows,
		Value: maxRows,
	}
}

// OptionSortAsc returns a query option that will sort the query results in ASCENDING order
func OptionSortAsc(field string) Option {
	return Option{
		Type:  OptionTypeSort,
		Name:  field,
		Value: OptionSortDirectionAscending,
	}
}

// OptionSortDesc returns a query option that will sort the query results in DESCENDING order
func OptionSortDesc(field string) Option {
	return Option{
		Type:  OptionTypeSort,
		Name:  field,
		Value: OptionSortDirectionDescending,
	}
}
