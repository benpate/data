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

    // SetCreated stamps the CreateDate and UpdateDate of the object, and makes a note
    SetCreated(comment string)

    // SetUpdated stamps the UpdateDate of the object, and makes a note
    SetUpdated(comment string)

    // SetDeleted marks the object virtually "deleted", and makes a note
    SetDeleted(comment string)
}
```

To implement this quickly, just attach the data.Journal object to your domain objects, and most of your work is already done.

### Using datasources

```go

// Configure your database
ds := mongodb.New(uri, dbname)

// Create a new session, one per server request.
session := ds.Session()
defer session.Close()

// Load from database into a person object
err := session.Load("Person", criteria, &person)

// Insert/Update a person object in the database
err := session.Save("Person", person)

// Delete a person from the database
err := session.Delete("Person", person)

```

### data.mongodb

This adapter implements the data interface for MongoDB.  It uses the standard MongoDB driver.

### data.memory

This adapter implements the data interface for an in-memory datastore.  It is the world's worst database, and should only be used for creating unit tests.  If you use this "database" in production (hell, or even as a proof-of-concept demo) then you deserve the merciless mockery that fate holds for you.

## Minimal Expression Builder

Every database has its own query language, so this library provides in intermediate format that should be easy to convert into whatever specific language you need to use.

```go
// build a data directly
c := data.Criteria{{"id", "=", 42}, {"deleteDate", "=", 0}}

// build a data incrementally
c := data.Criteria{}

c.Add("id", data.OperatorEqual, 42)
c.Add("deleteDate", data.OperatorEqual, 0)

// combine data expressions into a single value
finalResult = data.combined(data1, data2, data3)

// Constants define standard expected operators
data.OperatorEqual    = "="
data.OperatorNotEqual = "!="
data.LessThan         = "<"
data.LessOrEqual      = "<="
data.GreaterThan      = ">"
data.GreaterOrEqual   = ">="
```