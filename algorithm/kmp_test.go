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

func TestKMP1(t *testing.T) {
	assert.Equal(t, 2, KMP("北京天安门最美丽", "天安门"))
}
