package channel

import "fmt"

// result panic: too many concurrent operations on a single file or socket (max 1048575)
func testGoroutineFor() int {
	i := 0
	for  {
		go func() {
			i++
			fmt.Println(i)
		}()
	}
	return 0
}
