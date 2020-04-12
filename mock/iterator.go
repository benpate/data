package mock

import (
	"reflect"

	"github.com/benpate/data"
	"github.com/benpate/data/option"
	"github.com/benpate/derp"
)

// Iterator represents a generic set of data that can be returned by a datasource.
type Iterator struct {
	Data    []data.Object
	Options []option.Option
	Counter int
}

// NewIterator generates
func NewIterator(data []data.Object, options ...option.Option) *Iterator {
	return &Iterator{
		Data:    data,
		Options: options,
		Counter: 0,
	}
}

// Reset moves the iterator back to the beginning of the dataaset
func (iterator *Iterator) Reset() {
	iterator.Counter = 0
}

/// THESE FUNCTIONS IMPLEMENT THE Data.Iterator INTERFACE

// Next moves the Iterator to the next position in the dataset.
// If there is another record in the dataset, it returns TRUE, and
// writes the next record to the "output" variable.
// If there are no more records, it returns FALSE.
func (iterator *Iterator) Next(output data.Object) bool {

	if iterator.Counter >= len(iterator.Data) {
		return false
	}

	populateInterface(iterator.Data[iterator.Counter], output)
	iterator.Counter = iterator.Counter + 1
	return true
}

// Close prevents any further records from being read from the iterator
func (iterator *Iterator) Close() *derp.Error {
	iterator.Counter = len(iterator.Data) + 1
	return nil
}

/// THESE FUNCTIONS IMPLEMENT THE Sort.Interface INTERFACE

// Len returns the number of elements in the collection.
func (iterator *Iterator) Len() int {
	return len(iterator.Data)
}

// Less reports whether the element with index i should sort before the element with index j.
// A return value of TRUE means that the item in position "i" should appear ahead of the item in
// position "j".
func (iterator *Iterator) Less(i int, j int) bool {

	// Range guard
	if i >= len(iterator.Data) {
		return false
	}

	// Range guard
	if j >= len(iterator.Data) {
		return false
	}

	object1 := iterator.Data[i]
	object2 := iterator.Data[j]

	// Look through all options in order.
	for _, record := range iterator.Options {

		// Only use "sort" type options
		if record, ok := record.(option.SortConfig); ok {

			// Get interface{} values for each of the two fields to be compared
			field1 := reflect.Indirect(reflect.ValueOf(object1)).FieldByName(record.FieldName).Interface()
			field2 := reflect.Indirect(reflect.ValueOf(object2)).FieldByName(record.FieldName).Interface()

			// Use generic data.Compare function to compare them
			comparision, err := data.Compare(field1, field2)

			// If these two values cannot be compared, then they cannot be sorted either.  Return FALSE.
			if err != nil {
				return false
			}

			// Return result depends on the direction of the sort order
			switch record.Direction {

			case option.SortDirectionDescending:

				switch comparision {
				case 1:
					return true // IF (i > j) and sort is descending, then i SHOULD appear before j.
				case -1:
					return false // IF (i == j) and sort is descending, then i SHOULD NOT appear before j.
				default:
					// (i == j) so fall through to next comparison
				}

			default: // option.SortDirectionAscending

				switch comparision {
				case -1:
					return true // if (i < j) and sort is ascending, then i SHOULD appear before j.
				case 1:
					return false // if (i > j) and sort is ascending, then i SHOULD NOT appear before j.
				default:
					// (i == j) so fall through to next comparison
				}
			}
		}

		// Fall through to next iteration of loop => for this sort option, the two values are equal.
		// If there's another sort option, then use that as a secondary sort.  Otherwise, fall through
		// all the way to the end of the function
	}

	// Fall through to here means that the two values are equal.
	return false

	// return data.CompareLessThan(iterator.Data[i], iterator.Data[j])
}

// Swpa swpas the elements with indexes i and j
func (iterator *Iterator) Swap(i int, j int) {

	temp := iterator.Data[i]
	iterator.Data[i] = iterator.Data[j]
	iterator.Data[j] = temp
}
