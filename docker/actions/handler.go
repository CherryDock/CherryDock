package actions

import (
	"context"
	"github.com/docker/docker/client"
)

func Handle(action string, containerId string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	switch action {
	case "start":
		StartContainer(ctx, cli, containerId)
	case "stop":
		StopContainer(ctx, cli, containerId)
	default:
		panic("Unknow Action")
	}

}
