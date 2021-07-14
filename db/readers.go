package db

import "encoding/csv"

const (
	commits = "data/commits.csv"
	repos   = "data/repos.csv"
	actors  = "data/actors.csv"
	events  = "data/events.csv"
)

type Reader *csv.Reader

var (
	Events  Reader
	Commits Reader
	Actors  Reader
	Repos   Reader
)
