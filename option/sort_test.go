package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSortAsc(t *testing.T) {

	option := SortAsc("createDate")

	typed, ok := option.(SortOption)
	require.True(t, ok)

	assert.Equal(t, TypeSort, option.OptionType())
	assert.Equal(t, "createDate", typed.FieldName)
	assert.Equal(t, SortDirectionAscending, typed.Direction)
}

func TestSortDesc(t *testing.T) {

	option := SortDesc("updateDate")

	typed, ok := option.(SortOption)
	require.True(t, ok)

	assert.Equal(t, TypeSort, option.OptionType())
	assert.Equal(t, "updateDate", typed.FieldName)
	assert.Equal(t, SortDirectionDescending, typed.Direction)
}

// A directly-constructed SortOption (no constructor) still reports its type.
func TestSortOption_OptionType(t *testing.T) {

	option := SortOption{FieldName: "name", Direction: SortDirectionAscending}
	assert.Equal(t, TypeSort, option.OptionType())
}

// Only the exact DESC token sorts descending. Every other value -- a typo, a
// lowercase variant, or the empty string of a zero-value SortOption -- must
// fall back to ASCENDING so that adapters cannot disagree about it.
func TestSortOption_DirectionIsValidated(t *testing.T) {

	test := func(direction string, expectedDescending bool) {
		option := SortOption{FieldName: "name", Direction: direction}
		assert.Equal(t, expectedDescending, option.IsDescending(), "direction %q", direction)
		assert.Equal(t, !expectedDescending, option.IsAscending(), "direction %q", direction)
	}

	test(SortDirectionDescending, true)

	test(SortDirectionAscending, false)
	test("", false)
	test("desc", false)
	test("DESCENDING", false)
	test("sideways", false)
}

// The constructors must agree with the accessors.
func TestSortOption_ConstructorsMatchAccessors(t *testing.T) {

	assert.True(t, SortAsc("name").(SortOption).IsAscending())
	assert.False(t, SortAsc("name").(SortOption).IsDescending())

	assert.True(t, SortDesc("name").(SortOption).IsDescending())
	assert.False(t, SortDesc("name").(SortOption).IsAscending())
}

// A closure-driven test confirms field names are passed through unchanged,
// including the empty-string edge case.
func TestSort_FieldNames(t *testing.T) {

	test := func(fieldName string) {
		asc := SortAsc(fieldName).(SortOption)
		assert.Equal(t, fieldName, asc.FieldName)
		assert.Equal(t, SortDirectionAscending, asc.Direction)

		desc := SortDesc(fieldName).(SortOption)
		assert.Equal(t, fieldName, desc.FieldName)
		assert.Equal(t, SortDirectionDescending, desc.Direction)
	}

	test("")
	test("name")
	test("nested.field.path")
}
