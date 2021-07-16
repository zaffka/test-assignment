package db

import (
	"encoding/csv"
	"errors"
)

// ErrDBCorrupted is an error to be returned after db.Build call in a case of incomplete reader entities init.
var ErrDBCorrupted = errors.New("database file corrupted or has less entities")

// All the files expected to be in a data.tar.gz file.
const (
	commits = "data/commits.csv"
	repos   = "data/repos.csv"
	actors  = "data/actors.csv"
	events  = "data/events.csv"
)

// Package scoped enitities for all the files backed the database.
var (
	Events  Entity
	Commits Entity
	Actors  Entity
	Repos   Entity
)

// Entity is a struct to hold a csv.Reader for every file inside of the database.
type Entity struct {
	Reader *csv.Reader
}

// Read is a wrapper on the db.Entity.Reader of *csv.Reader type.
func (e *Entity) Read() ([]string, error) {
	return e.Reader.Read()
}

// Valid checks if the db.Entity.Reader is nil.
func (e *Entity) Valid() bool {
	return e.Reader != nil
}
