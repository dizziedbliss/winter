package project

import (
	"fmt"
	"os"
	"os/exec"
)

func (d *Project) Build() error {

	cmd := exec.Command("docker", "build", "--network=host", "-f", d.ConfigPath, d.Path)

	if d.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()

	if err != nil {
		d.DeployUI.Error(fmt.Errorf("build failed: %v", err))
		return fmt.Errorf("build failed: %v", err)
	}

	d.DeployUI.Success("Build Successful")
	return nil
}
