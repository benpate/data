package mongodb

import (
	"encoding/json"
	"testing"

	"github.com/benpate/data"
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

	// Test combining operators into a single bson.M
	ct := data.Expression{{"id", data.OperatorEqual, 42}}
	assert.Equal(t, toJSON(ExpressionToBSON(ct)), `{"id":{"$eq":42},"journal.deleteDate":0}`)

	ct.Add("createDate", data.OperatorGreaterThan, 10)
	assert.Equal(t, toJSON(ExpressionToBSON(ct)), `{"createDate":{"$gt":10},"id":{"$eq":42},"journal.deleteDate":0}`)

	ct.Add("createDate", data.OperatorLessThan, 20)
	assert.Equal(t, toJSON(ExpressionToBSON(ct)), `{"createDate":[{"$gt":10},{"$lt":20}],"id":{"$eq":42},"journal.deleteDate":0}`)

	// Test that all operators are translated correctly.
	ops := data.Expression{}
	ops.Add("=", data.OperatorEqual, 0)
	ops.Add("!=", data.OperatorNotEqual, 0)
	ops.Add("<", data.OperatorLessThan, 0)
	ops.Add("<=", data.OperatorLessOrEqual, 0)
	ops.Add(">", data.OperatorGreaterThan, 0)
	ops.Add(">=", data.OperatorGreaterOrEqual, 0)
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
}
