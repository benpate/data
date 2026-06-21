package data

// TransactionCallbackFunc is a function that is called after a
// transaction session is created.  The callback function can return a
// value (which is passed through) and an error.  If the error is nil
// then the transaction is committed to the database.  Otherwise,
// the transaction is rolled back.
type TransactionCallbackFunc func(Session) (any, error)
