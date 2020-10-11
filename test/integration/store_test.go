package integration

import (
	"os"
	"os/exec"
	"testing"
)

// expects DB empty
func TestStore(t *testing.T) {
	deleteAllWordsFromDocker(t)
}

// docker exec -ti some-mongo mongo /scripts/delete-words.js
func deleteAllWordsFromDocker(t *testing.T) {
	dockerExecutable, _ := exec.LookPath("docker")

	cmd := &exec.Cmd{
		Path:   dockerExecutable,
		Args:   []string{dockerExecutable, "exec", "some-mongo", "mongo","/scripts/delete-words.js"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// run `go version` command
	if err := cmd.Run(); err != nil {
		t.Errorf("Err %v", err)
	}
}
