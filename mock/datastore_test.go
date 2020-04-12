package mock

import (
	"context"
	"testing"

	"github.com/benpate/data"
	"github.com/benpate/data/journal"
	"github.com/benpate/derp"
	"github.com/stretchr/testify/assert"
)

func TestDatastore(t *testing.T) {

	ds := New()

	session := ds.Session(context.TODO())

	john := &testPerson{
		PersonID: "A",
		Name:     "John Connor",
		Email:    "john@connor.com",
	}

	// CREATE
	{
		err := session.Save("Person", john, "created in test suite")
		assert.Nil(t, err)
	}

	// READ & UPDATE
	{
		// Load a record from the db
		person := testPerson{}
		criteria := data.Expression{{"personId", "=", "A"}}
		err := session.Load("Person", criteria, &person)
		assert.Nil(t, err)
		assert.Equal(t, "A", person.PersonID)
		assert.Equal(t, "John Connor", person.Name)
		assert.Equal(t, "john@connor.com", person.Email)

		// Change some values
		person.Name = "Sarah Connor"
		person.Email = "sarah@sky.net"

		// Save the record
		err = session.Save("Person", &person, "Comment Here")
		assert.Nil(t, err)

		person2 := testPerson{}
		err = session.Load("Person", criteria, &person2)
		assert.Nil(t, err)
		assert.Equal(t, "Sarah Connor", person2.Name)
		assert.Equal(t, "sarah@sky.net", person2.Email)
		assert.Equal(t, "Comment Here", person2.Journal.Note)
	}

	// "NOT FOUND"
	{
		person := testPerson{}
		criteria := data.Expression{{"missingField", "=", "A"}}
		err := session.Load("Person", criteria, &person)
		assert.NotNil(t, err)
	}
}

func TestList(t *testing.T) {

	ds := New()

	session := ds.Session(context.TODO())

	session.Save("Person", &testPerson{PersonID: "A", Name: "Sarah Connor", Email: "sarah@sky.net"}, "Creating Record")
	session.Save("Person", &testPerson{PersonID: "B", Name: "John Connor", Email: "john@connor.com"}, "Creating Record")
	session.Save("Person", &testPerson{PersonID: "C", Name: "Kyle Reese", Email: "kyle@resistance.mil"}, "Creating Record")

	criteria := data.Expression{}
	criteria.Add("PersonID", "=", "A")

	it, err := session.List("Person", criteria)

	assert.Nil(t, err)

	person := testPerson{}

	assert.True(t, it.Next(&person))
	assert.Equal(t, "A", person.PersonID)
	assert.Equal(t, "Sarah Connor", person.Name)
	assert.Equal(t, "sarah@sky.net", person.Email)

	assert.False(t, it.Next(&person))
}

func TestErrors(t *testing.T) {

	ds := New()

	session := ds.Session(context.TODO())

	person := &testPerson{}

	{
		err := session.Load("MissingCollection", data.Expression{}, person)
		assert.NotNil(t, err)
		assert.Equal(t, derp.CodeNotFoundError, err.Code)
		assert.Equal(t, "Datastore.Load", err.Location)
		assert.Equal(t, "Collection does not exist", err.Message)
		assert.Equal(t, []interface{}{"MissingCollection"}, err.Details)
	}

	{
		err := session.Save("Person", person, "ERROR: Testing error codes")
		assert.NotNil(t, err)
		assert.Equal(t, derp.CodeInternalError, err.Code)
		assert.Equal(t, "Datastore.Save", err.Location)
		assert.Equal(t, "Synthetic Error", err.Message)
	}
}

// MODEL OBJECT

type testPerson struct {
	PersonID string
	Name     string
	Email    string
	journal.Journal
}

func (person testPerson) ID() string {
	return person.PersonID
}
