package monitoring

type MemoryInfo struct {
	MemoryUsage        float64
	Limit              float64
	UtilizationPercent float64
}

func getMemoryInfo(stats DockerStats) MemoryInfo {
	memoryStats := stats.MemoryStats

	limit := memoryStats.Limit
	usage := memoryStats.Usage
	cache := memoryStats.Stats.Cache

	memoryInfo := ComputeMemory(float64(limit), float64(usage), float64(cache))

	return memoryInfo
}

func ComputeMemory(limit float64, usage float64, cache float64) MemoryInfo {
	memoryUsage := usage - cache
	memoryUtilizationPercent := memoryUsage / limit * 100
	memoryStats := MemoryInfo{
		memoryUsage,
		limit,
		memoryUtilizationPercent,
	}

	return memoryStats
}

func byteConversion(value float64) float64 {
	var converted float64
	if value > 1073741824 {
		converted = value / float64(1073741824)
	} else {
		converted = value / float64(1048576)
	}
	return converted
}
