package mongodb

import (
	"context"

	"github.com/benpate/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Server is an abstract representation of a MongoDB database.  It implements the data.Server interface,
// so that it should be usable anywhere that requires a data.Server.
type Server struct {
	uri      string
	database string
}

// New returns a fully populated mongodb.Server.  It requires that you provide the URI for the mongodb
// cluster, along with the name of the database to be used for all transactions.
func New(uri string, database string) Server {

	return Server{
		uri:      uri,
		database: database,
	}
}

// Session returns a new client session that can be used to perform CRUD transactions on this datastore.
func (db Server) Session(ctx context.Context) data.Session {

	client, err := mongo.NewClient(options.Client().ApplyURI(db.uri))

	if err != nil {
		panic(err.Error())
	}

	if err := client.Connect(ctx); err != nil {
		panic(err.Error())
	}

	return Session{
		client:   client,
		context:  ctx,
		database: db.database,
	}
}
