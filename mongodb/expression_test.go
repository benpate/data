package mongodb

import (
	"encoding/json"
	"testing"

	"github.com/benpate/data/expression"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestExpression(t *testing.T) {

	// toJSON converts values into an easy-to-test JSON string
	toJSON := func(value interface{}) string {

		result, err := json.Marshal(value)

		if err != nil {
			return err.Error()
		}

		return string(result)
	}

	{
		// Test combining operators into a single bson.M
		pred := expression.New("age", expression.OperatorGreaterThan, 42)
		assert.Equal(t, toJSON(ExpressionToBSON(pred)), `{"age":{"$gt":42}}`)

		exp := pred.And("createDate", expression.OperatorEqual, 10)
		assert.Equal(t, toJSON(ExpressionToBSON(exp)), `{"$and":[{"age":{"$gt":42}},{"createDate":{"$eq":10}}]}`)

		exp = exp.And("createDate", expression.OperatorLessThan, 20)
		assert.Equal(t, toJSON(ExpressionToBSON(exp)), `{"$and":[{"age":{"$gt":42}},{"createDate":{"$eq":10}},{"createDate":{"$lt":20}}]}`)
	}

	{
		exp := expression.Or(
			expression.New("name", "=", "John Connor").And("favorite_color", "=", "blue"),
			expression.New("name", "=", "Sara Connor").And("favorite_color", "=", "green"),
		)

		assert.Equal(t, toJSON(ExpressionToBSON(exp)), `{"$or":[{"favorite_color":{"$eq":"blue"},"name":{"$eq":"John Connor"}},{"favorite_color":{"$eq":"green"},"name":{"$eq":"Sara Connor"}}]}`)
	}

	{
		exp := expression.New("name", "=", "John Connor").Or("favorite_color", "=", "blue")
		assert.Equal(t, toJSON(ExpressionToBSON(exp)), `{"$or:[{"favorite_color":{"$eq":"blue"}},{"name":{"$eq":"John Connor"}}]}`)
	}

	{
		exp := expression.And(
			expression.New("name", "=", "John Connor").Or("favorite_color", "=", "blue"),
			expression.New("name", "=", "Sara Connor").Or("favorite_color", "=", "green"),
		)

		t.Log(spew.Sdump(exp))

		assert.Equal(t, toJSON(ExpressionToBSON(exp)), `{"$and":[{"$or:[{"favorite_color":{"$eq":"blue"}},{"name":{"$eq":"John Connor"}}],{"$or:[{"favorite_color":{"$eq":"green"}},{"name":{"$eq":"Sara Connor"}}]}]}`)
		t.Error()
	}
	/*
		// Test that all operators are translated correctly.
		ops := expression.New{
		ops.Add("=", expression.OperatorEqual, 0)
		ops.Add("!=", expression.OperatorNotEqual, 0)
		ops.Add("<", expression.OperatorLessThan, 0)
		ops.Add("<=", expression.OperatorLessOrEqual, 0)
		ops.Add(">", expression.OperatorGreaterThan, 0)
		ops.Add(">=", expression.OperatorGreaterOrEqual, 0)
		ops.Add("OTHER", "OTHER", 0)

		assert.Equal(t, "=", ops[0].Name)
		assert.Equal(t, "=", ops[0].Operator)

		assert.Equal(t, "!=", ops[1].Name)
		assert.Equal(t, "!=", ops[1].Operator)

		assert.Equal(t, "<", ops[2].Name)
		assert.Equal(t, "<", ops[2].Operator)

		assert.Equal(t, "<=", ops[3].Name)
		assert.Equal(t, "<=", ops[3].Operator)

		assert.Equal(t, ">", ops[4].Name)
		assert.Equal(t, ">", ops[4].Operator)

		assert.Equal(t, ">=", ops[5].Name)
		assert.Equal(t, ">=", ops[5].Operator)

		assert.Equal(t, "OTHER", ops[6].Name)
		assert.Equal(t, "=", ops[6].Operator)
	*/
}
