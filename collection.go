package data

import (
	"github.com/benpate/data/option"
	"github.com/benpate/exp"
)

// Collection represents a single database collection (or table) that is opened to support a single transactional request, and then closed
// when this transaction is complete
type Collection interface {
	List(filter exp.Expression, options ...option.Option) (Iterator, error) // Query the collection
	Load(filter exp.Expression, target Object) error                        // Load an object
	Save(object Object, note string) error                                  // Insert/Update an object
	Delete(object Object, note string) error                                // Virtual delete of an object
	HardDelete(Object) error                                                // Physical delete of an object
}
