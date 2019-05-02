package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {

	type person struct {
		ID    int
		Name  string
		Email string
		CreateDate int64
	}

	john := person{
		ID:    42,
		Name:  "John Connor",
		Email: "john@connor.com",
		CreateDate: 0,
	}

	sarah := person{
		ID:    43,
		Name:  "Sarah Connor",
		Email: "sarah@sky.net",
		CreateDate: 1,
	}

	kyle := person{
		ID:    44,
		Name:  "Kyle Reese",
		Email: "kyle@resistance.org",
		CreateDate: 2,
	}

	{
		// Test INTEGER equality
		exp := Expression{{"id", "=", 42}}
		assert.Equal(t, true, exp.Match(john))
		assert.Equal(t, false, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))
	}

	{
		// Test INTEGER inequality
		exp := Expression{{"id", "!=", 44}}
		assert.Equal(t, true, exp.Match(john))
		assert.Equal(t, true, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))
	}

	{
		// Test INTEGER less than
		exp := Expression{{"id", "<", 44}}
		assert.Equal(t, true, exp.Match(john))
		assert.Equal(t, true, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))
	}

	{
		// Test INTEGER greater than
		exp := Expression{{"id", ">", 44}}
		assert.Equal(t, false, exp.Match(john))
		assert.Equal(t, false, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))
	}

	{
		// Test INTEGER greater or equal
		exp := Expression{{"id", ">=", 44}}
		assert.Equal(t, false, exp.Match(john))
		assert.Equal(t, false, exp.Match(sarah))
		assert.Equal(t, true, exp.Match(kyle))
	}

	{
		// Test INTEGER less or equal
		exp6 := Expression{{"id", "<=", 43}}
		assert.Equal(t, true, exp6.Match(john))
		assert.Equal(t, true, exp6.Match(sarah))
		assert.Equal(t, false, exp6.Match(kyle))
	}

	{
		// Test INTEGER type mismatch
		exp6 := Expression{{"id", ">=", "Michael Jackson"}}
		assert.Equal(t, false, exp6.Match(john))
		assert.Equal(t, false, exp6.Match(sarah))
		assert.Equal(t, false, exp6.Match(kyle))
	}

	{
		// Test INT64 equality
		exp := Expression{{"createDate", "=", 0}}
		assert.Equal(t, true, exp.Match(john))
		assert.Equal(t, false, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))
	}

	{
		// Test INT64 inequality
		exp := Expression{{"CREATEDATE", "!=", 1}}
		assert.Equal(t, true, exp.Match(john))
		assert.Equal(t, false, exp.Match(sarah))
		assert.Equal(t, true, exp.Match(kyle))
	}

	{
		// Test INT64 less than
		exp := Expression{{"createdate", "<", 2}}
		assert.Equal(t, true, exp.Match(john))
		assert.Equal(t, true, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))
	}

	{
		// Test INT64 greater than
		exp := Expression{{"cReAtEdAtE", ">", 1}}
		assert.Equal(t, false, exp.Match(john))
		assert.Equal(t, false, exp.Match(sarah))
		assert.Equal(t, true, exp.Match(kyle))
	}

	{
		// Test INT64 greater or equal
		exp := Expression{{"CrEaTeDaTe", ">=", 1}}
		assert.Equal(t, false, exp.Match(john))
		assert.Equal(t, true, exp.Match(sarah))
		assert.Equal(t, true, exp.Match(kyle))
	}

	{
		// Test INT64 less or equal
		exp := Expression{{"createDate", "<=", 3}}
		assert.Equal(t, true, exp.Match(john))
		assert.Equal(t, true, exp.Match(sarah))
		assert.Equal(t, true, exp.Match(kyle))
	}

	{
		// Test INT64 comparisons
		assert.Equal(t, true, Expression{{"createDate", "=", int64(0)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", "!=", int64(1)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", "<", int64(1)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", "<=", int64(0)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", "<=", int64(1)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", ">", int64(-1)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", ">=", int64(-1)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", ">=", int64(0)}}.Match(john))

		assert.Equal(t, false, Expression{{"createDate", "=", int64(1)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", "!=", int64(0)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", "<", int64(-1)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", "<=", int64(-1)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", ">", int64(1)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", ">=", int64(1)}}.Match(john))
	}

	{
		// Test INT64 type mismatch
		exp := Expression{{"createDate", "<=", "STRING"}}
		assert.Equal(t, false, exp.Match(john))
		assert.Equal(t, false, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))
	}

	// Test multiple fields
	{
		exp := Expression{{"name", "=", "John Connor"}, {"id", "=", 42}}
		assert.Equal(t, true, exp.Match(john))
		assert.Equal(t, false, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))
	}

	{
		// Test multiple fields
		exp := Expression{{"name", ">", "John Connor"}, {"id", "<", 44}}
		assert.Equal(t, false, exp.Match(john))
		assert.Equal(t, true, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))

		// Test pointers
		assert.Equal(t, false, exp.Match(&john))
		assert.Equal(t, true, exp.Match(&sarah))
		assert.Equal(t, false, exp.Match(&kyle))
	}

	{
		// Test missing fields
		exp := Expression{{"missing-field", ">", "John Connor"}}
		assert.Equal(t, false, exp.Match(john))
		assert.Equal(t, false, exp.Match(sarah))
		assert.Equal(t, false, exp.Match(kyle))
	}

	{
		// Test missing fields
		exp1 := Expression{{"id", "=", "John Connor"}}
		assert.Equal(t, false, exp1.Match(john))
		assert.Equal(t, false, exp1.Match(sarah))
		assert.Equal(t, false, exp1.Match(kyle))
	}

	{
		// Test string comparisons
		assert.Equal(t, true, Expression{{"name", "=", "John Connor"}}.Match(john))
		assert.Equal(t, true, Expression{{"name", ">=", "John Connor"}}.Match(john))
		assert.Equal(t, true, Expression{{"name", "<=", "John Connor"}}.Match(john))
		assert.Equal(t, true, Expression{{"name", "!=", "A"}}.Match(john))
		assert.Equal(t, true, Expression{{"name", "<", "Klaus"}}.Match(john))
		assert.Equal(t, true, Expression{{"name", "<=", "Kaus"}}.Match(john))
		assert.Equal(t, true, Expression{{"name", ">", "Ignacio"}}.Match(john))
		assert.Equal(t, true, Expression{{"name", ">=", "Ignacio"}}.Match(john))

		assert.Equal(t, false, Expression{{"name", "=", "Sarah Connor"}}.Match(john))
		assert.Equal(t, false, Expression{{"name", "<", "John Connor"}}.Match(john))
		assert.Equal(t, false, Expression{{"name", ">", "John Connor"}}.Match(john))
		assert.Equal(t, false, Expression{{"name", ">=", "Klaus"}}.Match(john))
		assert.Equal(t, false, Expression{{"name", "<=", "Ignacio"}}.Match(john))
		assert.Equal(t, false, Expression{{"name", "!=", "John Connor"}}.Match(john))
		assert.Equal(t, false, Expression{{"name", "<", "Ignacio"}}.Match(john))
		assert.Equal(t, false, Expression{{"name", ">", "Klaus"}}.Match(john))
	}

	{
		// Test INT / INT64 type mismatch
		assert.Equal(t, true, Expression{{"id", "=", int64(42)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", "=", 0}}.Match(john))
		assert.Equal(t, false, Expression{{"id", "=", int64(43)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", "=", 1}}.Match(john))

		assert.Equal(t, true, Expression{{"id", "<", int64(43)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", "<", 1}}.Match(john))
		assert.Equal(t, false, Expression{{"id", "<", int64(40)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", "<", 0}}.Match(john))

		assert.Equal(t, true, Expression{{"id", "<=", int64(42)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", "<=", 0}}.Match(john))
		assert.Equal(t, false, Expression{{"id", "<=", int64(40)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", "<=", -1}}.Match(john))
		assert.Equal(t, true, Expression{{"id", "<=", int64(43)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", "<=", 1}}.Match(john))
		assert.Equal(t, false, Expression{{"id", "<=", int64(40)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", "<=", -1}}.Match(john))

		assert.Equal(t, true, Expression{{"id", ">", int64(40)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", ">", -1}}.Match(john))
		assert.Equal(t, false, Expression{{"id", ">", int64(44)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", ">", 1}}.Match(john))

		assert.Equal(t, true, Expression{{"id", ">=", int64(42)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", ">=", 0}}.Match(john))
		assert.Equal(t, false, Expression{{"id", ">=", int64(43)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", ">=", 1}}.Match(john))
		assert.Equal(t, true, Expression{{"id", ">=", int64(42)}}.Match(john))
		assert.Equal(t, true, Expression{{"createDate", ">=", 0}}.Match(john))
		assert.Equal(t, false, Expression{{"id", ">=", int64(44)}}.Match(john))
		assert.Equal(t, false, Expression{{"createDate", ">=", 1}}.Match(john))
	}
}
