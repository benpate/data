package journal

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJournal_IsNew(t *testing.T) {

	// A zero-value journal has never been created, so it is "new"
	journal := Journal{}
	assert.True(t, journal.IsNew())

	// Once a CreateDate is set, the journal is no longer "new"
	journal.CreateDate = 1
	assert.False(t, journal.IsNew())
}

func TestJournal_IsDeleted(t *testing.T) {

	// A zero-value journal has not been deleted
	journal := Journal{}
	assert.False(t, journal.IsDeleted())

	// A positive DeleteDate marks the journal as deleted
	journal.DeleteDate = 1
	assert.True(t, journal.IsDeleted())

	// Negative dates are not considered "deleted" (boundary check)
	journal.DeleteDate = -1
	assert.False(t, journal.IsDeleted())
}

func TestJournal_Created(t *testing.T) {

	journal := Journal{CreateDate: 12345}
	assert.Equal(t, int64(12345), journal.Created())
}

func TestJournal_Updated(t *testing.T) {

	journal := Journal{UpdateDate: 67890}
	assert.Equal(t, int64(67890), journal.Updated())
}

func TestJournal_SetCreated(t *testing.T) {

	before := time.Now().UnixMilli()
	journal := Journal{}
	journal.SetCreated("created this object")
	after := time.Now().UnixMilli()

	// CreateDate and UpdateDate are both stamped with "now"
	assert.GreaterOrEqual(t, journal.CreateDate, before)
	assert.LessOrEqual(t, journal.CreateDate, after)
	assert.Equal(t, journal.CreateDate, journal.UpdateDate)

	// The note is recorded, and the object is no longer "new"
	assert.Equal(t, "created this object", journal.Note)
	assert.False(t, journal.IsNew())
}

// An empty note passed to SetCreated must not overwrite an existing note
func TestJournal_SetCreated_EmptyNote(t *testing.T) {

	journal := Journal{Note: "original note"}
	journal.SetCreated("")

	assert.Equal(t, "original note", journal.Note)
	assert.NotZero(t, journal.CreateDate)
}

func TestJournal_SetUpdated(t *testing.T) {

	before := time.Now().UnixMilli()
	journal := Journal{CreateDate: 1, Revision: 5}
	journal.SetUpdated("updated this object")
	after := time.Now().UnixMilli()

	// UpdateDate is stamped with "now", but CreateDate is untouched
	assert.GreaterOrEqual(t, journal.UpdateDate, before)
	assert.LessOrEqual(t, journal.UpdateDate, after)
	assert.Equal(t, int64(1), journal.CreateDate)

	// The Revision is incremented and the note is recorded
	assert.Equal(t, int64(6), journal.Revision)
	assert.Equal(t, "updated this object", journal.Note)
}

// An empty note passed to SetUpdated must not overwrite an existing note,
// but the Revision must still be incremented.
func TestJournal_SetUpdated_EmptyNote(t *testing.T) {

	journal := Journal{Note: "original note", Revision: 2}
	journal.SetUpdated("")

	assert.Equal(t, "original note", journal.Note)
	assert.Equal(t, int64(3), journal.Revision)
}

func TestJournal_SetDeleted(t *testing.T) {

	before := time.Now().UnixMilli()
	journal := Journal{CreateDate: 1, Revision: 9}
	journal.SetDeleted("deleted this object")
	after := time.Now().UnixMilli()

	// Both UpdateDate and DeleteDate are stamped with the same "now" value
	assert.GreaterOrEqual(t, journal.DeleteDate, before)
	assert.LessOrEqual(t, journal.DeleteDate, after)
	assert.Equal(t, journal.UpdateDate, journal.DeleteDate)

	// The Revision is incremented, the note is recorded, and the object reads as deleted
	assert.Equal(t, int64(10), journal.Revision)
	assert.Equal(t, "deleted this object", journal.Note)
	assert.True(t, journal.IsDeleted())
}

// An empty note passed to SetDeleted must not overwrite an existing note,
// but the object must still be marked deleted and the Revision incremented.
func TestJournal_SetDeleted_EmptyNote(t *testing.T) {

	journal := Journal{Note: "original note", Revision: 0}
	journal.SetDeleted("")

	assert.Equal(t, "original note", journal.Note)
	assert.Equal(t, int64(1), journal.Revision)
	assert.True(t, journal.IsDeleted())
}

func TestJournal_ETag(t *testing.T) {

	journal := Journal{Revision: 42}
	assert.Equal(t, "42", journal.ETag())

	// A zero-value journal returns the string "0"
	assert.Equal(t, "0", Journal{}.ETag())
}

// The full create -> update -> delete lifecycle should leave the journal in a
// consistent state, exercising the accessors together.
func TestJournal_Lifecycle(t *testing.T) {

	journal := Journal{}
	require.True(t, journal.IsNew())

	journal.SetCreated("create")
	require.False(t, journal.IsNew())
	require.False(t, journal.IsDeleted())
	require.Equal(t, journal.Created(), journal.Updated())
	require.Equal(t, int64(0), journal.Revision)

	journal.SetUpdated("update")
	require.Equal(t, int64(1), journal.Revision)
	require.GreaterOrEqual(t, journal.Updated(), journal.Created())

	journal.SetDeleted("delete")
	require.True(t, journal.IsDeleted())
	require.Equal(t, int64(2), journal.Revision)
	require.Equal(t, strconv.FormatInt(journal.Revision, 10), journal.ETag())
}
