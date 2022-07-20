package radixsort

import (
	"math"
)

func RadixSort(arr []int) []int {
	maxBits := howManyBits(maximum(arr))
	result := make([]int, len(arr))
	for i := 0; i < maxBits; i++ {
		count := make([]int, 10)
		division := math.Pow(10, float64(i))
		for _, num := range arr {
			n1 := num / int(division) % 10
			count[n1]++
		}
		for j := 1; j < len(count); j++ {
			count[j] += count[j-1]
		}
		for n := len(arr) - 1; n >= 0; n-- {
			num := arr[n] / int(division) % 10
			result[count[num]-1] = arr[n]
			count[num]--
		}
		copy(arr, result)
	}
	return result
}

func maximum(list []int) int {
	max := 0
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}

func howManyBits(number int) int {
	count := 0
	for number != 0 {
		number = number / 10
		count++
	}
	return count
}
