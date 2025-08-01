package option

// TypeFields is the token that designates the fields to be returned
const TypeFields = "FIELDS"

// FieldsOption is a query option that limits the number of rows to be included in a dataset
type FieldsOption []string

// Fields returns a query option that will limit the query results to a certain number of rows
func Fields(fields ...string) Option {
	return FieldsOption(fields)
}

// OptionType identifies this object as a query option
func (option FieldsOption) OptionType() string {
	return TypeFields
}

// Fields returns the maximum number of rows to include in a dataset
func (option FieldsOption) Fields() []string {
	return []string(option)
}
