package data

import (
	"github.com/benpate/data/expression"
	"github.com/benpate/data/option"
	"github.com/benpate/derp"
)

// Session represents a single database session, that is opened to support a single transactional request, and then closed
// when this transaction is complete
type Session interface {
	List(collection string, filter expression.Expression, options ...option.Option) (Iterator, *derp.Error)
	Load(collection string, filter expression.Expression, target Object) *derp.Error
	Save(collection string, object Object, note string) *derp.Error
	Delete(collection string, object Object, note string) *derp.Error
	Close()
}
