package radixsort

import "testing"

func TestRadixSort(t *testing.T) {
	arr := []int{
		1, 99, 1, 2, 7,
	}
	t.Log(RadixSort(arr))
}
