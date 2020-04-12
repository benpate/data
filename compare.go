package data

import "github.com/benpate/derp"

// Compare tries its best to muscle value2 and value2 into compatable types so that they can be compared.
// If value1 is LESS THAN value2, it returns -1, nil
// If value1 is EQUAL TO value2, it returns 0, nil
// If value1 is GREATER THAN value2, it returns 1, nil
// If the two values are not compatable, then it returns 0, [DERP] with an explanation of the error.
// Currently, this function ONLLY compares identical numeric or string types.  In the future, it *may*
// be expanded to perform simple type converstions between similar types.
func Compare(value1 interface{}, value2 interface{}) (int, *derp.Error) {

	switch v1 := value1.(type) {

	case int:

		if v2, ok := value2.(int); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case int8:

		if v2, ok := value2.(int8); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case int16:

		if v2, ok := value2.(int16); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case int32:

		if v2, ok := value2.(int32); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case int64:

		if v2, ok := value2.(int64); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case uint:

		if v2, ok := value2.(uint); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}
	case uint8:

		if v2, ok := value2.(uint8); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case uint16:

		if v2, ok := value2.(uint16); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case uint32:

		if v2, ok := value2.(uint32); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case uint64:

		if v2, ok := value2.(uint64); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case float32:

		if v2, ok := value2.(float32); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case float64:

		if v2, ok := value2.(float64); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}

	case string:

		if v2, ok := value2.(string); ok {

			if v1 < v2 {
				return -1, nil
			}
			if v1 == v2 {
				return 0, nil
			}
			return 1, nil
		}
	}

	return 0, derp.New(500, "data.Compare", "Incompatable Types", value1, value2)
}

// CompareEqual is a simplified version of Compare.  It ONLY returns true if the two provided values are EQUAL.
// In all other cases (including errors) it returns FALSE
func CompareEqual(value1 interface{}, value2 interface{}) bool {

	if result, err := Compare(value1, value2); err == nil {

		if result == 0 {
			return true
		}
	}

	return false
}

// CompareLessThan is a simplified version of Compare.  It ONLY returns true if value1 is verifiably LESS THAN value2.
// In all other cases (including errors) it returns FALSE
func CompareLessThan(value1 interface{}, value2 interface{}) bool {

	if result, err := Compare(value1, value2); err == nil {

		if result == -1 {
			return true
		}
	}

	return false
}

// CompareGreaterThan is a simplified version of Compare.  It ONLY returns true if value1 is verifiably GREATER THAN value2.
// In all other cases (including errors) it returns FALSE
func CompareGreaterThan(value1 interface{}, value2 interface{}) bool {

	if result, err := Compare(value1, value2); err == nil {

		if result == 1 {
			return true
		}
	}

	return false
}
