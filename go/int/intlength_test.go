package int

import (
	"fmt"
	"testing"
)

func TestIntLength(t *testing.T) {
	t.Log(intType())
	fmt.Println((^uint(0) >> 32&1))
}
