package deployment

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func Validate(file_path string) error {
	log.Printf("Validating Dockerfile at path: %s", file_path)

	//for now lets just use docker --build --check
	contextDir := filepath.Dir(file_path)

	cmd := exec.Command("docker", "build", "--check", "--build-arg", "BUILDKIT_DOCKERFILE_CHECK=error=true", "-f", file_path, contextDir)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("validation failed:\n%s", output)
	}

	log.Printf("validation succeeded: %s", output)

	return nil
}
