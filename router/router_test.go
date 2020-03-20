package router

import (
	"testing"

	"github.com/benpate/data/memory"
	"github.com/benpate/data/mongodb"
	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {

	config := map[string]string{
		"type": "memory",
	}

	result, err := New(config)

	assert.Nil(t, err)
	assert.IsType(t, &memory.Datastore{}, result)
}

func TestMongoDB(t *testing.T) {

	config := map[string]string{
		"type": "mongodb",
	}

	result, err := New(config)

	assert.Nil(t, err)
	assert.IsType(t, mongodb.Datastore{}, result)
}

func TestError(t *testing.T) {

	config := map[string]string{
		"type": "ERROR",
	}

	result, err := New(config)

	assert.Nil(t, result)

	assert.Equal(t, 500, err.ErrorCode())
}
