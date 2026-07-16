package project

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
)

func (d *Project) Scan() error {

	d.UI.Status("Searching for " + d.Config + " in Directory " + d.Path)

	err := filepath.WalkDir(d.Path, func(path string, dir fs.DirEntry, err error) error {

		if err != nil {
			d.UI.Error(err)
			return err
		}

		if !dir.IsDir() && dir.Name() == d.Config {
			d.UI.Status(fmt.Sprintf("Found %s in Directory %s", d.Config, path))
			d.ConfigPath = path
			return fs.SkipAll
		}

		return nil
	})

	if err != nil && err != fs.SkipAll {
		d.UI.Error(err)
		return err
	}

	if d.ConfigPath == "" {
		d.UI.Error(errors.New("config file not found"))
		return errors.New("config file not found")
	}

	return nil
}
