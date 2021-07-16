package stat

type eventType string

func (et eventType) IsCommitOrPR() bool {
	return et == pullRequest || et == commit
}

func (et eventType) IsCommit() bool {
	return et == commit
}

func (et eventType) IsWatch() bool {
	return et == watch
}

const (
	commit      eventType = "PushEvent"
	pullRequest eventType = "PullRequestEvent"
	watch       eventType = "WatchEvent"
)
