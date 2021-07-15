package db

import (
	"errors"
	"fmt"
	"io"
)

var ErrBrokenDBLink = errors.New("link between DB entities is broken")

type StatFunc func([]string)

func Iterate(statFuncs ...StatFunc) ([]string, error) {
	err := dbValid()
	if err != nil {
		return nil, err
	}

	skipCSVheader := true
	lineNumber := 1
	for {
		event, err := Events.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		actor, err := Actors.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		repo, err := Repos.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		if skipCSVheader {
			skipCSVheader = false

			continue
		}

		err = replaceEventIDsByNames(event, actor, repo)
		if err != nil {
			return nil, fmt.Errorf("line number %d: %w", lineNumber, err)
		}

		for _, sfn := range statFuncs {
			sfn(event)
		}

		lineNumber++
	}

	return nil, nil
}

func replaceEventIDsByNames(event, actor, repo []string) error {
	if event[2] != actor[0] || event[3] != repo[0] {
		return ErrBrokenDBLink
	}

	event[2] = actor[1]
	event[3] = repo[1]

	return nil
}
