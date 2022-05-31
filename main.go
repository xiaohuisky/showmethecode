package main

import (
	"fmt"
	"net/http"
	"showmethecode/go/gin"
	"time"
)

/*** escape 1 run: `go tool compile -S gin.go` 发现，内存是在堆上进行分配的
	"CALL    runtime.newobject(SB)"
**/
/*  escape 2 `go tool compile -m gin.go`  发现
gin.go:19:13: inlining call to fmt.Println
gin.go:16:10: new(Cursor) does not escape
gin.go:19:13: ... argument does not escape
gin.go:19:15: c.X escapes to heap
gin.go:19:20: c.Y escapes to heap
**/
/* escape 2 struct
type Cursor struct {
	X,Y int
}
*/

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

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
	// 测试 new 和 空值的地址是否相同，最终结果显示是相同的
	//a := struct{}{}
	//b := new(struct{})
	//fmt.Println(a == *b) // true

	//r := gin.New()
	//r.Use(gin.Logger())
	//r.SetFuncMap(template.FuncMap{
	//	"FormatAsDate": FormatAsDate,
	//})
	//r.LoadHtmlGlob("templates/*")
	//r.Static("/assets", "./static")
	//
	//stu1 := &student{Name: "Geek", Age: 20}
	//stu2 := &student{Name: "D", Age: 22}
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "css.tmpl", nil)
	//})
	//r.GET("/students", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "arr.tmpl", gin.H{
	//		"title":  "gee",
	//		"stuArr": [2]*student{stu1, stu2},
	//	})
	//})
	//
	//r.GET("/date", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "custom_func.tmpl", gin.H{
	//		"title": "gee",
	//		"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
	//	})
	//})
	//
	//r.Run(":9999")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World\n")
	})
	r.GET("/panic", func(c *gin.Context) {
		names := []string{"geek"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}

/*
func onlyForV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Fail(500, "internal server error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
*/
