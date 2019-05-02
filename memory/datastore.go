package memory

import (
	"context"
	"strings"

	"github.com/benpate/criteria"
	"github.com/benpate/data"
	"github.com/benpate/derp"
)

// DATABASE OBJECT
type Datastore map[string]Collection

type Collection map[string]data.Object

func (db *Datastore) Session(ctx context.Context) *Datastore {
	return db
}

func (db *Datastore) Load(collection string, filter criteria.Expression, target data.Object) *derp.Error {

	if collection, ok := (*db)[collection]; ok {

		for _, document := range collection {

			if person, ok := document.(*testPerson); ok {

				if filter.Match(*person) {

					switch target := target.(type) {
					case *testPerson:

						*target = *person
						return nil
					}
				}
			}
		}

		return derp.New(404, "Datastore.Load", "Document not found", filter)
	}

	return derp.New(404, "Datastore.Load", "Collection does not exist", collection)
}

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
