package mongodb

import (
	"github.com/benpate/data"
	"go.mongodb.org/mongo-driver/bson"
)

func Expression2BSON(c data.Expression) bson.M {

	result := bson.M{}

	for _, predicate := range c {

		newBSON := predicate2BSON(predicate)

		// If this field does not yet exist in the result, then just add it and continue the loop
		if _, ok := result[predicate.Name]; !ok {
			result[predicate.Name] = newBSON
			continue
		}

		// fall through to here means that the element already exists in the result.
		switch p := result[predicate.Name].(type) {

		case bson.D:
			result[predicate.Name] = append(p, newBSON)

		case bson.E:
			result[predicate.Name] = bson.D{p, newBSON}
		}
	}

	return result
}

// predicate2BSON converts a standard data.Predicate into a bson.E
func predicate2BSON(predicate data.Predicate) bson.E {
	operator := operator2BSON(predicate.Operator)
	return bson.E{Key: operator, Value: predicate.Value}
}

// operator2BSON converts a standard data.Operator into the operators used by mongodb
func operator2BSON(operator string) string {

	switch operator {
	case data.OperatorEqual:
		return "$eq"
	case data.OperatorNotEqual:
		return "$ne"
	case data.OperatorLessThan:
		return "$lt"
	case data.OperatorLessOrEqual:
		return "$le"
	case data.OperatorGreaterOrEqual:
		return "$ge"
	case data.OperatorGreaterThan:
		return "$gt"
	default:
		return "$eq"
	}
}
