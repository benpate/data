package mock

import (
	"context"
	"testing"

	"github.com/benpate/data"
	"github.com/benpate/data/option"
	"github.com/davecgh/go-spew/spew"
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

	it, err := s.List(collection, nil, option.SortAsc("Name"))

	if err != nil {
		t.Error(err)
	}

	spew.Dump(it)
	t.Fail()
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
			Name:     "Stonewall  Jackson",
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
