package expression

// AndExpression combines a series of sub-expressions using AND logic
type AndExpression []Expression

// And allows an additional predicate into this AndExpression
func (andExpression AndExpression) And(name string, operator string, value interface{}) AndExpression {
	return append(andExpression, New(name, operator, value))
}

// And combines one or more expression parameters into an AndExpression
func And(expressions ...Expression) AndExpression {

	result := AndExpression{}

	// Add each expression into our result one at a time.
	for _, item := range expressions {

		// Special case.  If the sub-expression is ALSO an AndExpression,
		if items, ok := item.(AndExpression); ok {

			// Then just APPEND it to our current result
			result = append(result, items...)

		} else {
			result = append(result, item)
		}
	}

	return result
}

// Match implements the Expression interface.  It loops through all sub-expressions and returns TRUE if all of them match
func (andExpression AndExpression) Match(fn MatcherFunc) bool {

	for _, expression := range andExpression {

		if expression.Match(fn) == false {
			return false
		}
	}

	return true
}
