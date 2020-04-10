package memory

import (
	"testing"

	"github.com/benpate/data"
	"github.com/stretchr/testify/assert"
)

func TestIterator1(t *testing.T) {

	data := []data.Object{
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
	}

	it := NewIterator(data)

	person := testPerson{}

	// First record
	assert.True(t, it.Next(&person))

	assert.Equal(t, "A", person.ID())
	assert.Equal(t, "John Connor", person.Name)
	assert.Equal(t, "john@connor.com", person.Email)

	// Second record
	assert.True(t, it.Next(&person))

	assert.Equal(t, "B", person.ID())
	assert.Equal(t, "Sarah Connor", person.Name)
	assert.Equal(t, "sarah@sky.net", person.Email)

	// Third record
	assert.True(t, it.Next(&person))

	assert.Equal(t, "C", person.ID())
	assert.Equal(t, "Kyle Reese", person.Name)
	assert.Equal(t, "kyle@resistance.mil", person.Email)

	// Done
	assert.False(t, it.Next(&person))
}
