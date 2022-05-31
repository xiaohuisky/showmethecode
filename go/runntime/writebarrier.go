package runntime

var A Wb
var B Wb

type Wb struct {
	Obj *int
}

// go build -gcflags "-N -I"
// go tool objdump -s 'main\.simpleSet' -S ./main.exe
/*
func simpleSet(c *int) {
	A.Obj = nil
	B.Obj = c
	// if GC Begin
	// scan A
	A.Obj = c
	B.Obj = nil
	//scan B
	a := []byte{
		1, 2,
	}
	b := "1"
	c := reflect.StringHeader{&a, &b}
}
*/
