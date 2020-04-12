package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareInt(t *testing.T) {

	{
		result, err := Compare(1, 1)

		assert.Equal(t, 0, result)
		assert.Nil(t, err)
	}

	{
		result, err := Compare(int8(1), int8(1))

		assert.Equal(t, 0, result)
		assert.Nil(t, err)
	}

	{
		result, err := Compare(int16(1), int16(1))

		assert.Equal(t, 0, result)
		assert.Nil(t, err)
	}

}
