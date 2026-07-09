package deployment

import (
	"fmt"
	"log"
	"os/exec"
)

func Build(file_path string) error {

	cmd := exec.Command("docker", "build", "--network=host", "-f", file_path, ".")

	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("build failed: %v\nOutput: %s", err, output)
	}

	log.Printf("build succeeded: %s", output)
	return nil
}
