package _map

import (
	"fmt"
	"sync"
	"time"
)

func runWithPanic() {
	s := make(map[int]int)
	n := 100
	for i := 0; i < n; i++ {
		go func(i int) {
			s[i] = i
		}(i)
	}
	for i := 0; i <= n; i++ {
		go func(i int) {
			fmt.Printf("第 %d 个元素是 %v", i, s[i])
		}(i)
	}
	time.Sleep(time.Second)
	// fatal error: concurrent map writes
}

// runWithSyncRWMutex mqp + sync.RWMutex
func runWithSyncRWMutex() {
	var lock sync.RWMutex
	s := make(map[int]int)
	n := 100
	for i := 0; i < n; i++ {
		go func(i int) {
			lock.Lock()
			s[i] = i
			lock.Unlock()
		}(i)
	}
	for i := 0; i <= n; i++ {
		go func(i int) {
			lock.RLock()
			fmt.Printf("第 %d 个元素是%v；", i, s[i])
			lock.RUnlock()
		}(i)
	}
	time.Sleep(time.Second)
}

// RunWithSyncMap
func RunWithSyncMap() {
	s := sync.Map{}
	n := 100
	for i := 0; i < n; i++ {
		go func(i int) {
			s.Store(i, i)
		}(i)
	}
	for i := 0; i <= n; i++ {
		go func(i int) {
			v, ok := s.Load(i)
			if ok {
				fmt.Printf("第 %d 个元素是%v；", i, v)
			}
		}(i)
	}
	time.Sleep(time.Second)
}
