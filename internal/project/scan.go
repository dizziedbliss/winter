package project

import (
	"errors"
	"io/fs"
	"path/filepath"
)

func (d *Project) Scan() error {

	d.Logger.Printf("Searching for %s in Directory %s", d.Config, d.Path)

	err := filepath.WalkDir(d.Path, func(path string, dir fs.DirEntry, err error) error {

		if err != nil {
			d.DeployUI.Error(err)
			return err
		}

		if !dir.IsDir() && dir.Name() == d.Config {
			d.DeployUI.Status("Found " + d.Config + " in Directory " + path)
			d.ConfigPath = path
			return fs.SkipAll
		}

		return nil
	})

	if err != nil && err != fs.SkipAll {
		d.DeployUI.Error(err)
		return err
	}

	if d.ConfigPath == "" {
		d.DeployUI.Error(errors.New("config file not found"))
		return errors.New("config file not found")
	}

	return nil
}
