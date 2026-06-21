package data

// Iterator interface allows callers to iterate over a large number of items in a data structure
type Iterator interface {

	// Next populates the provided target with the next item, returning FALSE when the iterator is exhausted
	Next(any) bool

	// Error returns the first error (if any) encountered while iterating
	Error() error

	// Count returns the number of records contained by this iterator
	Count() int

	// Close releases any resources held by the iterator
	Close() error
}
