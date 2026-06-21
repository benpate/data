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
// Object defines all of the methods that a domain object must provide to the data library
type Object interface {

    // ID returns the primary key of the object
    ID() string

    // Created and Updated return the Unix epoch (milliseconds) when the object was created/updated
    Created() int64
    Updated() int64

    // IsNew returns TRUE if the object has not yet been saved to the database
    IsNew() bool

    // IsDeleted returns TRUE if the object has been virtually deleted
    IsDeleted() bool

    // SetCreated stamps the CreateDate and UpdateDate of the object, and adds a note to the Journal.
    SetCreated(comment string)

    // SetUpdated stamps the UpdateDate of the object, and adds a note to the Journal.
    SetUpdated(comment string)

    // SetDeleted marks the object virtually "deleted", and adds a note to the Journal.
    SetDeleted(comment string)

    // ETag returns the signature or revision number of the object
    ETag() string
}
```

### Datasource Interface

```go

// Configure your database
server := mongodb.New(uri, dbname)

// Create a new session, one per server request
session, err := server.Session(ctx)
defer session.Close()

// Get a reference to the table/collection you're working with
collection := session.Collection("Person")

// LOAD the first person that matches the criteria into a person object
err := collection.Load(criteria, &person)

// INSERT/UPDATE a person object in the database - writing a "note" to the journal.
err := collection.Save(person, note)

// DELETE a person from the database. This is a "soft" delete that marks values as deleted but leaves them in the database.
err := collection.Delete(person, note)

// HARD DELETE all records that match the criteria, actually removing them from the database.
err := collection.HardDelete(criteria)

// QUERY many records from the database, by populating a slice of results
err := collection.Query(&people, criteria, options...)

// ITERATE over many records, using a more memory-friendly iterator to loop through a very large dataset.
it, err := collection.Iterator(criteria, options...)

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

There's a package for managing optional query arguments, such as sorting and row limits.  These options just encapsulate data.  It is the responsibility of each database adapter to implement each of these in its own query engine.

```go

// get a new iterator.  Sort results by name.  Return at most 100 rows.
it, err := collection.Iterator(criteria, option.SortAsc("name"), option.MaxRows(100))
```

**`SortAsc(fieldname)`** tells the database to sort by a particular field, in ascending order

**`SortDesc(fieldname)`** tells the database to sort by a particular field, in descending order

**`FirstRow()`** tells the database to return only the first matching row

**`MaxRows(count)`** tells the database to limit the number of records to the designated number of rows.

**`Fields(names...)`** tells the database to return only the named fields for each record

**`CaseSensitive(bool)`** tells the database whether string comparisons should consider letter case

## What matters here

This package contains **interfaces and data only** — there is no database logic. CRUD behavior lives entirely in the adapter packages ([data-mongo](https://github.com/benpate/data-mongo), [data-mock](https://github.com/benpate/data-mock), [data-slice](https://github.com/benpate/data-slice)). When changing this package, remember that every adapter must continue to satisfy the interfaces.

- **`option.FirstRow()` means "return only the first matching row" — it is not a pagination offset.** Despite the name, it carries no row number (`FirstRowOption` is an empty struct). Adapters translate it to a limit of one row (Mongo `SetLimit(1)`, slice `value[:1]`). Reaching for it to "skip to row N" will silently do the wrong thing.
- **Options only encapsulate intent; adapters decide what they mean.** An adapter is free to ignore an option it doesn't support, so a query that relies on `CaseSensitive` or `Fields` may behave differently across backends. Verify support in the specific adapter, not here.
- **`Delete` is a soft delete; `HardDelete` is permanent.** `Delete(object, note)` stamps the journal's `DeleteDate` and leaves the record in place; `HardDelete(criteria)` removes matching rows outright and takes no object or note.
- **The journal is the source of truth for object lifecycle.** Embedding `journal.Journal` satisfies all of `Object` except `ID()`. `SetUpdated`/`SetDeleted` bump `Revision` (the ETag); `SetCreated` does not. The `Revision` field serializes as `"signature"` for backward compatibility — don't rename the tag.

## Pull Requests Welcome

This library is a work in progress, and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making this library better, send in a pull request.  We're all in this together! 📚
