package deployment

import (
	"fmt"
	"os"
	"os/exec"
)

func (d *DeployOptions) Build() error {

	cmd := exec.Command("docker", "build", "--network=host", "-f", d.ConfigPath, d.Path)

	if d.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("build failed: %v", err)
	}

	d.Logger.Printf("build succeeded")
	return nil
}
