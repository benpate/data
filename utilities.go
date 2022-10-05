package data

// Slice converts an Iterator into a slice of items.
// You must include a constructor function that generates fully initialized values of the type you want to return.
func Slice[T any](iterator Iterator, constructor func() T) []T {

	result := make([]T, 0, iterator.Count())

	value := constructor()

	for iterator.Next(&value) {
		result = append(result, value)
		value = constructor()
	}

	return result
}
