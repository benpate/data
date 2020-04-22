package mockdb

import (
	"context"

	"github.com/benpate/data"
)

// Datastore is a mock database
type Datastore map[string]Collection

// New returns a fully initialized Database object
func New() data.Datastore {
	return Datastore{}
}

// Session returns a session that can be used as a mock database.
func (db Datastore) Session(ctx context.Context) data.Session {
	return Session{
		Datastore: &db,
	}
}
