package monitoring

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	json_utils "github.com/CherryDock/CherryDock/api/jsonutils"
	"github.com/docker/docker/client"
)

func getStats(containerId string) ContainerStats {
	cli, err := client.NewEnvClient()

	if err != nil {
		log.Println(err)
	}

	// Retrieve all statistics
	containerStats, err := cli.ContainerStats(context.Background(), containerId, false)
	containerStatsBody, err := ioutil.ReadAll(containerStats.Body)

	// Parse json to struct
	stats := DockerStats{}
	json.Unmarshal(containerStatsBody, &stats)

	cpuInfo := getCpuInfo(stats)
	memoryInfo := getMemoryInfo(stats)
	networkInfo := networkStats(stats)

	stt := ContainerStats{cpuInfo, memoryInfo, networkInfo}

	return stt
}

func SingleMonitoring(containerId string) []byte {

	stats := getStats(containerId)

	// Extract memory info
	memoryPercent := stats.MemoryInfo.UtilizationPercent
	memoryValue, memoryUnit := byteConversion(stats.MemoryInfo.MemoryUsage)

	// Extract network info
	networkIn := stats.NetworkInfo.In.Value
	networkInUnit := stats.NetworkInfo.In.Unit
	networkOut := stats.NetworkInfo.Out.Value
	networkOutUnit := stats.NetworkInfo.Out.Unit

	// Extract Cpu percentage of use
	cpuPercent := stats.CpuInfo.CpuPercent

	Info := ContainerInfo{
		{"CpuUsage", cpuPercent, string("%")},
		{"MemoryUsage", memoryPercent, string("%")},
		{"Memory", memoryValue, memoryUnit},
		{"NetworkInfoIn", networkIn, networkInUnit},
		{"NetworkInfoOut", networkOut, networkOutUnit},
	}

	return json_utils.FormatToJson(ContainerStat{containerId, Info})
}

func GlobalMonitoring() *GlobalStats {
	// Extract running containers id
	runningContainers := ContainersId(false)

	var containerInfo []Container
	var nbRunningContainers = len(runningContainers)
	var globalMemoryUsage = 0.0
	var globalCpuUsage = 0.0
	var memoryLimit float64
	var memoryUnit string

	var nbCpu int

	var wg sync.WaitGroup
	wg.Add(nbRunningContainers)

	// Retrieve container stat concurrently
	for i := 0; i < nbRunningContainers; i++ {
		go func(i int) {
			defer wg.Done()
			containerId := runningContainers[i]
			stats := getStats(containerId)

			// Extract memory info
			memoryLimit, memoryUnit = byteConversion(stats.MemoryInfo.Limit)
			memoryPercent := stats.MemoryInfo.UtilizationPercent
			memoryValue, memoryUnit := byteConversion(stats.MemoryInfo.MemoryUsage)
			globalMemoryUsage += memoryPercent

			// Extract cpu info
			nbCpu = stats.CpuInfo.NbCpu
			cpuPercent := stats.CpuInfo.CpuPercent
			globalCpuUsage += cpuPercent

			containerInfo = append(containerInfo,
				Container{
					containerId,
					Info{
						cpuPercent,
						memoryPercent,
						Memory{
							memoryValue,
							memoryUnit},
						stats.NetworkInfo,
					}})
		}(i)
	}
	wg.Wait()

	globalStats := GlobalStats{
		nbRunningContainers,
		nbCpu,
		Memory{memoryLimit, memoryUnit},
		globalMemoryUsage,
		globalCpuUsage,
		containerInfo,
	}

	return &globalStats
}

type ContainerStats struct {
	CpuInfo     CpuInfo
	MemoryInfo  MemoryInfo
	NetworkInfo NetworkInfo
}

type Info struct {
	CpuUsagePercent    float64
	MemoryUsagePercent float64
	Memory             Memory
	NetworkInfo        NetworkInfo
}

type Memory struct {
	Value float64
	Unit  string
}

type ContainerInfo []struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type ContainerStat struct {
	Id   string        `json:"Id"`
	Info ContainerInfo `json:"Info"`
}

type Container struct {
	Id   string
	Info Info
}

type GlobalStats struct {
	RunningContainers  int
	NbCpu              int
	MemoryLimit        Memory
	MemoryUsagePercent float64
	CpuUsagePercent    float64
	Containers         []Container
}
