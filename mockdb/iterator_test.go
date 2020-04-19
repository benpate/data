package mockdb

import (
	"context"
	"testing"

	"github.com/benpate/data"
	"github.com/benpate/data/option"
	"github.com/stretchr/testify/assert"
)

func TestIterator1(t *testing.T) {

	data := getTestData()

	it := NewIterator(data)

	person := testPerson{}
	counter := 0

	for it.Next(&person) {

		record := data[counter]

		if record, ok := record.(*testPerson); ok {

			assert.Equal(t, record.ID(), person.ID())
			assert.Equal(t, record.Name, person.Name)
			assert.Equal(t, record.Email, person.Email)

		} else {
			t.Error("??? record is not a testPerson type", data[counter])
		}

		counter = counter + 1
	}
}

func TestIterator2(t *testing.T) {

	const collection = "Person"

	data := getTestData()

	s := New().Session(context.TODO())

	for _, record := range data {
		s.Save(collection, record, "Initial Insert")
	}

	it, err := s.List(collection, nil, option.SortAsc("name"))

	if err != nil {
		t.Error(err)
	}

	// Check sort order
	var person testPerson

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "Andrew Jackson", person.Name)

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "Jackson Browne", person.Name)

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "Jessie Jackson", person.Name)

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "Joe Jackson", person.Name)

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "John Connor", person.Name)

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "Kendall Jackson", person.Name)

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "Kyle Reese", person.Name)

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "Michael Jackson", person.Name)

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "Sarah Connor", person.Name)

	if ok := it.Next(&person); !ok {
		t.Fail()
	}

	assert.Equal(t, "Stonewall Jackson", person.Name)

	assert.False(t, it.Next(&person))
}

func TestSafeFieldInterface(t *testing.T) {

	person := testPerson{
		Name: "Joe Jackson",
	}

	value, ok := safeFieldInterface(&person, "name")

	assert.True(t, ok)
	assert.Equal(t, "Joe Jackson", value)
}

func getTestData() []data.Object {

	return []data.Object{
		&testPerson{
			PersonID: "A",
			Name:     "John Connor",
			Email:    "john@connor.com",
		},
		&testPerson{
			PersonID: "B",
			Name:     "Sarah Connor",
			Email:    "sarah@sky.net",
		},
		&testPerson{
			PersonID: "C",
			Name:     "Kyle Reese",
			Email:    "kyle@resistance.mil",
		},
		&testPerson{
			PersonID: "D",
			Name:     "Michael Jackson",
			Email:    "michael@jackson.com",
		},
		&testPerson{
			PersonID: "E",
			Name:     "Joe Jackson",
			Email:    "joe@jackson.com",
		},
		&testPerson{
			PersonID: "F",
			Name:     "Andrew Jackson",
			Email:    "andrew@jackson.com",
		},
		&testPerson{
			PersonID: "G",
			Name:     "Jessie Jackson",
			Email:    "jessie@jackson.com",
		},

		&testPerson{
			PersonID: "H",
			Name:     "Stonewall Jackson",
			Email:    "stonewall@jackson.com",
		},
		&testPerson{
			PersonID: "I",
			Name:     "Kendall Jackson",
			Email:    "kendall@jackson.com",
		},
		&testPerson{
			PersonID: "J",
			Name:     "Jackson Browne",
			Email:    "jackson@jackson.com",
		},
	}
}

/*
func getTestStreamService() Stream {

	// Create service
	datasource := mock.New()
	factory := NewFactory(datasource)
	service := factory.Stream()

	// Initial data to load
	data := []*model.Stream{
		{
			StreamID: primitive.NewObjectID(),
			Label:    "My First Stream",
			Token:    "1-my-first-stream",
		},
		{
			StreamID: primitive.NewObjectID(),
			Label:    "My Second Stream",
			Token:    "2-my-second-stream",
		},
		{
			StreamID: primitive.NewObjectID(),
			Label:    "My Third Stream",
			Token:    "3-my-third-stream",
		},
	}

	// Populate datasource
	for _, record := range data {
		if err := service.Save(record, "comment"); err != nil {
			panic(err)
		}
	}

	return service
}
*/
