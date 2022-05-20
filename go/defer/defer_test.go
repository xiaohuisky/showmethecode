package _defer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefer1(t *testing.T) {
	expected := 5
	assert.Equal(t, expected, f1())
}

func TestDefer2(t *testing.T) {
	expected := 1
	assert.Equal(t, expected, f2())
}

func TestDefer3(t *testing.T) {
	expected := 6
	assert.Equal(t, expected, f3())
}
