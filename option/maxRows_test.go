package option

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaxRows(t *testing.T) {

	// MaxRows returns a concrete MaxRowsOption wrapped in the Option interface
	option := MaxRows(100)

	typed, ok := option.(MaxRowsOption)
	require.True(t, ok)

	assert.Equal(t, TypeMaxRows, option.OptionType())
	assert.Equal(t, int64(100), typed.MaxRows())
}

// A closure-driven test confirms the wrapped value across boundary inputs.
func TestMaxRows_Values(t *testing.T) {

	test := func(value int64) {
		typed := MaxRows(value).(MaxRowsOption)
		assert.Equal(t, value, typed.MaxRows())
		assert.Equal(t, TypeMaxRows, typed.OptionType())
	}

	test(0)
	test(1)
	test(-1)
	test(math.MaxInt64)
	test(math.MinInt64)
}
