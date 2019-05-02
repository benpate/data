package memory

import (
	"testing"
	"time"
	"context"
	"github.com/benpate/data"
	"github.com/stretchr/testify/assert"
)

func TestDatastore(t *testing.T) {

	ds := New()

	session := ds.Session(context.TODO())

	john := &testPerson{
		PersonID: "A",
		Name: "John Connor",
		Email: "john@connor.com",
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
		assert.Equal(t, "Comment Here", person2.Comment)
	}

	// "NOT FOUND"
	{
		person := testPerson{}
		criteria := data.Expression{{"missingField", "=", "A"}}
		err := session.Load("Person", criteria, &person); 
		assert.NotNil(t, err)
	}

}

// MODEL OBJECT

type testPerson struct {
	PersonID   string
	Name       string
	Email      string
	CreateDate int64 
	UpdateDate int64 
	Comment    string
}

func (person *testPerson) ID() string {
	return person.PersonID
}

func (person *testPerson) IsNew() bool {
	return person.CreateDate == 0
}

func (person *testPerson) SetCreated(comment string) {
	person.CreateDate = time.Now().Unix()
	person.Comment = comment
}

func (person *testPerson) SetUpdated(comment string) {
	person.UpdateDate = time.Now().Unix()
	person.Comment = comment
}

func (person *testPerson) SetDeleted(comment string) {
}
