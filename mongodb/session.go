package mongodb

import (
	"context"

	"github.com/benpate/data"
	"github.com/benpate/derp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Session struct {
	client   *mongo.Client
	context  context.Context
	database string
}

// Load retrieves a single object from the database
func (s Session) Load(collection string, criteria data.Expression, target data.Object) *derp.Error {

	criteriaBSON := Expression2BSON(criteria)

	if err := s.client.Database(s.database).Collection(collection).FindOne(s.context, criteriaBSON).Decode(target); err != nil {

		var errorCode int

		if err.Error() == "mongo: no documents in result" {
			errorCode = derp.CodeNotFoundError
		} else {
			errorCode = derp.CodeInternalError
		}

		return derp.New(errorCode, "mongodb.Load", "Error loading object", err.Error(), collection, criteria, criteriaBSON, target)
	}

	return nil
}

// Save inserts/updates a single object in the database.
func (s Session) Save(collection string, object data.Object, note string) *derp.Error {

	if object.IsNew() {
		object.SetCreated(note)
	}

	object.SetUpdated(note)

	filter := objectToFilter(object.ID())
	update := bson.D{{Key: "$set", Value: object}}
	opts := options.Update().SetUpsert(true)

	if _, err := s.client.Database(s.database).Collection(collection).UpdateOne(s.context, filter, update, opts); err != nil {
		return derp.New(derp.CodeInternalError, "mongodb.Save", "Error saving object", err.Error(), collection, filter, update, opts)
	}

	return nil
}

// Delete removes a single object from the database, using a "virtual delete"
func (s Session) Delete(collection string, object data.Object, note string) *derp.Error {

	object.SetDeleted(note)
	filter := objectToFilter(object.ID())
	update := bson.D{{Key: "$set", Value: object}}
	opts := options.Update().SetUpsert(true)

	if _, err := s.client.Database(s.database).Collection(collection).UpdateOne(s.context, filter, update, opts); err != nil {
		return derp.New(derp.CodeInternalError, "mongodb.Save", "Error updating object", err.Error(), collection, filter, update, opts)
	}

	return nil
}

// Close cleans up any remaining connections that need to be removed.
func (s Session) Close() {}

func objectToFilter(ID string) bson.D {

	objectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		objectID = primitive.NewObjectID()
	}

	return bson.D{
		{Key: "_id", Value: objectID},
		{Key: "journal.deleteDate", Value: 0},
	}
}
