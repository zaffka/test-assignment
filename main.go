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
	err := db.Build(dbFile)
	if err != nil {
		fmt.Printf("db opening failed with an error: %s\n", err)

		return
	}

	stat1 := &stat.Top{}
	_, err = db.Iterate(stat1.ActorsByCommitsAndPRs())
	if err != nil {
		fmt.Printf("failed to iterate over the DB: %s\n", err)
	}
}
