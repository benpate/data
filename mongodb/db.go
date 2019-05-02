package mongodb

import (
	"context"

	"github.com/benpate/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	uri      string
	database string
}

func NewDB(uri string, database string) DB {

	return DB{
		uri:      uri,
		database: database,
	}
}

func (db DB) Session(ctx context.Context) data.DBSession {

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
