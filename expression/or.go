package expression

// OrExpression compares a series of sub-expressions, using the OR logic
type OrExpression []Expression

// Or appends an additional predicate into the OrExpression
func (orExpression OrExpression) Or(name string, operator string, value interface{}) OrExpression {
	return append(orExpression, New(name, operator, value))
}

// Or combines one or more expression parameters into an OrExpression
func Or(expressions ...Expression) OrExpression {
	return OrExpression(expressions)
}

// Match implements the Expression interface.  It loops through all sub-expressions and returns TRUE if any of them match
func (orExpression OrExpression) Match(fn MatcherFunc) bool {

	for _, expression := range orExpression {

		if expression.Match(fn) == true {
			return true
		}
	}

	return false
}
