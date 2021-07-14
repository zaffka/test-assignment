package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zaffka/test-assignment/db"
)

func TestDBReading(t *testing.T) {
	assert.Nil(t, db.Commits)
	assert.Nil(t, db.Events)
	assert.Nil(t, db.Actors)
	assert.Nil(t, db.Repos)

	err := db.Build()
	assert.NoError(t, err)

	assert.NotNil(t, db.Commits)
	assert.NotNil(t, db.Events)
	assert.NotNil(t, db.Actors)
	assert.NotNil(t, db.Repos)
}
