package deployment

import "fmt"
import "path/filepath"

func Deploy(opts DeploymentOpts) error {

	opts.Config = "Dockerfile"
	absPath, err := filepath.Abs(opts.Path)
	if err != nil {
		return err
	}
	 
	opts.Path = absPath

	//first it needs to scan for config file in the current directory and subdirectories so use Scan.go

	path, err := Scan(opts)

	opts.ConfigPath = path

	if err != nil {
		return err
	}

	// Now validate the file from the path and check if it is a valid Dockerfile


	err = Validate(opts)
	if err != nil {
		return err
	}

	err = Build(opts)
	if err != nil {
		return err
	}

	fmt.Println("Build Successfull")

	return nil
}
