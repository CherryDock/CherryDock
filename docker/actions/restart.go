package actions

import (
	"context"
	"github.com/docker/docker/client"
	"log"
)

func RestartContainer(ctx context.Context, cli *client.Client, containerId string) bool {
	var succeed bool

	log.Printf("Try to restart container %s", containerId)
	if err := cli.ContainerRestart(ctx, containerId, nil); err != nil {
		log.Printf("Fail to restart container %s : %s", containerId, err)
		succeed = false
	} else {
		log.Printf("Succesfully restart container %s", containerId)
		succeed = true
	}
	return succeed
}
