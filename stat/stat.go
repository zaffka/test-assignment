package stat

import (
	"fmt"

	"github.com/zaffka/test-assignment/db"
)

type eventType string

func (et eventType) IsCommitOrPR() bool {
	return et == pullRequest || et == commit
}

const (
	commit      eventType = "PushEvent"
	pullRequest eventType = "PullRequestEvent"
)

type Top []string

func (top Top) ActorsByCommitsAndPRs() db.StatFunc {
	return func(event []string) { //TODO: convert event to a type
		if eventType(event[1]).IsCommitOrPR() {
			fmt.Println(event)
		}
	}
}
