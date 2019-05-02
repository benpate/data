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

func TestMatch(t *testing.T) {

	type person struct {
		ID    int
		Name  string
		Email string
	}

	john := person{
		ID:    42,
		Name:  "John Connor",
		Email: "john@connor.com",
	}

	sarah := person{
		ID:    43,
		Name:  "Sarah Connor",
		Email: "sarah@sky.net",
	}

	kyle := person{
		ID:    44,
		Name:  "Kyle Reese",
		Email: "kyle@resistance.org",
	}

	exp1 := Expression{{"name", "=", "John Connor"}}
	assert.Equal(t, true, exp1.Match(john))
	assert.Equal(t, false, exp1.Match(sarah))
	assert.Equal(t, false, exp1.Match(kyle))

	exp2 := Expression{{"email", "!=", "kyle@resistance.org"}}
	assert.Equal(t, true, exp2.Match(john))
	assert.Equal(t, true, exp2.Match(sarah))
	assert.Equal(t, false, exp2.Match(kyle))

	exp3 := Expression{{"id", "<", 44}}
	assert.Equal(t, true, exp3.Match(john))
	assert.Equal(t, true, exp3.Match(sarah))
	assert.Equal(t, false, exp3.Match(kyle))

	exp4 := Expression{{"id", ">", 44}}
	assert.Equal(t, false, exp4.Match(john))
	assert.Equal(t, false, exp4.Match(sarah))
	assert.Equal(t, false, exp4.Match(kyle))

	exp5 := Expression{{"id", ">=", 44}}
	assert.Equal(t, false, exp5.Match(john))
	assert.Equal(t, false, exp5.Match(sarah))
	assert.Equal(t, true, exp5.Match(kyle))

	exp6 := Expression{{"name", "<=", "Kyle Reese"}}
	assert.Equal(t, true, exp6.Match(john))
	assert.Equal(t, false, exp6.Match(sarah))
	assert.Equal(t, true, exp6.Match(kyle))

	exp7 := Expression{{"name", "=", "John Connor"}, {"id", "=", 42}}
	assert.Equal(t, true, exp7.Match(john))
	assert.Equal(t, false, exp7.Match(sarah))
	assert.Equal(t, false, exp7.Match(kyle))

	exp8 := Expression{{"name", ">", "John Connor"}, {"id", "<", 44}}
	assert.Equal(t, false, exp8.Match(john))
	assert.Equal(t, true, exp8.Match(sarah))
	assert.Equal(t, false, exp8.Match(kyle))
}
