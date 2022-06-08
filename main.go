package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"showmethecode/go/groupcache"
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

//type student struct {
//	Name string
//	Age  int8
//}
//
//func FormatAsDate(t time.Time) string {
//	year, month, day := t.Date()
//	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
//}

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

	/*
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello World\n")
		})
		r.GET("/panic", func(c *gin.Context) {
			names := []string{"geek"}
			c.String(http.StatusOK, names[100])
		})
		r.Run(":9999")
	*/
	/*
		var i int8 = -9
		fmt.Println("original:", i)
		fmt.Printf("%s\n", biu.ToBinaryString(i))
		r := i >> 2
		fmt.Printf("%s\n", biu.ToBinaryString(r))
		fmt.Println("result:", r)
	*/

	/*
		var i float32 = 99.99
		exponentLen := 8           // 指数部分的长度
		var middleware int64 = 127 // 中间数
		// 获取完整的二进制存储值
		str := biu.ToBinaryString(math.Float32bits(i))
		fmt.Println("str: ", str)
		// 只保留 0101 值
		newStr := strings.ReplaceAll(str[1:len(str)-1], " ", "")
		fmt.Println("newStr: ", newStr)
		// 数值切分逻辑

		sign := newStr[0:1]
		exponent := newStr[1 : 1+exponentLen]
		fraction := newStr[1+exponentLen:]
		fmt.Println("sign: ", sign)
		fmt.Println("exponent: ", exponent)
		fmt.Println("fraction: ", fraction)
		// 指数部分值计算逻辑
		decimalExponent, _ := strconv.ParseInt(exponent, 2, 32)
		fmt.Println("decimalExponent: ", decimalExponent)
		exponentValue := 1 << (decimalExponent - middleware)
		fmt.Println("exponentValue: ", exponentValue)
		// 有效数字部分计算逻辑
		decimalFraction, _ := strconv.ParseInt(fraction, 2, 32)
		fmt.Println("decimalFraction: ", decimalFraction)
		dividend := 1 << (len(newStr) - 1 - exponentLen)
		fractionValue := float64(decimalFraction)/float64(dividend) + 1
		fmt.Println("fractionValue: ", fractionValue)

		fmt.Println(fractionValue * float64(exponentValue))
	*/

	// group cache
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "Geecache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	gee := createGroup()
	if api {
		go startAPIServer(apiAddr, gee)
	}
	startCacheServer(addrMap[port], []string(addrs), gee)
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func createGroup() *groupcache.Group {
	return groupcache.NewGroup("scores", 2<<10, groupcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

// 启动缓存服务器：创建 HTTPPool，添加节点信息，注册到 gee 中，启动 HTTP 服务（共3个端口，8001/8002/8003），用户不感知。
func startCacheServer(addr string, addrs []string, gee *groupcache.Group) {
	peers := groupcache.NewHTTPPool(addr)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

// 启动一个 API 服务（端口 9999），与用户进行交互，用户感知。
func startAPIServer(apiAddr string, gee *groupcache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())

		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))

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
