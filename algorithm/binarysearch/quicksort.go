package binarysearch

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		key := arr[(left+right)/2]
		i := left
		j := right
		for {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		_quickSort(arr, left, i-1)
		_quickSort(arr, j+1, right)
	}
	return arr
}

func quickSort2(arr []int) []int {
	return _quickSort2(arr, 0, len(arr)-1)
}

func _quickSort2(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort2(arr, left, partitionIndex-1)
		_quickSort2(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, p, r int) int {
	x := arr[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
