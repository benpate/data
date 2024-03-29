# Data 📚

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/data)
[![Version](https://img.shields.io/github/v/release/benpate/data?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/data/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/benpate/data/go.yml?branch=main&style=flat-square)](https://github.com/benpate/data/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/data?style=flat-square)](https://goreportcard.com/report/github.com/benpate/data)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/data.svg?style=flat-square)](https://codecov.io/gh/benpate/data)

## Swappable Database Adapters for Go

This library provides a common interface for making simple database calls.  The goal of this package is to provide simple [CRUD operations](https://en.wikipedia.org/wiki/Create%2C_read%2C_update_and_delete) only, so each database will support many advanced features that are not available through this library.  Other modules (such as [data-mongo](https://github.com/benpate/data-mongo) and [data-mock](https://github.com/benpate/data-mock)) implement specific database adapters.

### The "Object" interface

The data library works with any object that implements the `Object` interface.  To implement this quickly in existing data models, you can just attach the `journal.Journal` object to your domain objects, and most of your work is already done.

```go
// Object wraps all of the methods that a Domain Object must provide to Presto
type Object interface {

    // ID returns the primary key of the object
    ID() string

    // IsNew returns TRUE if the object has not yet been saved to the database
    IsNew() bool

    // SetCreated stamps the CreateDate and UpdateDate of the object, and adds a note to the Journal.
    SetCreated(comment string)

    // SetUpdated stamps the UpdateDate of the object, and adds a note to the Journal.
    SetUpdated(comment string)

    // SetDeleted marks the object virtually "deleted", and adds a note to the Journal.
    SetDeleted(comment string)
}
```

### Datasource Interface

```go

// Configure your database
ds := mongodb.New(uri, dbname)

// Create a new session, one per server request
session := db.Session()
defer session.Close()

// Get a reference to the table/collection you're working with
collection := session.Collection("Person")

// LOAD from a person object from the database
err := collection.Load(criteria, &person)

// INSERT/UPDATE a person object in the database - writing a "note" to the journal.
err := collection.Save(person, note)

// DELETE a person from the database. This is a "soft" delete that marks values as deleted but leaves them in the database.
err := collection.Delete(person, note)

// HARD DELETE a person, actually removing the record from the database.
err := collection.HardDelete(criteria)

// QUERY many records from the database, by populating a slice of results
err := collection.Query(criteria, options...)

// ITERATE over many records, using a more memory-friendly iterator to loop through a very large dataset.
it, err := collection.Query(criteria, options...)

for it.Next(&person) {
    // do stuff with person.
}

// COUNT the number of records that match a particular criteria
count, err := collection.Count(criteria)

```

### Expression Builder

Every database has its own query language, so this library uses the [exp module](https://github.com/benpate/exp) to represent query expressions in an efficient intermediate format that should be easy to convert into whatever specific language you need to use.

```go
// build single predicate expressions
c := exp.Equal("name", "John Connor")

// or chain logical expressions together
c := exp.Equal("name", "John Connor").OrEqual("name", "Sarah Connor")
```

### Query Options

There's a package for managing optional query arguments, such as sorting and pagination.  These options just encapsulate data.  It is the responsibilty of each database adapter to implement each of these in its own query engine.

```go

// get a new iterator.  Sort results by first name.  Return only the first 100 rows.
it := session.List(collection, criteria, options.SortAsc("name"), options.MaxRows(100))
```

**`SortAsc(fieldname)`** tells the database to sort by a particular field, in ascending order

**`SortDesc(fieldname)`** tells the database to sort by a particular field, in descending order

**`FirstRow(count)`** tells the database to start returning records at the provided row number

**`MaxRows(count)`** tells the database to limit the number of records to the designated number of rows.

## Pull Requests Welcome

This library is a work in progress, and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making this library better, send in a pull request.  We're all in this together! 📚
