package monitoring

import "testing"

func TestComputeMemory(t *testing.T) {

	memoryStats := ComputeMemory(4.0, 2.0, 1.0)
	expected := MemoryInfo{
		1,
		4,
		25.0,
	}
	if memoryStats != expected {
		t.Errorf("Memory stats should be %v", expected)
	}

	memoryStats = ComputeMemory(1.0, 0, 0)
	expected = MemoryInfo{
		0,
		1,
		0,
	}

	if memoryStats != expected {
		t.Errorf("Memory stats should be %v", expected)
	}
}
