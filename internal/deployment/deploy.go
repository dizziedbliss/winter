package deployment

import "path/filepath"
import "fmt"

func Deploy(opts *DeployOptions) error {

	opts.Config = "Dockerfile"
	absPath, err := filepath.Abs(opts.Path)
	if err != nil {
		return err
	}

	opts.Path = absPath

	//first it needs to scan for config file in the current directory and subdirectories so use Scan.go

	if err := opts.Scan(); err != nil {
		return fmt.Errorf("Scan failed: %v", err)
	}

	// Now validate the file from the path and check if it is a valid Dockerfile

	if err := opts.Validate(); err != nil {
		return fmt.Errorf("Validation failed: %v", err)
	}

	// Now build the docker image from the path and config file

	if err := opts.Build(); err != nil {
		return fmt.Errorf("Build failed: %v", err)
	}

	opts.Logger.Println("Build Successfull")

	return nil
}
