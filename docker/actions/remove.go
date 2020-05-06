package actions

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

func RemoveContainer(ctx context.Context, cli *client.Client, containerId string) bool {
	var succeed bool

	log.Printf("Try to remove container %s", containerId)
	if err := cli.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{}); err != nil {
		log.Printf("Fail to remove container %s : %s", containerId, err)
		succeed = false
	} else {
		log.Printf("Succesfully remove container %s", containerId)
		succeed = true
	}
	return succeed
}
