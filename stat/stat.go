package stat

import (
	"github.com/zaffka/test-assignment/db"
)

// New is a constructor func returning a stat.Handler with preinitialized data fields.
func New() *Handler {
	return &Handler{
		List:   make(Records, 0, 100),
		Scores: make(map[string]*Record, 100),
	}
}

// Handler is struct holding the dependencies to collect the data.
// This data is using to collect stats executing specific stat funcs (db.RecHandleFunc).
type Handler struct {
	List   Records
	Scores map[string]*Record
}

// ActorsByCommitsAndPRs is a func responsible for collecting data about actors.
// It is handling only records scoring by pull-request or commit event type.
func (top *Handler) ActorsByCommitsAndPRs() db.RecHandleFunc {
	return func(event []string) { //TODO: use event as a type
		name := event[2]
		if eventType(event[1]).IsCommitOrPR() {
			record, exist := top.Scores[name]
			if exist {
				record.Score++
			} else {
				r := &Record{Name: name, Score: 1}
				top.Scores[name] = r
				top.List = append(top.List, r)
			}
		}
	}
}

// ReposByCommits is a func responsible for collecting data about repositories.
// It is handling only records scoring by commits number.
func (top *Handler) ReposByCommits() db.RecHandleFunc {
	return func(event []string) {
		repoName := event[3]
		if eventType(event[1]).IsCommit() {
			record, exist := top.Scores[repoName]
			if exist {
				record.Score++
			} else {
				r := &Record{Name: repoName, Score: 1}
				top.Scores[repoName] = r
				top.List = append(top.List, r)
			}
		}
	}
}

// ReposByWatchEvents is a func responsible for collecting data about repositories.
// It is handling only records scoring by watch events number.
func (top *Handler) ReposByWatchEvents() db.RecHandleFunc {
	return func(event []string) {
		repoName := event[3]
		if eventType(event[1]).IsWatch() {
			record, exist := top.Scores[repoName]
			if exist {
				record.Score++
			} else {
				r := &Record{Name: repoName, Score: 1}
				top.Scores[repoName] = r
				top.List = append(top.List, r)
			}
		}
	}
}
