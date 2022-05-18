package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNext(t *testing.T) {
	expected := []int{
		-1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0,
	}
	assert.Equal(t, expected, next("ewevefetene"))
}

func TestKMP(t *testing.T) {
	assert.Equal(t, 1, KMP("ttest", "tes"))
}
