package deployment

import "fmt"
import "path/filepath"

func Deploy(directory string) error {

	file_type := "Dockerfile"
	directory, err := filepath.Abs(directory)
	if err != nil {
		return err
	}

	//first it needs to scan for config file in the current directory and subdirectories so use Scan.go

	path, err := Scan(directory, file_type)

	if err != nil {
		return err
	}

	// Now validate the file from the path and check if it is a valid Dockerfile


	err = Validate(path)
	if err != nil {
		return err
	}

	err = Build(path)
	if err != nil {
		return err
	}

	fmt.Println("Build Successfull")

	return nil
}
