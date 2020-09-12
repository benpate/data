package expression

// Predicate represents a single true/false comparison
type Predicate struct {
	Field    string
	Operator string
	Value    interface{}
}

// New returns a fully populated Predicate
func New(field string, operator string, value interface{}) Predicate {
	return Predicate{
		Field:    field,
		Operator: operator,
		Value:    value,
	}
}

// Equal creates a new Predicate using an "Equals" comparison
func Equal(field string, value interface{}) Predicate {
	return New(field, OperatorEqual, value)
}

// NotEqual creates a new Predicate using an "Not Equals" comparison
func NotEqual(field string, value interface{}) Predicate {
	return New(field, OperatorNotEqual, value)
}

// LessThan creates a new Predicate using an "Less Than" comparison
func LessThan(field string, value interface{}) Predicate {
	return New(field, OperatorLessThan, value)
}

// LessOrEqual creates a new Predicate using an "Less Or Equal" comparison
func LessOrEqual(field string, value interface{}) Predicate {
	return New(field, OperatorLessOrEqual, value)
}

// GreaterThan creates a new Predicate using an "Greater Than" comparison
func GreaterThan(field string, value interface{}) Predicate {
	return New(field, OperatorGreaterThan, value)
}

// GreaterOrEqual creates a new Predicate using an "Greater Or Equal" comparison
func GreaterOrEqual(field string, value interface{}) Predicate {
	return New(field, OperatorGreaterOrEqual, value)
}

// And combines this predicate with another one (created from the arguments) into an AndExpression
func (predicate Predicate) And(field string, operator string, value interface{}) AndExpression {
	return AndExpression{predicate, New(field, operator, value)}
}

// Or combines this predicate with another one (created from the arguments) into an OrExpression
func (predicate Predicate) Or(field string, operator string, value interface{}) OrExpression {
	return OrExpression{predicate, New(field, operator, value)}
}

// Match implements the Expression interface.  It uses a MatcherFunc to determine if this predicate matches an arbitrary dataset.
func (predicate Predicate) Match(fn MatcherFunc) bool {
	return fn(predicate)
}
