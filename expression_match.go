package data

import (
	"reflect"
	"strings"
)

// Match uses reflection to compare this expression with an arbitrary struct.
// It is included here to simplify testing and development, but should not be used
// for production-ready code.
func (exp Expression) Match(object interface{}) bool {

	value := reflect.ValueOf(object)

	if value.Kind() == reflect.Ptr {
		value = reflect.Indirect(value)
	}

	for _, predicate := range exp {

		field := value.FieldByNameFunc(func(name string) bool {
			return strings.ToUpper(name) == strings.ToUpper(predicate.Name)
		})

		if field.IsValid() == false {
			return false
		}

		comparison, ok := compareOK(field.Interface(), predicate.Value)

		if ok == false {
			return false
		}

		switch predicate.Operator {

		case OperatorNotEqual:
			if comparison == 0 {
				return false
			}

		case OperatorLessThan:
			if comparison > -1 {
				return false
			}

		case OperatorLessOrEqual:
			if comparison > 0 {
				return false
			}

		case OperatorGreaterThan:
			if comparison < 1 {
				return false
			}

		case OperatorGreaterOrEqual:
			if comparison < 0 {
				return false
			}

		default:
			if comparison != 0 {
				return false
			}
		}
	}

	return true
}

func compareOK(value1 interface{}, value2 interface{}) (int, bool) {

	switch v1 := value1.(type) {

	case int:

		if v2, ok := value2.(int); ok {

			if v1 < v2 {
				return -1, true
			} else if v1 == v2 {
				return 0, true
			} else {
				return 1, true
			}
		}

	case int64:

		if v2, ok := value2.(int64); ok {

			if v1 < v2 {
				return -1, true
			} else if v1 == v2 {
				return 0, true
			} else {
				return 1, true
			}
		}

	case string:

		if v2, ok := value2.(string); ok {

			if v1 < v2 {
				return -1, true
			} else if v1 == v2 {
				return 0, true
			} else {
				return 1, true
			}
		}

	case byte:

		if v2, ok := value2.(byte); ok {

			if v1 < v2 {
				return -1, true
			} else if v1 == v2 {
				return 0, true
			} else {
				return 1, true
			}
		}
	}

	return 0, false
}
