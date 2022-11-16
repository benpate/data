package option

// TypeFields is the token that designates the fields to be returned
const TypeFields = "MAXROWS"

// FieldsConfig is a query option that limits the number of rows to be included in a dataset
type FieldsConfig []string

// Fields returns a query option that will limit the query results to a certain number of rows
func Fields(fields []string) Option {
	return FieldsConfig(fields)
}

// OptionType identifies this record as a query option
func (fieldsConfig FieldsConfig) OptionType() string {
	return TypeFields
}

// Fields returns the maximum number of rows to include in a dataset
func (fieldsConfig FieldsConfig) Fields() []string {
	return []string(fieldsConfig)
}
