package data

import "context"

// Session represents a single database session, that is opened to support a single transactional request, and then closed
// when this transaction is complete
type Session interface {

	// Collection returns a handle to the named collection (or table) within this session
	Collection(collection string) Collection

	// Context returns the context.Context that scopes this session
	Context() context.Context

	// Close releases any resources held by this session
	Close()
}
