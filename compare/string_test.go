package compare

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsString(t *testing.T) {

	require.True(t, Contains("one", "on"))
	require.True(t, Contains("one", "ne"))
	require.True(t, Contains("three", "th"))
	require.True(t, Contains("three", "thr"))
	require.True(t, Contains("three", "hre"))
	require.True(t, Contains("three", "hree"))

	require.False(t, Contains("one", "four"))
}

func TestContainsArray(t *testing.T) {

	require.True(t, Contains([]string{"one", "two", "three"}, "one"))
	require.True(t, Contains([]string{"one", "two", "three"}, "two"))
	require.True(t, Contains([]string{"one", "two", "three"}, "three"))
	require.False(t, Contains([]string{"one", "two", "three"}, "four"))
	
	require.False(t, Contains([]string{}, "empty"))
}