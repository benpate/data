package mockdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopulateInterface(t *testing.T) {

	type person struct {
		ID    int
		Name  string
		Email string
	}

	john := person{
		ID:    1,
		Name:  "John Connor",
		Email: "john@connor.com",
	}

	sarah := person{
		ID:    2,
		Name:  "Sarah Connor",
		Email: "sarah@sky.net",
	}

	target := person{}

	// Populate directly from object
	populateInterface(john, &target)
	assert.Equal(t, 1, target.ID)
	assert.Equal(t, "John Connor", target.Name)
	assert.Equal(t, "john@connor.com", target.Email)

	// Overwrite and populate from pointer
	populateInterface(&sarah, &target)
	assert.Equal(t, 2, target.ID)
	assert.Equal(t, "Sarah Connor", target.Name)
	assert.Equal(t, "sarah@sky.net", target.Email)
}
