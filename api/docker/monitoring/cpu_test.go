package monitoring

import (
	"testing"
)

type statsTests struct {
	cpuUsageT      int64
	precpuUsageT   int64
	cpuSysUsage    int64
	precpuSysUsage int64
	cpuPercpuUsage []int64
	expected       CpuInfo
}

func TestGetCpuInfo(t *testing.T) {

	// Test case systemDelta > 0 &  case systemDelta < 0
	tests := []statsTests{
		{5, 1, 2000, 1900, []int64{1, 2}, CpuInfo{2, 8}},
		{5, 1, 1800, 2000, []int64{1, 2}, CpuInfo{2, 0}},
	}

	for _, test := range tests {
		var statsTest DockerStats
		// Fill docker stats sdk struct with test values
		statsTest.CPUStats.CPUUsage.TotalUsage = test.cpuUsageT
		statsTest.PrecpuStats.CPUUsage.TotalUsage = test.precpuUsageT
		statsTest.CPUStats.SystemCPUUsage = test.cpuSysUsage
		statsTest.PrecpuStats.SystemCPUUsage = test.precpuSysUsage
		statsTest.CPUStats.CPUUsage.PercpuUsage = test.cpuPercpuUsage

		testCpuInfo := getCpuInfo(statsTest)

		if test.expected.NbCpu != testCpuInfo.NbCpu || test.expected.CpuPercent != testCpuInfo.CpuPercent {
			t.Fatalf("Fail, cpu computed info are not correct")
		}
	}
}
