package mockdb

import (
	"sort"
	"strings"

	"github.com/benpate/data"
	"github.com/benpate/data/expression"
	"github.com/benpate/data/option"
	"github.com/benpate/derp"
)

// Session is a mock database session
type Session struct {
	Datastore *Datastore
}

// List retrieves a group of records as an Iterator.
func (session Session) List(collection string, criteria expression.Expression, options ...option.Option) (data.Iterator, *derp.Error) {

	result := []data.Object{}

	if session.hasCollection(collection) == false {
		return NewIterator(result), derp.New(404, "mockdb.Load", "Collection does not exist", collection)
	}

	c := session.getCollection(collection)

	for _, document := range c {

		if (criteria == nil) || (criteria.Match(MatcherFunc(document))) {
			result = append(result, document)
		}
	}

	iterator := NewIterator(result, options...)

	sort.Sort(iterator)

	return iterator, nil

}

// Load retrieves a single record from the mock collection.
func (session Session) Load(collection string, criteria expression.Expression, target data.Object) *derp.Error {

	if session.hasCollection(collection) == false {
		return derp.New(404, "mockdb.Load", "Collection does not exist", collection)
	}

	c := session.getCollection(collection)

	for _, document := range c {

		if (criteria == nil) || (criteria.Match(MatcherFunc(document))) {
			return populateInterface(document, target)
		}
	}

	return derp.New(404, "mockdb.Load", "Document not found", criteria)
}

// Save adds/inserts a new record into the mock database
func (session Session) Save(collection string, object data.Object, comment string) *derp.Error {

	if strings.HasPrefix(comment, "ERROR") {
		return derp.New(500, "mockdb.Save", "Synthetic Error", comment)
	}

	c := session.getCollection(collection)

	object.SetUpdated(comment)

	if object.IsNew() {
		object.SetCreated(comment)
		(*session.Datastore)[collection] = append(c, object)
		return nil
	}

	if index := c.FindByObjectID(object.ID()); index >= 0 {
		c[index] = object
		return nil
	}

	return derp.New(500, "mockdb.Save", "Object Not Found", "attempted to update object, but it does not exist in the datastore", object)
}

// Delete PERMANENTLY removes a record from the mock database.
func (session Session) Delete(collection string, object data.Object, comment string) *derp.Error {

	if strings.HasPrefix(comment, "ERROR") {
		return derp.New(500, "mockdb.Delete", "Synthetic Error", comment)
	}

	c := session.getCollection(collection)

	if index := c.FindByObjectID(object.ID()); index >= 0 {
		(*session.Datastore)[collection] = append(c[:index], c[index+1:]...)
	}

	return nil
}

// Close cleans up any remaining data created by the mock session.
func (session Session) Close() {}

func (session Session) hasCollection(collection string) bool {

	_, ok := (*session.Datastore)[collection]

	return ok
}

// getCollection loads (and creates, if necessary) the named collection in this datastore
func (session Session) getCollection(collection string) Collection {

	if session.hasCollection(collection) == false {
		(*session.Datastore)[collection] = NewCollection()
	}

	return (*session.Datastore)[collection]
}
