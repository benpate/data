package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {

	a := All()

	assert.Equal(t, a, And())
}
