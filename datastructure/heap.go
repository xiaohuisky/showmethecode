package datastructure

var a []int   // 数组，用于存储数据
var n int     // 堆可以存储的最大数据个数
var count int // 堆中已经存储的数据个数

func Heap(capacity int) {
	a = make([]int, capacity)
	n = capacity
	count = 0
}

func insert(data int) {
	if count >= n { // 堆满了
		return
	}
	count++
	a[count] = data
	i := count
	for i/2 > 0 && a[i] > a[i/2] { // 自下向上堆化
		swap(a, i, i/2) // 交换下标为 i 和 i/2 的两个元素
	}
	i = i / 2
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func removeMax() int {
	if count == 0 {
		return -1 // 堆中没有数据
	}
	a[1] = a[count]
	count--
	return heapify(a, count, 1)
}

func heapify(a []int, n, i int) int { // 自上往下堆化
	for true {
		maxPos := i
		if i*2 <= n && a[i] < a[i*2] {
			maxPos = i * 2
		}
		if i*2+1 <= n && a[maxPos] < a[i*2+1] {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		swap(a, i, maxPos)
		i = maxPos
	}
	return -1
}

func buildHeap(a []int, n int) {
	for i := n / 2; i >= 1; i-- {
		heapify(a, n, i)
	}
}

// n 表示数据的个数，数组 a 中的数据从下标 1 到 n 的位置
func sort(a []int, n int) {
	buildHeap(a, n)
	k := n
	for k > 1 {
		swap(a, 1, k)
		k--
		heapify(a, k, 1)
	}
}
