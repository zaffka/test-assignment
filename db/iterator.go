package db

import (
	"errors"
	"fmt"
	"io"
)

// ErrBrokenDBLink is returning if we have no id-alignment between db.Events and db.Actors/db.Repos.
var ErrBrokenDBLink = errors.New("link between DB entities is broken")

// RecHandleFunc represents any func to be called on every single db record.
type RecHandleFunc func([]string)

// Iterate cycling thru the database and executes any number of db.RecHandleFuncs on each db row.
func Iterate(recHandleFuncs ...RecHandleFunc) error {
	err := dbValid()
	if err != nil {
		return err
	}

	skipCSVheader := true
	lineNumber := 1
	for {
		event, err := Events.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		actor, err := Actors.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		repo, err := Repos.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		if skipCSVheader {
			skipCSVheader = false

			continue
		}

		err = replaceEventIDsByNames(event, actor, repo)
		if err != nil {
			return fmt.Errorf("line number %d: %w", lineNumber, err)
		}

		for _, rfn := range recHandleFuncs {
			rfn(event)
		}

		lineNumber++
	}

	return nil
}

func replaceEventIDsByNames(event, actor, repo []string) error {
	if event[2] != actor[0] || event[3] != repo[0] {
		return ErrBrokenDBLink
	}

	event[2] = actor[1]
	event[3] = repo[1]

	return nil
}
