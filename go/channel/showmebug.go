package channel

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 通过协程、通道实现并发按特定顺序打印字符串
// 有三个函数分别打印，“dog”，“cat”，“fish”， 要求每个函数起一个goroutine，请按照dog，cat，fish的顺序，打印四次，输出到控制台。

func sequenceFmt() int {
	var wg sync.WaitGroup
	var dogCounter uint64
	var fishCounter uint64
	var catCounter uint64
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	catCh := make(chan struct{}, 1)

	wg.Add(3)
	dogCh <- struct{}{}
	go dog(&wg, dogCounter, dogCh, fishCh)
	go fish(&wg, fishCounter, fishCh, catCh)
	go cat(&wg, catCounter, catCh, dogCh)
	wg.Wait()
	return 0
}

func dog(wg *sync.WaitGroup, counter uint64, dogCh, fishCh chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<-dogCh
		fmt.Println("dog")
		atomic.AddUint64(&counter, 1)
		fishCh <- struct{}{}
	}
}

func fish(wg *sync.WaitGroup, counter uint64, fishCh, catCh chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<-fishCh
		fmt.Println("fish")
		atomic.AddUint64(&counter, 1)
		catCh <- struct{}{}
	}
}

func cat(wg *sync.WaitGroup, counter uint64, catCh, dogCh chan struct{}) {
	for {
		if counter >= uint64(100) {
			wg.Done()
			return
		}
		<-catCh
		fmt.Println("cat")
		atomic.AddUint64(&counter, 1)
		dogCh <- struct{}{}
	}
}
