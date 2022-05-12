package int

import (
	"fmt"
)
const intSize = 32 << (^uint(0) >> 63)

func intType() int {
	fmt.Println(intSize)
	fmt.Println((^uint(0) >> 63))

	return 0
}
