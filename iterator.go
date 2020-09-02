package data

import "github.com/benpate/derp"

// Iterator interface allows callers to iterator over a large number of items in an array/slice
type Iterator interface {
	Next(Object) bool
	Count() int
	Close() *derp.Error
}
