package deployment

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Build(opts DeploymentOpts) error {

	cmd := exec.Command("docker", "build", "--network=host", "-f", opts.ConfigPath, ".")

	if opts.Verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = nil
		cmd.Stderr = nil
	}

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("build failed: %v", err)
	}

	log.Printf("build succeeded")
	return nil
}
