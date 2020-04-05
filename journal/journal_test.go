package journal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJournal(t *testing.T) {

	j := Journal{}

	// Test empty states
	assert.Equal(t, int64(0), j.CreateDate)
	assert.Equal(t, int64(0), j.UpdateDate)
	assert.Equal(t, int64(0), j.DeleteDate)
	assert.Equal(t, "", j.Note)
	assert.Equal(t, int64(0), j.Revision)
	assert.Equal(t, "0", j.ETag())
	assert.True(t, j.IsNew())
	assert.False(t, j.IsDeleted())

	// Test SetUpdated()
	createDate := time.Now().Unix()
	j.SetCreated("CREATED")
	assert.Equal(t, createDate, j.CreateDate)
	assert.Equal(t, createDate, j.UpdateDate)
	assert.Equal(t, int64(0), j.DeleteDate)
	assert.Equal(t, "CREATED", j.Note)
	assert.Equal(t, int64(0), j.Revision)
	assert.Equal(t, "0", j.ETag())
	assert.False(t, j.IsNew())
	assert.False(t, j.IsDeleted())

	updateDate := time.Now().Unix()
	j.SetUpdated("UPDATED")
	assert.Equal(t, createDate, j.CreateDate)
	assert.Equal(t, updateDate, j.UpdateDate)
	assert.Equal(t, int64(0), j.DeleteDate)
	assert.Equal(t, "UPDATED", j.Note)
	assert.Equal(t, int64(1), j.Revision)
	assert.Equal(t, "1", j.ETag())
	assert.False(t, j.IsNew())
	assert.False(t, j.IsDeleted())

	updatedAgainDate := time.Now().Unix()
	j.SetUpdated("")
	assert.Equal(t, createDate, j.CreateDate)
	assert.Equal(t, updatedAgainDate, j.UpdateDate)
	assert.Equal(t, int64(0), j.DeleteDate)
	assert.Equal(t, "UPDATED", j.Note)
	assert.Equal(t, int64(2), j.Revision)
	assert.Equal(t, "2", j.ETag())
	assert.False(t, j.IsNew())
	assert.False(t, j.IsDeleted())

	deletedDate := time.Now().Unix()
	j.SetDeleted("DELETED")
	assert.Equal(t, createDate, j.CreateDate)
	assert.Equal(t, updatedAgainDate, j.UpdateDate)
	assert.Equal(t, deletedDate, j.DeleteDate)
	assert.Equal(t, "DELETED", j.Note)
	assert.Equal(t, int64(3), j.Revision)
	assert.Equal(t, "3", j.ETag())
	assert.False(t, j.IsNew())
	assert.True(t, j.IsDeleted())
}
