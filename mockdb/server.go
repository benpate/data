package mockdb

import (
	"context"

	"github.com/benpate/data"
)

// Server is a mock database
type Server map[string]Collection

// New returns a fully initialized Database object
func New() data.Server {
	return Server{}
}

// Session returns a session that can be used as a mock database.
func (db Server) Session(ctx context.Context) data.Session {
	return Session{
		Server: &db,
	}
}
