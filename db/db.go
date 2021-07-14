package db

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	_ "embed" // data file is hardcoded using embed functionality
	"encoding/csv"
	"errors"
	"fmt"
	"io"
)

//go:embed .data/data.tar.gz
var dbFile []byte

const preallocBufSize = 2048 // in bytes (256Kb)

// Build reads gzipped tar file with the database and initializes db.Events, db.Commits, db.Repos, db.Actors entities.
func Build() error {
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
			buf := bytes.NewBuffer(make([]byte, 0, preallocBufSize)) // allocating 256Kb buffer

			_, err := io.Copy(buf, tr)
			if err != nil {
				return fmt.Errorf("failed to read a file: %w", err)
			}

			csvRdr := csv.NewReader(buf)
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

	return nil
}
