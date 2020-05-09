package actions

import (
	"context"
	"github.com/CherryDock/CherryDock/api/docker/monitoring"
	"github.com/CherryDock/CherryDock/api/jsonutils"
	"github.com/docker/docker/client"
	"log"
)

type action func(ctx context.Context, cli *client.Client, containerId string) bool

type ActionState struct {
	ContainerId string
	Succeed     bool
}

func ActionSingleContainer(singleAction action, containerdId string) bool {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		log.Println("Fail to create docker client")
	}

	succeed := singleAction(ctx, cli, containerdId)

	return succeed
}

func ActionOnAllContainer(singleAction action, all bool) []byte {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		log.Println("Fail to create docker client")
	}

	var states []ActionState

	allContainers := monitoring.ContainersId(all)
	for _, id := range allContainers {
		succeed := singleAction(ctx, cli, id)
		states = append(states, ActionState{id, succeed})
	}

	return jsonutils.FormatToJson(states)
}
