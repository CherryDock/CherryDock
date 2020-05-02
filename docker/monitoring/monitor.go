package monitoring

import (
	json_utils "github.com/Fszta/DockerMonitoring/jsonutils"
	"log"
	"sync"
	"time"
)

type GlobalMemoryStats struct {
	RunningContainers int
	GlobalUsePercent float64
	ContainersMemory []ContainerMemory
}

func memoryMonitoring() ([]ContainerMemory, int,float64){
	log.Println("Monitor running containers memory...")

	// Extract running containers id
	runningContainers := ContainersId(false)

	// Extract memory statistics
	var containersMemory []ContainerMemory
	var nbContainers = len(runningContainers)
	var globalMemoryUsage = 0.0
	var wg sync.WaitGroup
	wg.Add(nbContainers)

	for i := 0; i < nbContainers; i++ {
		go func(i int) {
			defer wg.Done()
			containerId := runningContainers[i]
			memory := getContainerMemory(containerId)
			globalMemoryUsage += memory.Memory.UtilizationPercent
			containersMemory = append(containersMemory, memory)
		}(i)
	}
	wg.Wait()

	return containersMemory, nbContainers, globalMemoryUsage
}

func GetMemoryStats() []byte {
	memoryStats,runningContainers,globalMemoryPercent := memoryMonitoring()

	globalStats := GlobalMemoryStats{
		runningContainers,
		globalMemoryPercent,
		memoryStats,
	}
	return json_utils.FormatToJson(globalStats)
}

func scheduleMonitor(timeInterval time.Duration) {
	ticker := time.NewTicker(timeInterval * time.Second)
	go func() {
		for t := range ticker.C {
			_ = t
			log.Println("Hello !!")
		}
	}()
}
