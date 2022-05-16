package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTryAppendInt(t *testing.T) {
	assert.Equal(t, 512, AppendInt(254))
	assert.Equal(t, 608, AppendInt(260))
	assert.Equal(t, 608, AppendInt(300))
	assert.Equal(t, 768, AppendInt(400))
	assert.Equal(t, 848, AppendInt(500))
}

func TestCapacityExpansion(t *testing.T) {
	expected := 512
	assert.Equal(t, expected, CapacityExpansion(511))
}

func TestJsonUnmarshal(t *testing.T) {
	assert.NotEqual(t, nil, JsonUnmarshal())
}

func TestMakeSliceInit(T *testing.T) {
	assert.Equal(T, 513, MakeSliceInit(513))
}

func TestSliceAppend(T *testing.T) {
	SliceAppend()
}
