package deployment

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Validate(opts DeploymentOpts) error {
	log.Printf("Validating Dockerfile at path: %s", opts.Path)

	//for now lets just use docker --build --check
	contextDir := filepath.Dir(opts.ConfigPath)

	cmd := exec.Command("docker", "build", "--check", "--build-arg", "BUILDKIT_DOCKERFILE_CHECK=error=true", "-f", opts.ConfigPath, contextDir)

	if opts.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = nil
		cmd.Stderr = nil
	}
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("validation failed:\n%s", err)
	}

	log.Printf("validation succeeded")

	return nil
}
