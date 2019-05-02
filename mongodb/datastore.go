package mongodb

import (
	"context"

	"github.com/benpate/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Datastore struct {
	uri      string
	database string
}

func NewDB(uri string, database string) Datastore {

	return Datastore{
		uri:      uri,
		database: database,
	}
}

func (db Datastore) Session(ctx context.Context) data.Session {

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
