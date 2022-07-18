package linesort

func countingSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// 寻找最大的元素
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// 初始化一个长度为 max+1 的数组
	countList := make([]int, max+1, max+1)

	// 计数
	for i := 0; i < len(arr); i++ {
		countList[arr[i]]++
	}

	// 统计计数累计值
	for i := 1; i < max+1; i++ {
		countList[i] += countList[i-1]
	}

	//初始化返回数组
	outPutList := make([]int, len(arr))
	// 将元素放到正确的位置上
	for i := 0; i < len(outPutList); i++ {
		outPutList[countList[arr[i]]-1] = arr[i]
		countList[arr[i]]--
	}
	return outPutList
}
