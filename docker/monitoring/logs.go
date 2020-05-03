package monitoring

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io/ioutil"
	"log"
)

func RetrieveLogs(containerId string) []byte {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Println(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(context.Background(), containerId, options)
	if err != nil {
		log.Printf("Fail to get logs for container %s", containerId)
	} else {

	}
	logs, err := ioutil.ReadAll(out)
	return logs
}
