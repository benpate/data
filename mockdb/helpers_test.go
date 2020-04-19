package mockdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {

	{
		first, last := split("Hello World")
		assert.Equal(t, "Hello World", first)
		assert.Equal(t, "", last)
	}

	{
		first, last := split("Hello.World")
		assert.Equal(t, "Hello", first)
		assert.Equal(t, "World", last)
	}

	{
		first, last := split("Hello.World.World.World")
		assert.Equal(t, "Hello", first)
		assert.Equal(t, "World.World.World", last)

		last2, last3 := split(last)
		assert.Equal(t, "World", last2)
		assert.Equal(t, "World.World", last3)

	}
}
