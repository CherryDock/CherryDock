package monitoring

type CpuInfo struct {
	NbCpu      int
	CpuPercent float64
}

func getCpuInfo(stats DockerStats) CpuInfo {

	cpuStats := stats.CPUStats
	precpuStats := stats.PrecpuStats

	percpuUsage := cpuStats.CPUUsage.PercpuUsage
	nbCpu := len(percpuUsage)

	cpuPercent := 0.0
	cpuDelta := cpuStats.CPUUsage.TotalUsage - precpuStats.CPUUsage.TotalUsage

	systemDelta := cpuStats.SystemCPUUsage - precpuStats.SystemCPUUsage

	if systemDelta > 0.0 {
		cpuPercent = float64(cpuDelta) / float64(systemDelta) * (100.0 * float64(nbCpu))
	}

	return CpuInfo{nbCpu, cpuPercent}
}
