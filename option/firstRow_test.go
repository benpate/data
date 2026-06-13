package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFirstRow(t *testing.T) {

	// FirstRow returns an empty FirstRowOption wrapped in the Option interface
	option := FirstRow()

	_, ok := option.(FirstRowOption)
	require.True(t, ok)

	assert.Equal(t, TypeFirstRow, option.OptionType())
}
