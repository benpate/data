package data

import (
	"github.com/benpate/data/expression"
	"github.com/benpate/data/option"
)

// Collection represents a single database collection (or table) that is opened to support a single transactional request, and then closed
// when this transaction is complete
type Collection interface {
	List(filter expression.Expression, options ...option.Option) (Iterator, error)
	Load(filter expression.Expression, target Object) error
	Save(object Object, note string) error
	Delete(object Object, note string) error
}
