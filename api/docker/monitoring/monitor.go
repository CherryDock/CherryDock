package monitoring

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
	"time"

	json_utils "github.com/CherryDock/CherryDock/api/jsonutils"
	"github.com/docker/docker/client"
)

func getStats(containerId string) *ContainerStats {
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

	return &stt
}

func SingleMonitoring(containerId string) []byte {

	stats := getStats(containerId)

	// Extract memory info
	memoryPercent := stats.MemoryInfo.UtilizationPercent
	memoryValue, memoryUnit := byteConversion(stats.MemoryInfo.MemoryUsage)

	// Extract Cpu percentage of use
	cpuPercent := stats.CpuInfo.CpuPercent

	containerInfo := Container{
		containerId,
		Info{
			cpuPercent,
			memoryPercent,
			Memory{
				memoryValue,
				memoryUnit},
			stats.NetworkInfo,
		}}

	return json_utils.FormatToJson(containerInfo)
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
	var stats *ContainerStats

	var nbCpu int

	var wg sync.WaitGroup
	wg.Add(nbRunningContainers)

	// Retrieve container stat concurrently
	for i := 0; i < nbRunningContainers; i++ {
		go func(i int) {
			defer wg.Done()
			containerId := runningContainers[i]
			stats = getStats(containerId)

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

type (
	// GlobalStats define format of monitoring statistics json
	GlobalStats struct {
		RunningContainers  int
		NbCpu              int
		MemoryLimit        Memory
		MemoryUsagePercent float64
		CpuUsagePercent    float64
		Containers         []Container
	}

	// ContainerStats define format of statistics
	// for a single container
	ContainerStats struct {
		CpuInfo     CpuInfo
		MemoryInfo  MemoryInfo
		NetworkInfo NetworkInfo
	}

	Info struct {
		CpuUsagePercent    float64
		MemoryUsagePercent float64
		Memory             Memory
		NetworkInfo        NetworkInfo
	}

	MemoryInfo struct {
		MemoryUsage        float64
		Limit              float64
		UtilizationPercent float64
	}

	NetworkInfo struct {
		In struct {
			Value float64
			Unit  string
		}
		Out struct {
			Value float64
			Unit  string
		}
	}
	Rx struct {
		Value float64
		Unit  string
	}
	Tx struct {
		Value float64
		Unit  string
	}

	CpuInfo struct {
		NbCpu      int
		CpuPercent float64
	}

	Memory struct {
		Value float64
		Unit  string
	}

	Container struct {
		Id   string
		Info Info
	}

	DockerStats struct {
		BlkioStats struct {
			IoMergedRecursive       []interface{} `json:"io_merged_recursive"`
			IoQueueRecursive        []interface{} `json:"io_queue_recursive"`
			IoServiceBytesRecursive []struct {
				Major int    `json:"major"`
				Minor int    `json:"minor"`
				Op    string `json:"op"`
				Value int    `json:"value"`
			} `json:"io_service_bytes_recursive"`
			IoServiceTimeRecursive []interface{} `json:"io_service_time_recursive"`
			IoServicedRecursive    []struct {
				Major int    `json:"major"`
				Minor int    `json:"minor"`
				Op    string `json:"op"`
				Value int    `json:"value"`
			} `json:"io_serviced_recursive"`
			IoTimeRecursive     []interface{} `json:"io_time_recursive"`
			IoWaitTimeRecursive []interface{} `json:"io_wait_time_recursive"`
			SectorsRecursive    []interface{} `json:"sectors_recursive"`
		} `json:"blkio_stats"`
		CPUStats struct {
			CPUUsage struct {
				PercpuUsage       []int64 `json:"percpu_usage"`
				TotalUsage        int64   `json:"total_usage"`
				UsageInKernelmode int     `json:"usage_in_kernelmode"`
				UsageInUsermode   int64   `json:"usage_in_usermode"`
			} `json:"cpu_usage"`
			OnlineCpus     int   `json:"online_cpus"`
			SystemCPUUsage int64 `json:"system_cpu_usage"`
			ThrottlingData struct {
				Periods          int `json:"periods"`
				ThrottledPeriods int `json:"throttled_periods"`
				ThrottledTime    int `json:"throttled_time"`
			} `json:"throttling_data"`
		} `json:"cpu_stats"`
		ID          string `json:"id"`
		MemoryStats struct {
			Limit    int64 `json:"limit"`
			MaxUsage int   `json:"max_usage"`
			Stats    struct {
				ActiveAnon              int   `json:"active_anon"`
				ActiveFile              int   `json:"active_file"`
				Cache                   int   `json:"cache"`
				Dirty                   int   `json:"dirty"`
				HierarchicalMemoryLimit int64 `json:"hierarchical_memory_limit"`
				HierarchicalMemswLimit  int   `json:"hierarchical_memsw_limit"`
				InactiveAnon            int   `json:"inactive_anon"`
				InactiveFile            int   `json:"inactive_file"`
				MappedFile              int   `json:"mapped_file"`
				Pgfault                 int   `json:"pgfault"`
				Pgmajfault              int   `json:"pgmajfault"`
				Pgpgin                  int   `json:"pgpgin"`
				Pgpgout                 int   `json:"pgpgout"`
				Rss                     int   `json:"rss"`
				RssHuge                 int   `json:"rss_huge"`
				TotalActiveAnon         int   `json:"total_active_anon"`
				TotalActiveFile         int   `json:"total_active_file"`
				TotalCache              int   `json:"total_cache"`
				TotalDirty              int   `json:"total_dirty"`
				TotalInactiveAnon       int   `json:"total_inactive_anon"`
				TotalInactiveFile       int   `json:"total_inactive_file"`
				TotalMappedFile         int   `json:"total_mapped_file"`
				TotalPgfault            int   `json:"total_pgfault"`
				TotalPgmajfault         int   `json:"total_pgmajfault"`
				TotalPgpgin             int   `json:"total_pgpgin"`
				TotalPgpgout            int   `json:"total_pgpgout"`
				TotalRss                int   `json:"total_rss"`
				TotalRssHuge            int   `json:"total_rss_huge"`
				TotalUnevictable        int   `json:"total_unevictable"`
				TotalWriteback          int   `json:"total_writeback"`
				Unevictable             int   `json:"unevictable"`
				Writeback               int   `json:"writeback"`
			} `json:"stats"`
			Usage int `json:"usage"`
		} `json:"memory_stats"`
		Name     string `json:"name"`
		Networks struct {
			Eth0 struct {
				RxBytes   int `json:"rx_bytes"`
				RxDropped int `json:"rx_dropped"`
				RxErrors  int `json:"rx_errors"`
				RxPackets int `json:"rx_packets"`
				TxBytes   int `json:"tx_bytes"`
				TxDropped int `json:"tx_dropped"`
				TxErrors  int `json:"tx_errors"`
				TxPackets int `json:"tx_packets"`
			} `json:"eth0"`
		} `json:"networks"`
		NumProcs  int `json:"num_procs"`
		PidsStats struct {
			Current int `json:"current"`
		} `json:"pids_stats"`
		PrecpuStats struct {
			CPUUsage struct {
				PercpuUsage       []int64 `json:"percpu_usage"`
				TotalUsage        int64   `json:"total_usage"`
				UsageInKernelmode int     `json:"usage_in_kernelmode"`
				UsageInUsermode   int64   `json:"usage_in_usermode"`
			} `json:"cpu_usage"`
			OnlineCpus     int   `json:"online_cpus"`
			SystemCPUUsage int64 `json:"system_cpu_usage"`
			ThrottlingData struct {
				Periods          int `json:"periods"`
				ThrottledPeriods int `json:"throttled_periods"`
				ThrottledTime    int `json:"throttled_time"`
			} `json:"throttling_data"`
		} `json:"precpu_stats"`
		Preread      time.Time `json:"preread"`
		Read         time.Time `json:"read"`
		StorageStats struct {
		} `json:"storage_stats"`
	}
)
