package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpression(t *testing.T) {

	c1 := Expression{}

	c1.Add("id", OperatorEqual, 123)
	c1.Add("name", OperatorEqual, "John Connor")

	assert.Equal(t, 2, len(c1))

	assert.Equal(t, "id", c1[0].Name)
	assert.Equal(t, OperatorEqual, c1[0].Operator)
	assert.Equal(t, 123, c1[0].Value)

	assert.Equal(t, "name", c1[1].Name)
	assert.Equal(t, OperatorEqual, c1[1].Operator)
	assert.Equal(t, "John Connor", c1[1].Value)

	c2 := Expression{{"email", OperatorEqual, "john@connor.com"}}

	c3 := c1.Join(c2)

	assert.Equal(t, 3, len(c3))

	assert.Equal(t, "id", c3[0].Name)
	assert.Equal(t, OperatorEqual, c3[0].Operator)
	assert.Equal(t, 123, c3[0].Value)

	assert.Equal(t, "name", c3[1].Name)
	assert.Equal(t, OperatorEqual, c3[1].Operator)
	assert.Equal(t, "John Connor", c3[1].Value)

	assert.Equal(t, "email", c3[2].Name)
	assert.Equal(t, OperatorEqual, c3[2].Operator)
	assert.Equal(t, "john@connor.com", c3[2].Value)
}

func TestPredicate(t *testing.T) {

	p := Predicate{}

	p = Predicate{"id", "=", 0}.Validate()
	assert.Equal(t, OperatorEqual, p.Operator)

	p = Predicate{"id", "!=", 0}.Validate()
	assert.Equal(t, OperatorNotEqual, p.Operator)

	p = Predicate{"id", "<", 0}.Validate()
	assert.Equal(t, OperatorLessThan, p.Operator)

	p = Predicate{"id", "<=", 0}.Validate()
	assert.Equal(t, OperatorLessOrEqual, p.Operator)

	p = Predicate{"id", ">", 0}.Validate()
	assert.Equal(t, OperatorGreaterThan, p.Operator)

	p = Predicate{"id", ">=", 0}.Validate()
	assert.Equal(t, OperatorGreaterOrEqual, p.Operator)

	p = Predicate{"id", "ANYTHING-ELSE", 0}.Validate()
	assert.Equal(t, OperatorEqual, p.Operator)
}
