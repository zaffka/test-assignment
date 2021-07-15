package db

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
)

// Size in bytes (256Kb) for the initial buffer to where the entities from the tar.gz file read.
const preallocBufSize = 2048

// Build accepts gzipped tar file as an argument initializes readers: db.Events, db.Commits, db.Repos, db.Actors.
func Build(dbFile []byte) error {
	gr, err := gzip.NewReader(bytes.NewBuffer(dbFile))
	if err != nil {
		return err
	}

	tr := tar.NewReader(gr)

	for {
		header, err := tr.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return fmt.Errorf("database file corrupted: %w", err)
		}

		if header.Typeflag == tar.TypeReg {
			buf := bytes.NewBuffer(make([]byte, 0, preallocBufSize)) // preallocating 256Kb

			_, err := io.Copy(buf, tr)
			if err != nil {
				return fmt.Errorf("failed to read a file: %w", err)
			}

			csvRdr := &Entity{Reader: csv.NewReader(buf)}
			switch header.Name {
			case events:
				Events = csvRdr
			case commits:
				Commits = csvRdr
			case actors:
				Actors = csvRdr
			case repos:
				Repos = csvRdr
			}
		}
	}

	return dbValid()
}
