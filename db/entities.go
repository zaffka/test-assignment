package db

import (
	"encoding/csv"
	"errors"
)

// ErrDBCorrupted is an error to be returned after db.Build call in a case of incomplete reader entities init.
var ErrDBCorrupted = errors.New("database file corrupted or has less entities")

const (
	commits = "data/commits.csv"
	repos   = "data/repos.csv"
	actors  = "data/actors.csv"
	events  = "data/events.csv"
)

var (
	Events  *Entity
	Commits *Entity
	Actors  *Entity
	Repos   *Entity
)

type Entity struct {
	Reader *csv.Reader
}

func (e *Entity) Read() ([]string, error) {
	return e.Reader.Read()
}

func (e *Entity) Valid() bool {
	return e.Reader != nil
}
