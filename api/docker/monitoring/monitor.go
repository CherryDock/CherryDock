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

func getStats(containerID string) MonitoringStats {
	cli, err := client.NewEnvClient()

	if err != nil {
		log.Println(err)
	}

	// Retrieve all statistics
	containerStats, err := cli.ContainerStats(context.Background(), containerID, false)
	containerStatsBody, err := ioutil.ReadAll(containerStats.Body)

	// Parse json to struct
	stats := DockerStats{}
	json.Unmarshal(containerStatsBody, &stats)

	cpuInfo := getCpuInfo(stats)
	memoryInfo := getMemoryInfo(stats)
	networkInfo := networkStats(stats)

	stt := MonitoringStats{cpuInfo, memoryInfo, networkInfo}

	return stt
}

func SingleMonitoring(containerID string) []byte {

	stats := getStats(containerID)

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

	Info := Info{
		{"CpuUsage", cpuPercent, string("%")},
		{"MemoryUsage", memoryPercent, string("%")},
		{"Memory", memoryValue, memoryUnit},
		{"NetworkInfoIn", networkIn, networkInUnit},
		{"NetworkInfoOut", networkOut, networkOutUnit},
	}

	return json_utils.FormatToJson(ContainerStat{containerID, Info})
}

func GlobalMonitoring() *GlobalStats {
	// Extract running containers id
	runningContainers := ContainersId(false)

	var containersStat []ContainerStat
	var globalInfo Info
	var nbRunningContainers = len(runningContainers)
	var globalMemoryUsage = 0.0
	var globalCPUUsage = 0.0
	var memoryLimit float64
	var memoryUnit string
	var nbCPU int

	var wg sync.WaitGroup
	wg.Add(nbRunningContainers)

	// Retrieve container stat concurrently
	for i := 0; i < nbRunningContainers; i++ {
		go func(i int) {
			defer wg.Done()
			containerID := runningContainers[i]
			stats := getStats(containerID)

			// Extract memory info
			memoryLimit, memoryUnit = byteConversion(stats.MemoryInfo.Limit)
			memoryPercent := stats.MemoryInfo.UtilizationPercent
			memoryValue, memoryUnit := byteConversion(stats.MemoryInfo.MemoryUsage)
			globalMemoryUsage += memoryPercent

			// Extract network info
			networkIn := stats.NetworkInfo.In.Value
			networkInUnit := stats.NetworkInfo.In.Unit
			networkOut := stats.NetworkInfo.Out.Value
			networkOutUnit := stats.NetworkInfo.Out.Unit

			// Extract cpu info
			nbCPU = stats.CpuInfo.NbCpu
			cpuPercent := stats.CpuInfo.CpuPercent
			globalCPUUsage += cpuPercent

			containersStat = append(containersStat,
				ContainerStat{containerID,
					Info{
						{"CpuUsage", cpuPercent, string("%")},
						{"MemoryUsage", memoryPercent, string("%")},
						{"Memory", memoryValue, memoryUnit},
						{"NetworkInfoIn", networkIn, networkInUnit},
						{"NetworkInfoOut", networkOut, networkOutUnit},
					}})
		}(i)
	}
	wg.Wait()

	globalInfo = Info{
		{"RunningContainers", float64(nbRunningContainers), ""},
		{"NbCpu", float64(nbCPU), ""},
		{"Memory", memoryLimit, memoryUnit},
		{"MemoryUsage", globalMemoryUsage, string("%")},
		{"CpuUsage", globalCPUUsage, string("%")},
	}

	globalStats := GlobalStats{
		globalInfo,
		containersStat,
	}
	return &globalStats
}

type MonitoringStats struct {
	CpuInfo     CpuInfo
	MemoryInfo  MemoryInfo
	NetworkInfo NetworkInfo
}

type Info []struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type ContainerStat struct {
	ID   string `json:"Id"`
	Info Info   `json:"Info"`
}

type GlobalStats struct {
	Info       Info            `json:"GlobalInfo"`
	Containers []ContainerStat `json:"Containers"`
}
