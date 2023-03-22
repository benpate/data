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

// Channel converts an Iterator into a channel of items.
// You must include a constructor function that generates fully initialized values of the type you want to return.
func Channel[T any](iterator Iterator, constructor func() T) chan T {

	if iterator.Count() == 0 {
		return nil
	}

	result := make(chan T, 1) // Length of 1 to prevent blocking on the first item.

	go func() {

		value := constructor()

		for iterator.Next(&value) {
			result <- value
			value = constructor()
		}
		close(result)
	}()

	return result
}

// Channel converts an Iterator into a channel of items.
// You must include a constructor function that generates fully initialized values of the type you want to return.
func ChannelWithCancel[T any](iterator Iterator, constructor func() T, cancel <-chan bool) chan T {

	if iterator.Count() == 0 {
		return nil
	}

	result := make(chan T, 1) // Length of 1 to prevent blocking on the first item.

	go func() {
		value := constructor()

		for iterator.Next(&value) {
			select {
			case <-cancel:
				close(result)
				return

			default:
				result <- value
			}

			value = constructor()
		}

		close(result)
	}()

	return result
}
