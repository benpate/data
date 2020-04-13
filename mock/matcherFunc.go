package mock

import (
	"reflect"

	"github.com/benpate/data"
	"github.com/benpate/data/compare"
	"github.com/benpate/data/expression"
)

// MatcherFunc is a helper function that uses reflection to look inside a generic data.Object and match it.
// Because it uses reflection, it should be considered SLOW, and only be used in the mock library.
func MatcherFunc(object data.Object) expression.MatcherFunc {

	return func(predicate expression.Predicate) bool {

		value := reflect.Indirect(reflect.ValueOf(object))
		structure := value.Type()

		_, field, ok := findField(structure, value, predicate.Field)

		if ok == false {
			return false
		}

		result, _ := compare.WithOperator(field.Interface(), predicate.Operator, predicate.Value)

		return result

		/*

			// Search every field in the structure
			for index := 0; index < structure.NumField(); index = index + 1 {

				field := structure.Field(index)

				// If the field has a bson tag...
				if bson, ok := field.Tag.Lookup("bson"); ok {

					// If the bson tag matches the predicate field
					if bson == predicate.Field {
						result, _ := compare.WithOperator(value.Field(index).Interface(), predicate.Operator, predicate.Value)

						return result
					}
				}
			}

			return false
		*/
	}
}
