package slice

import "fmt"

func WhatIsA() []int {
	var a []int
	f(a)
	return a // a 等于什么？为什么？
}

func f(a []int) {
	fmt.Println(a) // []
	a = append(a, 1)
	fmt.Println(a) // [1]
}
