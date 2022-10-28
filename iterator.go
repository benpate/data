package data

// Iterator interface allows callers to iterator over a large number of items in an array/slice
type Iterator interface {
	Next(any) bool
	Error() error
	Count() int
	Close() error
}
