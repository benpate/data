package data

import (
	"context"
)

// Datastore is an abstract representation of a database and its connection information.
type Datastore interface {
	Session(context.Context) Session
}
