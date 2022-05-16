package channel

import (
	"github.com/stretchr/testify/assert"
	_ "net/http/pprof"

	"testing"
)

func TestPrint(t *testing.T) {
	assert.Equal(t, 0, sequenceFmt())
}

func TestTestGoroutineFor(t *testing.T) {
	t.Log(testGoroutineFor())
}

func TestControlGoroutine(t *testing.T) {
	ControlGoroutine(5)
}

func TestMapGoroutineWrite(t *testing.T) {
	strs := []string{
		"a",
		"a",
		"a",
		"al",
		"a",
	}
	r := MapGoroutineWrite(strs, 4)
	assert.Equal(t, 5, r[rune('a')])
}

func TestForSelect(t *testing.T) {
	forSelect()
}
