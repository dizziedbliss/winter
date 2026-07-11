package deployment

import (
	"errors"
	"io/fs"
	"log"
	"path/filepath"
)

func Scan(opts DeploymentOpts) (string, error) {

	log.Printf("Searching for %s in Directory %s", opts.Config, opts.Path)

	var found string

	err := filepath.WalkDir(opts.Path, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !d.IsDir() && d.Name() == opts.Config {
			log.Printf("Found %s in Directory %s", opts.Config, path)
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
