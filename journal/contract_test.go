package journal_test

import (
	"testing"

	"github.com/benpate/data"
	"github.com/benpate/data/journal"
	"github.com/stretchr/testify/require"
)

// personFixture is a minimal domain object, built the way the README recommends:
// embed journal.Journal, then supply the only method it does not provide.
type personFixture struct {
	PersonID string
	journal.Journal
}

// ID returns the unique identifier for this fixture
func (person personFixture) ID() string {
	return person.PersonID
}

// RULE: Journal's mutators use pointer receivers, so only a POINTER to an
// embedding type satisfies data.Object. Asserting the value form here would
// not compile. Adapters must hold pointers.
var _ data.Object = &personFixture{}

// TestJournal_SatisfiesObject pins the contract that an embedded Journal supplies
// every data.Object method except ID
func TestJournal_SatisfiesObject(t *testing.T) {

	// Drive the fixture strictly through the data.Object interface
	var object data.Object = &personFixture{PersonID: "123"}

	require.True(t, object.IsNew())
	require.Equal(t, "123", object.ID())

	// Creation stamps the dates but leaves Revision (and therefore ETag) at zero
	object.SetCreated("created")
	require.False(t, object.IsNew())
	require.False(t, object.IsDeleted())
	require.Equal(t, object.Created(), object.Updated())
	require.Equal(t, "0", object.ETag())

	// Every subsequent write advances the ETag, which adapters use for optimistic concurrency
	object.SetUpdated("updated")
	require.Equal(t, "1", object.ETag())

	// A virtual delete leaves the object readable, but marked
	object.SetDeleted("deleted")
	require.True(t, object.IsDeleted())
	require.Equal(t, "2", object.ETag())
}
