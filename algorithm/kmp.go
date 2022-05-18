package algorithm

func KMP(a, b string) int {
	n, m := len(a), len(b)
	next := next(b)
	i, j := 0, 0
	for i < n && j < m {
		if j == -1 || a[i] == b[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == m {
		return i - j
	}
	return -1
}

func next(b string) []int {
	m := len(b)
	next := make([]int, m)
	next[0] = -1
	next[1] = 0
	i, j := 0, 1
	for j < m-1 {
		if i == -1 || b[i] == b[j] {
			i++
			j++
			next[j] = i
		} else {
			i = next[i]
		}
	}
	return next
}
