package main

import (
	"fmt"

	_ "embed"

	"github.com/zaffka/test-assignment/db"
	"github.com/zaffka/test-assignment/stat"
)

//go:embed data.tar.gz
var dbFile []byte

func main() {
	// parse datafile
	err := db.Build(dbFile)
	if err != nil {
		fmt.Printf("db opening failed with an error: %s\n", err)

		return
	}

	// init stat handlers
	actors := stat.New()
	repoCommits := stat.New()
	repoEvents := stat.New()

	// iterate over the database using stat handlers on each db record
	err = db.Iterate(
		actors.ActorsByCommitsAndPRs(),
		repoCommits.ReposByCommits(),
		repoEvents.ReposByWatchEvents(),
	)
	if err != nil {
		fmt.Printf("failed to iterate over the DB: %s\n", err)
	}

	// sorting, slicing and printing necessary results as a table

	top10repComm := repoCommits.List.SortedByScore()[:10]
	top10repComm.PrintTable(
		"Top 10 repositories by commits",
	)

	top10repEvents := repoEvents.List.SortedByScore()[:10]
	top10repEvents.PrintTable(
		"Top 10 repositories by watch events",
	)

	top10actors := actors.List.SortedByScore()[:10]
	top10actors.PrintTable(
		"Top 10 users by PR/Commits",
	)
}
