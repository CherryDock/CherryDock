package actions

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

func StartContainer(ctx context.Context, cli *client.Client, containerId string) bool {
	var succeed bool

	log.Printf("Try to start container %s", containerId)
	if err := cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err != nil {
		log.Printf("Fail to start container %s : %s",containerId,err)
		succeed = false
	} else {
		log.Printf("Successfully start container %s", containerId)
		succeed = true
	}
	return succeed
}

