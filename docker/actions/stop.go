package actions

import (
	"context"
	"github.com/docker/docker/client"
	"log"
)

func StopContainer(ctx context.Context, cli *client.Client, containerId string) {
	log.Printf("Stop container %s", containerId)
	if err := cli.ContainerStop(ctx, containerId, nil); err != nil {
		panic(err)
	}
	log.Printf("Succesfully stop container %s", containerId)
}
