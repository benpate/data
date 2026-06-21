package data

import (
	"context"

	"github.com/benpate/data/option"
	"github.com/benpate/exp"
)

// Collection represents a single database collection (or table) that is opened to support a single transactional request, and then closed
// when this transaction is complete
type Collection interface {

	// Context returns the context.Context that scopes this collection's transaction
	Context() context.Context

	// Count returns the number of records that match the provided criteria
	Count(criteria exp.Expression, options ...option.Option) (int64, error)

	// Query populates the target (typically a slice) with all records that match the criteria
	Query(target any, criteria exp.Expression, options ...option.Option) error

	// Iterator returns an Iterator over all records that match the criteria
	Iterator(criteria exp.Expression, options ...option.Option) (Iterator, error)

	// Load populates the target with the first record that matches the criteria
	Load(criteria exp.Expression, target Object, options ...option.Option) error

	// Save inserts or updates the object, recording the note in its journal
	Save(object Object, note string) error

	// Delete virtually ("soft") deletes the object, recording the note in its journal
	Delete(object Object, note string) error

	// HardDelete permanently removes all records that match the criteria
	HardDelete(criteria exp.Expression) error
}
