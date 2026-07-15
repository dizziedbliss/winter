package project

import (
	"fmt"
	"path/filepath"
)

func Deploy(p *Project) error {

	p.Config = "Dockerfile"
	absPath, err := filepath.Abs(p.Path)
	if err != nil {
		p.DeployUI.Error(err)
		return err
	}

	p.Path = absPath

	//first it needs to scan for config file in the current directory and subdirectories so use Scan.go

	if err := p.Scan(); err != nil {
		p.DeployUI.Error(err)
		return fmt.Errorf("Scan failed: %v", err)
	}

	// Now validate the file from the path and check if it is a valid Dockerfile

	if err := p.Validate(); err != nil {
		p.DeployUI.Error(err)
		return fmt.Errorf("Validation failed: %v", err)
	}

	// Now build the docker image from the path and config file

	if err := p.Build(); err != nil {
		p.DeployUI.Error(err)
		return fmt.Errorf("Build failed: %v", err)
	}

	p.DeployUI.Success("Build Successful")

	return nil
}
