package mongodb

import (
	"github.com/benpate/data/expression"
	"go.mongodb.org/mongo-driver/bson"
)

// ExpressionToBSON converts a data.Expression value into pure bson.
func ExpressionToBSON(criteria expression.Expression) bson.M {

	switch c := criteria.(type) {

	case expression.Predicate:

		result := bson.M{}
		result[c.Field] = bson.M{operatorBSON(c.Operator): c.Value}
		return result

	case expression.AndExpression:

		if len(c) == 0 {
			return nil
		}

		array := bson.A{}

		for _, exp := range c {
			array = append(array, ExpressionToBSON(exp))
		}

		return bson.M{"$and": array}

	case expression.OrExpression:

		if len(c) == 0 {
			return nil
		}

		array := bson.A{}

		for _, exp := range c {
			array = append(array, ExpressionToBSON(exp))
		}

		return bson.M{"$or": array}
	}

	return bson.M{}
}

// operatorBSON converts a standard data.Operator into the operators used by mongodb
func operatorBSON(operator string) string {

	switch operator {
	case expression.OperatorEqual:
		return "$eq"
	case expression.OperatorNotEqual:
		return "$ne"
	case expression.OperatorLessThan:
		return "$lt"
	case expression.OperatorLessOrEqual:
		return "$le"
	case expression.OperatorGreaterOrEqual:
		return "$ge"
	case expression.OperatorGreaterThan:
		return "$gt"
	default:
		return "$eq"
	}
}
