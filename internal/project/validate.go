package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (d *Project) Validate() error {
	d.UI.Status("Validating Dockerfile at path: " + d.Path)

	//for now lets just use docker --build --check
	contextDir := filepath.Dir(d.ConfigPath)

	cmd := exec.Command("docker", "build", "--check", "--build-arg", "BUILDKIT_DOCKERFILE_CHECK=error=true", "-f", d.ConfigPath, contextDir)

	if d.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()

	if err != nil {
		d.UI.Error(fmt.Errorf("validation failed:\n%s", err))
		return fmt.Errorf("validation failed:\n%s", err)
	}

	d.UI.Status("Dockerfile validation successful")

	return nil
}
