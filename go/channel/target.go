package channel

import (
	"fmt"
	"sync"
)

// active object对象
type Service struct {
	channel chan int `desc:"即将加入到数据slice的数据"`
	data    []int    `desc:"数据slice"`
}

// 新建一个size大小缓存的active object对象
func NewService(size int, done func()) *Service {
	s := &Service{
		channel: make(chan int, size),
		data:    make([]int, 0),
	}

	go func() {
		s.schedule()
		done()
	}()
	return s
}

// 把管道中的数据append到slice中
func (s *Service) schedule() {
	for v := range s.channel {
		s.data = append(s.data, v)
	}
}

// 增加一个值
func (s *Service) Add(v int) {
	s.channel <- v
}

// 管道使用完关闭
func (s *Service) Close() {
	close(s.channel)
}

// 返回slice
func (s *Service) Slice() []int {
	return s.data
}

func Target() {

	// 1. 新建一个active object, 并增加结束信号
	c := make(chan struct{})
	s := NewService(100, func() { c <- struct{}{} })

	// 2. 起n个goroutine不断执行增加操作
	n := 10000
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(a int) {
			s.Add(a)
			wg.Done()
		}(i)
	}
	wg.Wait()
	s.Close()

	<-c

	// 3. 校验所有结果是否都被添加上
	fmt.Println("done len:", len(s.Slice()))
}
