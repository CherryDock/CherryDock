package test

import "testing"
import "github.com/CherryDock/CherryDock/docker/monitoring"

func TestComputeMemory(t *testing.T) {

	memoryStats := monitoring.ComputeMemory(4.0, 2.0, 1.0)
	expected := monitoring.MemoryInfo{
		1,
		4,
		25.0,
	}
	if memoryStats != expected {
		t.Errorf("Memory stats should be %v", expected)
	}

	memoryStats = monitoring.ComputeMemory(1.0, 0, 0)
	expected = monitoring.MemoryInfo{
		0,
		1,
		0,
	}

	if memoryStats != expected {
		t.Errorf("Memory stats should be %v", expected)
	}
}
