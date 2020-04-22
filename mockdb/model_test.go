package mockdb

import "github.com/benpate/data/journal"

type testPerson struct {
	PersonID        string `bson:"_id"`
	Name            string `bson:"name"`
	Email           string `bson:"email"`
	Age             int    `bson:"age"`
	journal.Journal `bson:"journal"`
}

func (person testPerson) ID() string {
	return person.PersonID
}
