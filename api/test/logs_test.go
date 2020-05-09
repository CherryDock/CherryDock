package test

import (
	"github.com/CherryDock/CherryDock/api/docker/monitoring"
	"testing"
)

type TestLogs struct {
	ContainerId         string
	ExpectedLogsContent []byte
	ExpectedStatus      bool
}

func TestRetrieveLogs(t *testing.T) {
	var expectedLogs []byte
	expectedStatus := false

	testLogsBadId := TestLogs{"xxxxxx", expectedLogs, expectedStatus}

	logs, succeed := monitoring.RetrieveLogs(testLogsBadId.ContainerId)

	if len(logs) != 0 && succeed != expectedStatus {
		t.Fatalf("Logs should be empty & succeed equal to false")
	}
}
