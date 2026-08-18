package option

// TypeSort is the token that designates a Sort order
const TypeSort = "SORT"

// SortDirectionAscending is the token that designates that records should be sorted lowest to highest
const SortDirectionAscending = "ASC"

// SortDirectionDescending is the token that designates that records should be sorted highest to lowest
const SortDirectionDescending = "DESC"

// SortOption identifies the field and direction to use when sorting a dataset
type SortOption struct {
	FieldName string // FieldName is the name of the field to sort on
	Direction string // Direction is the sort order, either SortDirectionAscending or SortDirectionDescending
}

// OptionType identifies this object as a query option
func (option SortOption) OptionType() string {
	return TypeSort
}

// IsDescending returns TRUE if this option sorts from highest to lowest.
// Any direction other than SortDirectionDescending sorts ASCENDING.
func (option SortOption) IsDescending() bool {
	// Adapters must branch on this, not on the raw Direction field. Direction is
	// exported and unvalidated, so it can hold a typo or the empty string (from a
	// zero-value SortOption); ascending is the safe default for anything unrecognized.
	return option.Direction == SortDirectionDescending
}

// IsAscending returns TRUE if this option sorts from lowest to highest
func (option SortOption) IsAscending() bool {
	return !option.IsDescending()
}

// SortAsc returns a query option that will sort the query results in ASCENDING order
func SortAsc(fieldName string) Option {
	return SortOption{
		FieldName: fieldName,
		Direction: SortDirectionAscending,
	}
}

// SortDesc returns a query option that will sort the query results in DESCENDING order
func SortDesc(fieldName string) Option {
	return SortOption{
		FieldName: fieldName,
		Direction: SortDirectionDescending,
	}
}
