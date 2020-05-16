package actions

import (
	"context"
	"github.com/CherryDock/CherryDock/api/docker/monitoring"
	"github.com/CherryDock/CherryDock/api/jsonutils"
	"github.com/docker/docker/client"
	"log"
)

func ActionSingleContainer(singleAction action, containerId string) bool {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		log.Println("Fail to create docker client")
	}

	var succeed bool
	err = singleAction(ctx, cli, containerId)

	if err != nil {
		succeed = false
	} else {
		succeed = true
	}

	return succeed
}

func ActionOnAllContainer(singleAction action, all bool) []byte {
	ctx := context.Background()
	cli, err := client.NewEnvClient()

	if err != nil {
		log.Println("Fail to create docker client")
	}

	var states []actionState
	var succeed bool

	allContainers := monitoring.ContainersId(all)
	for _, id := range allContainers {
		err = singleAction(ctx, cli, id)
		if err != nil {
			succeed = false
		} else {
			succeed = true
		}
		states = append(states, actionState{id, succeed})
	}

	return jsonutils.FormatToJson(states)
}
