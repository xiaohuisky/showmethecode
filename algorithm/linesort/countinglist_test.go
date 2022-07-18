package linesort

import "testing"

func TestCountingSort(t *testing.T) {
	arr := []int{
		0, 3, 5, 2, 2, 4, 2, 1, 7, 8, 10,
	}
	t.Log(countingSort(arr))
}
