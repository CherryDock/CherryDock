package actions

import (
	"context"
	"github.com/docker/docker/client"
	"log"
)

func UnpauseContainer(ctx context.Context, cli *client.Client, containerId string) bool {
	var succeed bool

	log.Printf("Try to unpause tcontainer %s", containerId)
	if err := cli.ContainerUnpause(ctx, containerId); err != nil {
		log.Printf("Fail to unpause container %s : %s", containerId, err)
		succeed = false
	} else {
		log.Printf("Successfully unpause container %s", containerId)
		succeed = true
	}
	return succeed
}
