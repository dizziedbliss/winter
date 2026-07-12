package deployment

import (
	"errors"
	"io/fs"
	"path/filepath"
)

func (d *DeployOptions)Scan() ( error) {

	d.Logger.Printf("Searching for %s in Directory %s", d.Config, d.Path)

	err := filepath.WalkDir(d.Path, func(path string, dir fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !dir.IsDir() && dir.Name() == d.Config {
			d.Logger.Printf("Found %s in Directory %s", d.Config, path)
			d.ConfigPath = path
			return fs.SkipAll
		}

		return nil
	})


	if err != nil && err != fs.SkipAll {
		return err
	}

	if d.ConfigPath == "" {
		return errors.New("config file not found")
	}

	return nil
}
