package mockdb

import "github.com/benpate/data"

// Collection is a mock table
type Collection []data.Object

// FindByObjectID does a linear search on the collection for the first object with a matching ID()
func (collection Collection) FindByObjectID(objectID string) int {

	for index, object := range collection {

		if object.ID() == objectID {
			return index
		}
	}

	return -1
}
