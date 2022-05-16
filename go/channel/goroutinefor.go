package channel

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	_ "net/http/pprof"
)

// result panic: too many concurrent operations on a single file or socket (max 1048575)
func testGoroutineFor() int {
	i := 0
	for {
		go func() {
			i++
			fmt.Println(i)
		}()
	}
	return 0
}

// 浏览器中打开链接127.0.0.1:6060/debug/pprof/goroutine 会下载goroutine文件。
// 下载后，在命令行下执行： go tool pprof -http=":8081" goroutine 会自动打开浏览器，清晰的看到goroutine的数量以及调用关系。
func ControlGoroutine(concurrent int) {
	w := sync.WaitGroup{}
	ch := make(chan bool, concurrent)
	i := 0
	for {
		i++
		ch <- true
		w.Add(1)
		go func(i int) {
			defer w.Done()
			log.Println(http.ListenAndServe("localhost:6060", nil))
			time.Sleep(time.Millisecond)
			log.Println("Goroutine", runtime.NumGoroutine()) // max is 9
			log.Println(i)
			<-ch
		}(i)
	}
	w.Wait()
}

func MapGoroutineWrite(strs []string, concurrent int) map[rune]int {
	w := sync.WaitGroup{}
	ch := make(chan bool, concurrent)
	sm := sync.Map{}
	for _, str := range strs {
		w.Add(1)
		ch <- true
		go func(str string) {
			defer w.Done()
			for _, v := range str {
				n, ok := sm.Load(v)
				if ok {
					nInt, ok := n.(int)
					if ok {
						nInt++
						sm.Store(v, nInt)
					}
				} else {
					sm.LoadOrStore(v, 1)
				}
			}
			<-ch
		}(str)
	}
	w.Wait()
	fmt.Println("sm", sm)
	r := make(map[rune]int)
	sm.Range(func(key, value any) bool {
		fmt.Println(key, value)
		k, ok1 := key.(rune)
		v, ok2 := value.(int)
		if ok1 && ok2 {
			r[k] = v
		}
		return true
	})
	return r
}
