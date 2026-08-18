package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFields(t *testing.T) {

	// Fields returns a concrete FieldsOption wrapped in the Option interface
	option := Fields("name", "email", "age")

	typed, ok := option.(FieldsOption)
	require.True(t, ok)

	assert.Equal(t, TypeFields, option.OptionType())

	// Confirm the data, type, and length of the returned slice
	fields := typed.Fields()
	assert.Len(t, fields, 3)
	assert.Equal(t, []string{"name", "email", "age"}, fields)
}

// A closure-driven test confirms that the variadic argument list is preserved
// exactly across a range of input lengths, including the empty case.
func TestFields_Values(t *testing.T) {

	test := func(input ...string) {
		typed := Fields(input...).(FieldsOption)
		assert.Equal(t, TypeFields, typed.OptionType())
		assert.Len(t, typed.Fields(), len(input))
		assert.Equal(t, input, typed.Fields())
	}

	test()
	test("one")
	test("one", "two")
	test("a", "b", "c", "d", "e")
}

// A spread argument must not alias the option: writes to the caller's slice
// after construction must not be visible through the option.
func TestFields_DoesNotAliasCallerSlice(t *testing.T) {

	names := []string{"name", "email"}
	typed := Fields(names...).(FieldsOption)

	names[0] = "mutated"

	assert.Equal(t, []string{"name", "email"}, typed.Fields())
}

// The accessor must hand back a copy, so an adapter cannot corrupt the option
// it was given by writing through the returned slice.
func TestFields_AccessorReturnsCopy(t *testing.T) {

	typed := Fields("name", "email").(FieldsOption)

	// Ranging (rather than indexing) keeps this safe for a nil slice, which is
	// what Fields() returns when the option carries no field names.
	out := typed.Fields()
	for index := range out {
		out[index] = "mutated"
	}

	assert.Equal(t, []string{"name", "email"}, typed.Fields())
}
