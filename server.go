// Package data defines a minimal set of interfaces for abstracting database access
// (servers, sessions, collections, and the objects they store) so that application
// code can perform CRUD operations without depending on a specific database driver.
package data

import (
	"context"
)

// Server is an abstract representation of a database and its connection information.
type Server interface {

	// Session opens a new Session against the database, scoped to the provided context
	Session(context.Context) (Session, error)

	// WithTransaction opens a session, runs the callback, and commits or rolls back based on its result
	WithTransaction(context.Context, TransactionCallbackFunc) (any, error)
}
