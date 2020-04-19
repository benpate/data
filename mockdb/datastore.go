package mockdb

import (
	"context"
	"sort"
	"strings"

	"github.com/benpate/data"
	"github.com/benpate/data/expression"
	"github.com/benpate/data/option"
	"github.com/benpate/derp"
)

// Datastore is a mock database
type Datastore map[string]Collection

// Collection is a mock table
type Collection map[string]data.Object

// New returns a fully initialized Database object
func New() data.Datastore {
	return &Datastore{}
}

// Session returns a session that can be used as a mock database.
func (db *Datastore) Session(ctx context.Context) data.Session {
	return db
}

// List retrieves a group of records as an Iterator.
func (db *Datastore) List(collection string, criteria expression.Expression, options ...option.Option) (data.Iterator, *derp.Error) {

	result := []data.Object{}

	if collection, ok := (*db)[collection]; ok {

		for _, document := range collection {

			if (criteria == nil) || (criteria.Match(MatcherFunc(document))) {
				result = append(result, document)
			}
		}

		iterator := NewIterator(result, options...)

		sort.Sort(iterator)

		return iterator, nil

	}

	return NewIterator(result), derp.New(404, "Datastore.Load", "Collection does not exist", collection)
}

// Load retrieves a single record from the mock collection.
func (db *Datastore) Load(collection string, criteria expression.Expression, target data.Object) *derp.Error {

	if collection, ok := (*db)[collection]; ok {

		for _, document := range collection {

			if (criteria == nil) || (criteria.Match(MatcherFunc(document))) {
				return populateInterface(document, target)
			}
		}

		return derp.New(404, "Datastore.Load", "Document not found", criteria)
	}

	return derp.New(404, "Datastore.Load", "Collection does not exist", collection)
}

// Save adds/inserts a new record into the mock database
func (db *Datastore) Save(collection string, object data.Object, comment string) *derp.Error {

	if strings.HasPrefix(comment, "ERROR") {
		return derp.New(500, "Datastore.Save", "Synthetic Error", comment)
	}

	if _, ok := (*db)[collection]; !ok {
		(*db)[collection] = Collection{}
	}

	if object.IsNew() {
		object.SetCreated(comment)
	}
	object.SetUpdated(comment)
	(*db)[collection][object.ID()] = object

	return nil
}

// Delete PERMANENTLY removes a record from the mock database.
func (db *Datastore) Delete(collection string, object data.Object, comment string) *derp.Error {

	if strings.HasPrefix(comment, "ERROR") {
		return derp.New(500, "Datastore.Delete", "Synthetic Error", comment)
	}

	if _, ok := (*db)[collection]; !ok {
		(*db)[collection] = Collection{}
	}

	delete((*db)[collection], object.ID())

	return nil
}

// Close cleans up any remaining data created by the mock session.
func (db *Datastore) Close() {}
