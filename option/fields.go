package option

// TypeFields is the token that designates the fields to be returned
const TypeFields = "FIELDS"

// FieldsOption is a query option that limits which fields are included in a dataset
type FieldsOption []string

// Fields returns a query option that will limit the query results to a certain set of fields
func Fields(fields ...string) Option {
	return FieldsOption(fields)
}

// OptionType identifies this object as a query option
func (option FieldsOption) OptionType() string {
	return TypeFields
}

// Fields returns the names of the fields to include in a dataset
func (option FieldsOption) Fields() []string {
	return []string(option)
}
