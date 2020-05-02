package monitoring

import (
	"context"
	"github.com/Fszta/DockerMonitoring/jsonutils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

type containerInfo struct {
	Name []string
	Status string
	Id string
	Image string
	Command string
	Created int64
	Ports []types.Port
}

func GetContainersInfo(allContainers bool) []byte{
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Println(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: allContainers})

	var containersInfo []containerInfo
	for _, container := range containers {
		containersInfo = append(
			containersInfo, containerInfo{
			container.Names,
			container.Status,
			container.ID[:10],
			container.Image,
			container.Command,
			container.Created,
			container.Ports})
	}

	return jsonutils.FormatToJson(containersInfo)
}

func ContainersId(all bool) []string{
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	runningContainers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: all})

	var containersId []string
	for _, container := range runningContainers {
		containersId = append(containersId, container.ID[:10])
	}

	return containersId
}