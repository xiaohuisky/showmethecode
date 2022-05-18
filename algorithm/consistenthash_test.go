package algorithm

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestConsistentHash(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key)) // 为了方便测试，自己实现的简单的 hash 算法
		return uint32(i)
	})
	hash.Add("2", "6", "4")
	keys1 := []int{2, 4, 6, 12, 14, 16, 22, 24, 26}
	map1 := map[int]string{
		2:  "2",
		4:  "4",
		6:  "6",
		12: "2",
		14: "4",
		16: "6",
		22: "2",
		24: "4",
		26: "6",
	}
	assert.Equal(t, hash.keys, keys1)
	assert.Equal(t, hash.hashMap, map1)

	expected := "2"
	key2 := hash.Get("1")
	key12 := hash.Get("12")
	key22 := hash.Get("22")
	assert.Equal(t, expected, key2)
	assert.Equal(t, expected, key12)
	assert.Equal(t, expected, key22)
}
