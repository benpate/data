# Journal

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

// Created() int64
// Updated() int64
// IsNew() bool
// IsDeleted() bool
// SetCreated(string)
// SetUpdated(string)
// SetDeleted(string)
// ETag() string

```

## What matters here

- **`SetUpdated` and `SetDeleted` increment `Revision`; `SetCreated` does not.** `Revision` backs `ETag()`, so a freshly created object has an ETag of `"0"`. Adapters rely on this for optimistic concurrency — keep the increment in the mutators.
- **The `Revision` field serializes as `"signature"`** in its `json`/`bson`/`path` tags, for backward compatibility with stored data. Renaming the tag would orphan existing records.

## Pull Requests Welcome

This library is a work in progress, and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making this library better, send in a pull request.  We're all in this together! 📚