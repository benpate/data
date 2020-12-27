package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// This tests our ability to "collapse" OrExpressions into a single expression, which should keep
// expression trees simpler, and make it easier to traverse/troubleshoot them.
func TestOrExpression(t *testing.T) {

	exp := Or(
		Or(
			New("field0", "=", 0),
		),
		Or(
			New("field1", "=", 1),
			New("field2", "=", 2),
		),
		Or(
			New("field3", "=", 3),
			New("field4", "=", 4),
			Or(
				New("field5", "=", 5),
				New("field6", "=", 6),
			),
		),
	)

	assert.Equal(t, "field0", exp[0].(Predicate).Field)
	assert.Equal(t, "field1", exp[1].(Predicate).Field)
	assert.Equal(t, "field2", exp[2].(Predicate).Field)
	assert.Equal(t, "field3", exp[3].(Predicate).Field)
}
