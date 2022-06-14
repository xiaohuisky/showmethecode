package _map

import (
	"testing"
)

func TestRunWithPanic(t *testing.T) {
	runWithPanic()
}

func TestRunWithSyncRWMutex(t *testing.T) {
	runWithSyncRWMutex()
}

func TestRunWithSyncMap(t *testing.T) {
	RunWithSyncMap()
}

func TestRunWithSyncRWMutex2(t *testing.T) {
	runWithSyncRWMutex2()
}
