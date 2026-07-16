package project

import (
	"fmt"
	"os"
	"os/exec"
)

func (d *Project) Build() error {

	d.UI.Status("Building Docker image from path: " + d.Path)

	cmd := exec.Command("docker", "build", "--network=host", "-f", d.ConfigPath, d.Path)

	if d.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()

	if err != nil {
		d.UI.Error(fmt.Errorf("build failed: %v", err))
		return fmt.Errorf("build failed: %v", err)
	}

	d.UI.Status("Build Successful")
	return nil
}
