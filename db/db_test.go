package db_test

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
	"github.com/zaffka/test-assignment/db"
)

//go:embed .test_data/valid.tar.gz
var validDB []byte

//go:embed .test_data/invalid.tar.gz
var invalidDB []byte

func TestDB(t *testing.T) {
	t.Run("invalid db file", func(t *testing.T) {
		assert.Nil(t, db.Commits)
		assert.Nil(t, db.Events)
		assert.Nil(t, db.Actors)
		assert.Nil(t, db.Repos)

		err := db.Build(invalidDB)
		assert.Equal(t, db.ErrDBCorrupted, err)
	})
	t.Run("valid db file", func(t *testing.T) {
		err := db.Build(validDB)
		assert.NoError(t, err)

		assert.NotNil(t, db.Commits)
		assert.NotNil(t, db.Events)
		assert.NotNil(t, db.Actors)
		assert.NotNil(t, db.Repos)
	})
}
