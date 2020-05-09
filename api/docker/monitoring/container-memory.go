package monitoring

import (
	"context"
	jsonutils "github.com/CherryDock/CherryDock/api/jsonutils"
	"github.com/docker/docker/client"
	"io/ioutil"
	"log"
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
	memoryStats := extractMemoryStats(memoryData)
	log.Println(ContainerMemory{containerId,memoryStats})

	return 	ContainerMemory{containerId,memoryStats}
}

func extractMemoryStats(memoryField interface{}) MemoryStats {

	// Extract subfield to compute basics memory metrics
	limit := memoryField.(map[string]interface{})["limit"].(float64)
	usage := memoryField.(map[string]interface{})["usage"].(float64)
	cache := memoryField.(map[string]interface{})["stats"].(map[string]interface{})["cache"].(float64)

	var memoryStats MemoryStats = ComputeStats(limit,usage,cache)

	return memoryStats
}

func ComputeStats(limit float64, usage float64, cache float64) MemoryStats {
	memoryUsage := usage - cache
	memoryUtilizationPercent := memoryUsage/limit * 100
	memoryStats := MemoryStats{
		memoryUsage,
		limit,
		memoryUtilizationPercent,
	}

	return memoryStats
}

