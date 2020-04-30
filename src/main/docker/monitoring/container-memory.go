package monitoring

import (
	"context"
	jsonutils "github.com/Fszta/DockerMonitoring/src/main/jsonutils"
	"github.com/docker/docker/client"
	"io/ioutil"
)

type MemoryStats struct {
	MemoryUsage float64
	Limit float64
	UtilizationPercent float64
}

type ContainerMemory struct {
	ContainerId string
	Memory      MemoryStats
}

func getContainerMemory(containerId string) ContainerMemory {
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}
	containerStats, err := cli.ContainerStats(context.Background(),containerId,false)
	containerStatsBody, err := ioutil.ReadAll(containerStats.Body)

	// Get data from json
	dockerStats:= jsonutils.ParseJson(containerStatsBody)

	// Extract memory field
	memoryData := dockerStats["memory_stats"]
	memoryStats := getMemoryStats(memoryData)

	return 	ContainerMemory{containerId,memoryStats}
}

func getMemoryStats(memoryField interface{}) MemoryStats {
	// Extract subfield to compute basics memory metrics
	limit := memoryField.(map[string]interface{})["limit"].(float64)
	usage := memoryField.(map[string]interface{})["usage"].(float64)
	cache := memoryField.(map[string]interface{})["stats"].(map[string]interface{})["cache"].(float64)

	// Compute memory usage in byte
	memoryUsage := usage - cache
	memoryUtilizationPercent := memoryUsage/limit * 100

	memoryStats := MemoryStats{memoryUsage,limit,memoryUtilizationPercent}

	return memoryStats
}
