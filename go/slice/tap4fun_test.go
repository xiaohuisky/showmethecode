package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhatIsA(t *testing.T) {
	f := WhatIsA()
	assert.Equal(t, []int(nil), f)
}
