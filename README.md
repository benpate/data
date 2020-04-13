# Data 📚

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/benpate/data)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/data?style=flat-square)](https://goreportcard.com/report/github.com/benpate/data)
[![Build Status](http://img.shields.io/travis/benpate/data.svg?style=flat-square)](https://travis-ci.org/benpate/data)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/data.svg?style=flat-square)](https://codecov.io/gh/benpate/data)

## Swappable Database Adapters for Go

This library helps to make simple database calls as easily as possible, by providing a simple, common interface that we can implement for every database we need.  The goal of this package is to provide simple [CRUD operations](https://en.wikipedia.org/wiki/Create%2C_read%2C_update_and_delete) only, so each database will support many advanced features that are not available through this library.

### The "Object" interface

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

To implement this quickly, just attach the journal.Journal object to your domain objects, and most of your work is already done.

### Using datasources

```go

// Configure your database
ds := mongodb.New(uri, dbname)

// Create a new session, one per server request.
session := ds.Session()
defer session.Close()

// Returns an iterator that will loop through all records that match the provided criteria.
it, err := session.List("Person", criteria, options...)


// Load from database into a person object
err := session.Load("Person", criteria, &person)

// Insert/Update a person object in the database
err := session.Save("Person", person)

// Delete a person from the database
err := session.Delete("Person", person)

```

### data.mongodb

This adapter implements the data interface for MongoDB.  It uses the standard MongoDB driver.

### data.mock

This adapter implements the data interface for an in-memory datastore.  It is the world's worst database, and should only be used for creating unit tests.  If you use this "database" in production (hell, or even as a proof-of-concept demo) then you deserve the merciless mockery that fate holds for you.

## Retrieving Record Sets

This library also includes an "iterator" interface, for retrieving large sets of data from the datasource efficiently.

```go

// Create an object for the iterator to populate
person := Person{}

// Create the iterator.  Requires a collection name, criteria expression (below), and options (also below, such as sorting and pagination)
it := session.List(CollectionName, CriteriaExpression, Options)

for it.Next(&person) {

    // person.Name...
    // person.Email...
}
```

### Expression Builder

Every database has its own query language, so this library provides in intermediate format that should be easy to convert into whatever specific language you need to use.  Check out the README.md in the expression folder for more details and examples.

```go
// build single predicate expressions
c := expression.New("name", "=", "John Connor")

// or chain logical expressions together
c := expression.New("name", "=", "John Connor").Or("name", "=", "Sarah Connor")
```

### Query Options

There's a package for managing optional query arguments, such as sorting and pagination.  These options just encapsulate data.  It is the responsibilty of 
each database adapter to implement each of these in its own query engine.

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