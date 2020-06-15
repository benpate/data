package data

import (
	"github.com/benpate/data/expression"
	"github.com/benpate/data/option"
	"github.com/benpate/derp"
)

// Collection represents a single database collection (or table) that is opened to support a single transactional request, and then closed
// when this transaction is complete
type Collection interface {
	List(filter expression.Expression, options ...option.Option) (Iterator, *derp.Error)
	Load(filter expression.Expression, target Object) *derp.Error
	Save(object Object, note string) *derp.Error
	Delete(object Object, note string) *derp.Error
}
