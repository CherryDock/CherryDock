package actions

import (
	"context"
	"github.com/docker/docker/client"
	"log"
)

func StopContainer(ctx context.Context, cli *client.Client, containerId string) bool {
	var succeed bool

	log.Printf("Try to stop container %s", containerId)
	if err := cli.ContainerStop(ctx, containerId, nil); err != nil {
		log.Printf("Fail to stop container %s : %s", containerId, err)
		succeed = false
	} else {
		log.Printf("Succesfully stop container %s", containerId)
		succeed = true
	}
	return succeed
}
