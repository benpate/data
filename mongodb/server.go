package mongodb

import (
	"context"
	"time"

	"github.com/benpate/data"
	"github.com/benpate/derp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Server is an abstract representation of a MongoDB database.  It implements the data.Server interface,
// so that it should be usable anywhere that requires a data.Server.
type Server struct {
	Database *mongo.Database
}

// New returns a fully populated mongodb.Server.  It requires that you provide the URI for the mongodb
// cluster, along with the name of the database to be used for all transactions.
func New(uri string, database string) (Server, *derp.Error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		return Server{}, derp.Wrap(err, "data.mongodb.New", "Error creating mongodb client")
	}

	// Context for connecting to the server only
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return Server{}, derp.Wrap(err, "data.mongodb.New", "Error connecting to mongodb Server")
	}

	result := Server{
		Database: client.Database(database),
	}

	return result, nil
}

// Session returns a new client session that can be used to perform CRUD transactions on this datastore.
func (server Server) Session(ctx context.Context) data.Session {

	return Session{
		database: server.Database,
		context:  ctx,
	}
}
