package channel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	assert.Equal(t, 0, sequenceFmt())
}

func TestTestGoroutineFor(t *testing.T) {
	t.Log(testGoroutineFor())
}