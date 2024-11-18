package data

// Object interface defines all of the methods required for the `data`
// library to create, read, update, and delete objects in the database.
type Object interface {

	// The unique identifier for this object
	ID() string

	// Unix epoch time when this object was created
	Created() int64

	// Unix epoch time when this object was updated
	Updated() int64

	// IsNew returns TRUE if the object has not yet been saved to the database
	IsNew() bool

	// IsDeleted returns TRUE if the object has been virtually deleted
	IsDeleted() bool

	// SetCreated stamps the CreateDate and UpdateDate of the object, and makes a note
	SetCreated(comment string)

	// SetUpdated stamps the UpdateDate of the object, and makes a note
	SetUpdated(comment string)

	// SetDeleted marks the object virtually "deleted", and makes a note
	SetDeleted(comment string)

	// ETag returns the signature or revision number of the object
	ETag() string
}
