// Package journal provides a simple implementation of a change-tracking
// "journal" that can be embedded into other data objects to track their
// creation, update, and deletion history.
package journal

import (
	"strconv"
	"time"
)

// Journal tracks a summary of changes to an object over time.  Journal implements
// *most* of the data.Object interface (except for the ID function) right out of
// the box, so it's a useful example for implementing the data.Object interface,
// or even to embed directly into an existing model object.
type Journal struct {
	CreateDate int64  `path:"createDate" json:"createDate" bson:"createDate"` // CreateDate is the Unix epoch (milliseconds) that the attached object was created
	UpdateDate int64  `path:"updateDate" json:"updateDate" bson:"updateDate"` // UpdateDate is the Unix epoch (milliseconds) that the attached object was last updated
	DeleteDate int64  `path:"deleteDate" json:"deleteDate" bson:"deleteDate"` // DeleteDate is the Unix epoch (milliseconds) that the attached object was deleted
	Note       string `path:"note"       json:"note"       bson:"note"`       // Note is a human-readable description of the change that was made
	Revision   int64  `path:"signature"  json:"signature"  bson:"signature"`  // Signature is a unique identifier for this version of the object
}

// IsNew returns TRUE if the object has not yet been saved to the database
func (journal Journal) IsNew() bool {
	return (journal.CreateDate == 0)
}

// IsDeleted returns TRUE if the object has been "virtually deleted" from the database
func (journal Journal) IsDeleted() bool {
	return (journal.DeleteDate > 0)
}

// Created returns the Unix epoch (in milliseconds) when the object containing this journal was created
func (journal Journal) Created() int64 {
	return journal.CreateDate
}

// Updated returns the Unix epoch (in milliseconds) when the object containing this journal was updated
func (journal Journal) Updated() int64 {
	return journal.UpdateDate
}

// SetCreated must be called whenever a new object is added to the database
func (journal *Journal) SetCreated(note string) {

	timestamp := time.Now().UnixMilli()
	journal.CreateDate = timestamp
	journal.UpdateDate = timestamp

	if note != "" {
		journal.Note = note
	}
}

// SetUpdated must be called whenever an existing object is updated in the database
func (journal *Journal) SetUpdated(note string) {

	journal.UpdateDate = time.Now().UnixMilli()
	journal.Revision = journal.Revision + 1

	if note != "" {
		journal.Note = note
	}
}

// SetDeleted must be called to "virtual-delete" an object in the database
func (journal *Journal) SetDeleted(note string) {

	timestamp := time.Now().UnixMilli()
	journal.UpdateDate = timestamp
	journal.DeleteDate = timestamp
	journal.Revision = journal.Revision + 1

	if note != "" {
		journal.Note = note
	}
}

// ETag returns the signature for this object.  It currently returns the "revision number"
// which should be fine unless we make out-of-band updates to objects.
func (journal Journal) ETag() string {
	return strconv.FormatInt(journal.Revision, 10)
}
