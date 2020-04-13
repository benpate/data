package expression

// NotExpression negates the single expression it contains
type NotExpression struct {
	SubExpression Expression
}

// Not generates a NotExpression around the argument provided
func Not(expression Expression) NotExpression {
	return NotExpression{
		SubExpression: expression,
	}
}

// Match implements the Expression interface.  It returns the opposite of
// the result from Matching the SubExpression.
func (notExpression NotExpression) Match(fn MatcherFunc) bool {
	return (notExpression.SubExpression.Match(fn) == false)
}
