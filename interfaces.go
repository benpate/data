package data

import (
	"context"

	"github.com/benpate/derp"
)

// Datastore is an abstract representation of a database and its connection information.
type Datastore interface {
	Session(context.Context) Session
}

// Session represents a single database session, that is opened to support a single transactional request, and then closed
// when this transaction is complete
type Session interface {
	Load(collection string, filter Expression, target Object) *derp.Error
	Save(collection string, object Object, note string) *derp.Error
	Delete(collection string, object Object, note string) *derp.Error
	Close()
}

// Object interface defines all of the methods that a Domain Object must provide to Presto
type Object interface {

	// ID returns the primary key of the object
	ID() string

	// IsNew returns TRUE if the object has not yet been saved to the database
	IsNew() bool

	// SetCreated stamps the CreateDate and UpdateDate of the object, and makes a note
	SetCreated(comment string)

	// SetUpdated stamps the UpdateDate of the object, and makes a note
	SetUpdated(comment string)

	// SetDeleted marks the object virtually "deleted", and makes a note
	SetDeleted(comment string)
}
