package option

import "slices"

// TypeFields is the token that designates the fields to be returned
const TypeFields = "FIELDS"

// FieldsOption is a query option that limits which fields are included in a dataset
type FieldsOption []string

// Fields returns a query option that will limit the query results to a certain set of fields
func Fields(fields ...string) Option {
	// Copy the caller's slice. A variadic spread (`Fields(names...)`) passes the
	// caller's own backing array, so storing it directly would let later writes
	// to `names` silently rewrite this option.
	return FieldsOption(slices.Clone(fields))
}

// OptionType identifies this object as a query option
func (option FieldsOption) OptionType() string {
	return TypeFields
}

// Fields returns the names of the fields to include in a dataset
func (option FieldsOption) Fields() []string {
	// Copy on the way out too, so an adapter cannot corrupt the option it was handed
	return slices.Clone(option)
}
