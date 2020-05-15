package actions

import (
	"context"
	"github.com/docker/docker/client"
	"log"
)

func KillContainer(ctx context.Context, cli *client.Client, containerId string) bool {
	var succeed bool

	log.Printf("Try to kill container %s", containerId)
	if err := cli.ContainerKill(ctx, containerId, ""); err != nil {
		log.Printf("Fail to kill container %s : %s", containerId, err)
		succeed = false
	} else {
		log.Printf("Succesfully kill container %s", containerId)
		succeed = true
	}
	return succeed
}
