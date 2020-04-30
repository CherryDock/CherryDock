package actions

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

func StartContainer(ctx context.Context, cli *client.Client, containerId string) {

	log.Printf("Start container %s", containerId)

	if err := cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, containerId, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	log.Printf("Successfully tart container %s", containerId)
	log.Println(out)
}
