package mongodb

import (
	"context"

	"github.com/benpate/data"
	"github.com/benpate/derp"
	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Session represents a single database session, such as a session encompassing all of the database queries to respond to
// a single REST service call.
type Session struct {
	client   *mongo.Client
	context  context.Context
	database string
}

// Load retrieves a single object from the database
func (s Session) Load(collection string, criteria data.Expression, target data.Object) *derp.Error {

	criteriaBSON := ExpressionToBSON(criteria)

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

	object.SetUpdated(note)

	// If new, then INSERT the object
	if object.IsNew() {
		object.SetCreated(note)

		if _, err := s.client.Database(s.database).Collection(collection).InsertOne(s.context, object); err != nil {
			return derp.New(derp.CodeInternalError, "mongodb.Save", "Error inserting object", err.Error(), collection, object)
		}

		return nil
	}

	// Fall through to here means UPDATE object

	filter := objectToFilter(object)
	update := bson.M{"$set": object}

	if _, err := s.client.Database(s.database).Collection(collection).UpdateOne(s.context, filter, update); err != nil {
		return derp.New(derp.CodeInternalError, "mongodb.Save", "Error saving object", err.Error(), collection, filter, update)
	}

	return nil
}

// Delete removes a single object from the database, using a "virtual delete"
func (s Session) Delete(collection string, object data.Object, note string) *derp.Error {

	if object.IsNew() {
		return nil
	}

	// Use virtual delete to mark this object as deleted.
	object.SetDeleted(note)
	return s.Save(collection, object, note)
}

// Close cleans up any remaining connections that need to be removed.
func (s Session) Close() {}

func objectToFilter(object data.Object) bson.M {

	objectID, err := primitive.ObjectIDFromHex(object.ID())

	if err != nil {
		objectID = primitive.NewObjectID()
	}

	result := bson.M{
		"_id":                objectID,
		"journal.deleteDate": 0,
	}

	spew.Dump(result)
	return result
}
