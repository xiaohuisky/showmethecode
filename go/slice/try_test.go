package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTryAppendInt(t *testing.T) {
	a := make([]int, 11, 20)
	assert.Equal(t, a, AppendInt(0))
}

func TestCapacityExpansion(t *testing.T) {
	expected := 512
	assert.Equal(t, expected, CapacityExpansion(511))
}

func TestJsonUnmarshal(t *testing.T) {
	assert.NotEqual(t, nil, JsonUnmarshal())
}

func TestMakeSliceInit( T *testing.T) {
	assert.Equal(T, 513, MakeSliceInit(513))
}