package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpression(t *testing.T) {

	testFn := func(predicate Predicate) bool {

		if value, ok := predicate.Value.(bool); ok {
			return value
		}

		t.Error("Bad Predicate Value for Test Function")

		return false
	}

	// Test Simple AND
	assert.True(t, New("", "", true).And("", "", true).And("", "", true).Match(testFn))
	assert.False(t, New("", "", false).And("", "", true).And("", "", true).Match(testFn))
	assert.False(t, New("", "", true).And("", "", false).And("", "", true).Match(testFn))
	assert.False(t, New("", "", true).And("", "", true).And("", "", false).Match(testFn))
	assert.False(t, New("", "", false).And("", "", false).And("", "", false).Match(testFn))

	// Test Simple OR
	assert.True(t, New("", "", true).Or("", "", true).Or("", "", true).Or("", "", true).Match(testFn))
	assert.True(t, New("", "", false).Or("", "", true).Or("", "", true).Or("", "", true).Match(testFn))
	assert.True(t, New("", "", true).Or("", "", false).Or("", "", true).Or("", "", true).Match(testFn))
	assert.True(t, New("", "", true).Or("", "", true).Or("", "", false).Or("", "", true).Match(testFn))
	assert.False(t, New("", "", false).Or("", "", false).Or("", "", false).Or("", "", false).Match(testFn))

	// Test Compound AND/OR
	assert.True(t, Or(
		New("", "", true).And("", "", true),
		New("", "", true).And("", "", true),
	).Match(testFn))

	assert.True(t, Or(
		New("", "", false).And("", "", true),
		New("", "", true).And("", "", true),
	).Match(testFn))

	assert.True(t, Or(
		New("", "", true).And("", "", true),
		New("", "", false).And("", "", true),
	).Match(testFn))

	assert.False(t, Or(
		New("", "", false).And("", "", true),
		New("", "", false).And("", "", true),
	).Match(testFn))

}

func TestExpressionLots(t *testing.T) {

	expression := And(
		Or(
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
		),
		Or(
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
		),
		Or(
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
		),
		Or(
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
			New("field", "operator", "value").And("field", "operator", "value").And("field", "operator", "value"),
		),
	)

	result := expression.Match(func(predicate Predicate) bool {
		return true
	})

	assert.True(t, result)
}
