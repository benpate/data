package data

// Predicate represents a single expression, such as [name = "John Connor"]
type Predicate struct {
	Name     string      // The name of the field being compared
	Operator string      // The type of comparison (=, !=, >, >=, <, <=).  If this value is empty string, it is assumed to be "="
	Value    interface{} // The value that the field is being compared to
}

// Validate verifies that the values in the predicate are well formed.  At this point,
// this just means that the operator is one of the valid operations.
func (p Predicate) Validate() Predicate {

	switch p.Operator {
	case OperatorEqual:
	case OperatorNotEqual:
	case OperatorLessThan:
	case OperatorLessOrEqual:
	case OperatorGreaterThan:
	case OperatorGreaterOrEqual:
	default:
		p.Operator = OperatorEqual
	}

	return p
}
