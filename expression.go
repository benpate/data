// Package data provides a data structure for defining simple database filters.  This
// is not able to represent every imaginable query criteria, but it does a good job of making
// common criteria simple to format and pass around in your application.
package data

// Expression represents a comparison or filter, to be used when loading objects from the database.
// Each service and database driver is expected to use this common format for query criteria, and
// then map it into the specific format required for that database.
type Expression []struct {
	Name     string      // The name of the field being compared
	Operator string      // The type of comparison (=, !=, >, >=, <, <=).  If this value is empty string, it is assumed to be "="
	Value    interface{} // The value that the field is being compared to
}

// Add appends a new predicate to an existing criteria expression.
func (exp *Expression) Add(name string, operator string, value interface{}) *Expression {

	p := Predicate{name, operator, value}.Validate()

	(*exp) = append(*exp, p)

	return exp
}

// Combine merges together the predicates of all provided criteria expressions, into a single criteria expression.
func (exp Expression) Join(criteriaToCombine ...Expression) Expression {

	for _, criteria := range criteriaToCombine {
		exp = append(exp, criteria...)
	}

	return exp
}
