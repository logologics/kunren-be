package integration

import (
	"context"
	"fmt"
	"log"
	"testing"
)

// expects DB empty
func TestStore(t *testing.T) {
	deleteAllWordsFromDockerCli(t)

}

func deleteAllWordsFromDockerCli(t *testing.T) {
	ctx := context.Background()
	execID, err := DockerExec(ctx, "some-mongo", []string{"mongo", "/scripts/delete-words.js"})
	if err != nil {
		log.Fatal(err)
	}
	execRes, err := DockerInspectExecResp(ctx, execID.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res |%v|\n", execRes)
	t.Errorf("hi")
}
