package test

import "testing"
import "github.com/Fszta/DockerMonitoring/docker/monitoring"

func TestComputeStats(t *testing.T) {

	memoryStats := monitoring.ComputeStats(4.0, 2.0, 1.0)
	expected := monitoring.MemoryStats{
		1,
		4,
		25.0,
	}
	if memoryStats != expected {
		t.Errorf("Memory stats should be %v", expected)
	}

	memoryStats = monitoring.ComputeStats(1.0, 0, 0)
	expected = monitoring.MemoryStats{
		0,
		1,
		0,
	}

	if memoryStats != expected {
		t.Errorf("Memory stats should be %v", expected)
	}
}
