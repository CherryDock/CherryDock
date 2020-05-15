package actions

import (
	"context"
	"github.com/docker/docker/client"
	"log"
)

func PauseContainer(ctx context.Context, cli *client.Client, containerId string) bool {
	var succeed bool

	log.Printf("Try to pause container %s", containerId)
	if err := cli.ContainerPause(ctx, containerId); err != nil {
		log.Printf("Fail to pause container %s : %s", containerId, err)
		succeed = false
	} else {
		log.Printf("Successfully pause container %s", containerId)
		succeed = true
	}
	return succeed
}
