package data

// Object interface defines all of the methods that a Domain Object must provide to Presto
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

	// Get implements the path.Getter interface.
	// It returns a value stored in the object.
	// Dotted values reference sub-objects.
	GetPath(string) (interface{}, bool)

	// SetPath implements the path.Setter interface.
	// It applies a value to the object's properties.
	// Dotted values reference sub-objects.
	SetPath(string, interface{}) error
}
