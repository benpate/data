package mongodb

import (
	"encoding/json"
	"testing"

	"github.com/benpate/criteria"
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
	ct := data.Expression{{"id", criteria.OperatorEqual, 42}}
	assert.Equal(t, toJSON(Expression2BSON(ct)), `{"id":{"Key":"$eq","Value":42}}`)

	ct.Add("createDate", criteria.OperatorGreaterThan, 10)
	assert.Equal(t, toJSON(Expression2BSON(ct)), `{"createDate":{"Key":"$gt","Value":10},"id":{"Key":"$eq","Value":42}}`)

	ct.Add("createDate", criteria.OperatorLessThan, 20)
	assert.Equal(t, toJSON(Expression2BSON(ct)), `{"createDate":[{"Key":"$gt","Value":10},{"Key":"$lt","Value":20}],"id":{"Key":"$eq","Value":42}}`)

	// Test that all operators are translated correctly.
	ops := data.Expression{}
	ops.Add("=", criteria.OperatorEqual, 0)
	ops.Add("!=", criteria.OperatorNotEqual, 0)
	ops.Add("<", criteria.OperatorLessThan, 0)
	ops.Add("<=", criteria.OperatorLessOrEqual, 0)
	ops.Add(">", criteria.OperatorGreaterThan, 0)
	ops.Add(">=", criteria.OperatorGreaterOrEqual, 0)
	ops.Add("OTHER", "OTHER", 0)

	t.Fail()
}
