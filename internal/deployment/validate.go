package deployment

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (d *DeployOptions) Validate() error {
	d.Logger.Printf("Validating Dockerfile at path: %s", d.Path)

	//for now lets just use docker --build --check
	contextDir := filepath.Dir(d.ConfigPath)

	cmd := exec.Command("docker", "build", "--check", "--build-arg", "BUILDKIT_DOCKERFILE_CHECK=error=true", "-f", d.ConfigPath, contextDir)

	if d.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("validation failed:\n%s", err)
	}

	d.Logger.Printf("validation succeeded")

	return nil
}
