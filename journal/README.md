# data.Journal

This package provides *most* of the data and features required by the data.Object interface.  


```go

// Here's a simple definition for a person object
type Person struct {
    PersonID string
    Name     string
    Email    string

    journal.Journal  // Include this line here for all data tracking info.
}

// ID is the only field that you need to define separately from the Journal, 
//in order to make this fit the data.Object interface
func (person *Person) ID() string {
    return person.PersonID
}

// journal.Journal defines all of the other data points and methods required 
// to meet the data.Object interface

// IsNew() bool
// SetCreated(string)
// SetUpdated(string)
// SetDeleted(string)

```