package option

// TypeCaseSensitive is the token that designates case sensitivity of search results
const TypeCaseSensitive = "CASESENSITIVE"

// CaseSensitiveOption is a query option that designates case sensitivity of search results.
type CaseSensitiveOption bool

// CaseSensitive returns a query option that will designate the case sensitivity of search results
func CaseSensitive(caseSensitive bool) Option {
	return CaseSensitiveOption(caseSensitive)
}

// OptionType identifies this object as a query option
func (option CaseSensitiveOption) OptionType() string {
	return TypeCaseSensitive
}

// CaseSensitive returns TRUE if the results should consider letter case
func (option CaseSensitiveOption) CaseSensitive() bool {
	return bool(option)
}
