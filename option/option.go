// Package option provides an interface for query options that can be used to modify
// database queries.
package option

// Option is an interface that wraps all of the query options that can be used to modify
// a database query.
type Option interface {
	OptionType() string // OptionType identifies the type of query option, and may be used for debugging purposes.
}
