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
		assert.Nil(t, db.Commits.Reader)
		assert.Nil(t, db.Events.Reader)
		assert.Nil(t, db.Actors.Reader)
		assert.Nil(t, db.Repos.Reader)

		err := db.Build(invalidDB)
		assert.Equal(t, db.ErrDBCorrupted, err)
	})
	t.Run("valid db file", func(t *testing.T) {
		err := db.Build(validDB)
		assert.NoError(t, err)

		assert.NotNil(t, db.Commits.Reader)
		assert.NotNil(t, db.Events.Reader)
		assert.NotNil(t, db.Actors.Reader)
		assert.NotNil(t, db.Repos.Reader)
	})
}
