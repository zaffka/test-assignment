package main

import (
	"fmt"

	"github.com/zaffka/test-assignment/db"
)

var version = "dev"

func main() {
	fmt.Printf("app:%s:opening database\n", version)

	err := db.Build()
	if err != nil {
		fmt.Printf("db opening failed with an error: %s", err)
	}
}
