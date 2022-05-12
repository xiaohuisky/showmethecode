package showmethecode

/*** escape 1 run: `go tool compile -S main.go` 发现，内存是在堆上进行分配的
	"CALL    runtime.newobject(SB)"
**/
/*  escape 2 `go tool compile -m main.go`  发现
main.go:19:13: inlining call to fmt.Println
main.go:16:10: new(Cursor) does not escape
main.go:19:13: ... argument does not escape
main.go:19:15: c.X escapes to heap
main.go:19:20: c.Y escapes to heap
**/
/* escape 2 struct
type Cursor struct {
	X,Y int
}
*/

func main() {
	/* escape 1
	var a [1]int
	c := a[:]
	fmt.Println(c)
	*/
	/* escape 2
	c := new(Cursor)
	c.Y = 1
	c.X = 2
	fmt.Println(c.X, c.Y)
	 */
}


