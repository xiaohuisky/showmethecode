package binarysearch

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{
		4, 5, 2, 3, 1,
	}
	t.Log(quickSort2(arr))
}
