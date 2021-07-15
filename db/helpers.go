package db

func dbValid() error {
	if !Events.Valid() || !Commits.Valid() || !Actors.Valid() || !Repos.Valid() {
		return ErrDBCorrupted
	}

	return nil
}
