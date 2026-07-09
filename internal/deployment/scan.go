package deployment

import (
	"errors"
	"io/fs"
	"log"
	"path/filepath"
)

func Scan(directory string, config string) (string, error) {

	log.Printf("Searching for %s in Directory %s", config, directory)

	var found string

	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !d.IsDir() && d.Name() == config {
			log.Printf("Found %s in Directory %s", config, path)
			found = path
			return fs.SkipAll
		}

		return nil
	})


	if err != nil && err != fs.SkipAll {
		return "", err
	}

	if found == "" {
		return "", errors.New("config file not found")
	}

	return found, nil
}
