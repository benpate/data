package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareInt(t *testing.T) {

	{
		result, err := Interface(1, 1)

		assert.Equal(t, 0, result)
		assert.Nil(t, err)
	}

	{
		result, err := Interface(int8(1), int8(1))

		assert.Equal(t, 0, result)
		assert.Nil(t, err)
	}

	{
		result, err := Interface(int16(1), int16(1))

		assert.Equal(t, 0, result)
		assert.Nil(t, err)
	}

}
