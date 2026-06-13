package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCaseSensitive(t *testing.T) {

	// CaseSensitive returns a concrete CaseSensitiveOption wrapped in the Option interface
	option := CaseSensitive(true)

	typed, ok := option.(CaseSensitiveOption)
	require.True(t, ok)

	assert.Equal(t, TypeCaseSensitive, option.OptionType())
	assert.True(t, typed.CaseSensitive())
}

// A closure-driven test confirms the wrapped boolean is preserved in both directions.
func TestCaseSensitive_Values(t *testing.T) {

	test := func(value bool) {
		typed := CaseSensitive(value).(CaseSensitiveOption)
		assert.Equal(t, value, typed.CaseSensitive())
		assert.Equal(t, TypeCaseSensitive, typed.OptionType())
	}

	test(true)
	test(false)
}
