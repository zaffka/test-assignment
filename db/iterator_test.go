package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zaffka/test-assignment/db"
)

func TestIterate(t *testing.T) {
	testResetDB(t)

	type args struct {
		statFuncs []db.RecHandleFunc
	}
	score := 0
	tests := []struct {
		name      string
		args      args
		precondFn func()
		wantScore int
		wantErr   error
	}{
		{
			name:      "uninit db",
			args:      args{},
			precondFn: func() {},
			wantScore: 0,
			wantErr:   db.ErrDBCorrupted,
		},
		{
			name: "valid db",
			args: args{},
			precondFn: func() {
				err := db.Build(validDB)
				if err != nil {
					t.Fatal(err)
				}
			},
			wantScore: 0,
			wantErr:   nil,
		},
		{
			name: "valid with single stat func db",
			args: args{
				statFuncs: []db.RecHandleFunc{
					func(s []string) {
						if s[2] == "AdrianWilczynski" {
							score++
						}
					},
				},
			},
			precondFn: func() {
				err := db.Build(validDB)
				if err != nil {
					t.Fatal(err)
				}
			},
			wantScore: 2,
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.precondFn()
			err := db.Iterate(tt.args.statFuncs...)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantScore, score)
		})
	}
}

func testResetDB(t *testing.T) {
	t.Helper()

	db.Actors.Reader = nil
	db.Events.Reader = nil
	db.Repos.Reader = nil
	db.Commits.Reader = nil
}
