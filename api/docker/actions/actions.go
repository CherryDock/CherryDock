package actions

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type action func(ctx context.Context, cli *client.Client, containerId string) error

type actionState struct {
	containerId string
	succeed     bool
}

func KillContainer(ctx context.Context, cli *client.Client, containerId string) error {
	err := cli.ContainerKill(ctx, containerId, "")
	return err
}

func PauseContainer(ctx context.Context, cli *client.Client, containerId string) error {
	err := cli.ContainerPause(ctx, containerId)
	return err
}

func RemoveContainer(ctx context.Context, cli *client.Client, containerId string) error {
	err := cli.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{})
	return err
}

func RestartContainer(ctx context.Context, cli *client.Client, containerId string) error {
	err := cli.ContainerRestart(ctx, containerId, nil)
	return err
}

func StartContainer(ctx context.Context, cli *client.Client, containerId string) error {
	err := cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{})
	return err
}

func StopContainer(ctx context.Context, cli *client.Client, containerId string) error {
	err := cli.ContainerStop(ctx, containerId, nil)
	return err
}

func UnpauseContainer(ctx context.Context, cli *client.Client, containerId string) error {
	err := cli.ContainerUnpause(ctx, containerId)
	return err
}
