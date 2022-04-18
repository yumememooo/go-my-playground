package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

//這個範例用來試著創建大量的go，然後透過gin的pprof接口監看效能
func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	v1 := engine.Group("/v1")
	{
		v1.POST("/send", postData) //http://127.0.0.1:8888/v1/send
	}

	pprof.Register(engine) // pprof monitor
	endpoint := fmt.Sprintf(":%d", 8888)
	go func() { //分析用
		err := engine.Run(endpoint)
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Printf("end call")
	select {} //避免退出
}
func postData(c *gin.Context) {
	fmt.Println("Hello World!,I am done.")
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

// 1. 先觀察啟動的量
// http://localhost:8888/debug/pprof/
// 5	goroutine
// 8	heap

// 2. 用baton 打ＰＯＳＴ http://127.0.0.1:8888/v1/send
//go get -u github.com/americanexpress/baton  go1.17
//baton -b test -m POST -u http://127.0.0.1:8888/v1/send -c 10 -t 10
// =========================== Results ========================================
// Total requests:                                150850
// Time taken to complete requests:        10.000603959s
// Requests per second:                            15082
// Number of 2xx responses:                       150850
// ===========================================================================
// 過程中增長到 17 goroutine

// 結束
//== 5	goroutine
//== 66	heap
// Showing nodes accounting for 4610.84kB, 100% of 4610.84kB total
// Showing top 10 nodes out of 37
//       flat  flat%   sum%        cum   cum%
//  2050.25kB 44.47% 44.47%  2050.25kB 44.47%  runtime.allocm
//   512.31kB 11.11% 55.58%   512.31kB 11.11%  regexp/syntax.(*compiler).inst
//   512.20kB 11.11% 66.69%   512.20kB 11.11%  runtime.malg
//   512.04kB 11.11% 77.79%   512.04kB 11.11%  runtime.bgscavenge
//   512.03kB 11.10% 88.90%   512.03kB 11.10%  net/textproto.(*Reader).ReadMIMEHeader
//   512.01kB 11.10%   100%   512.01kB 11.10%  github.com/gin-gonic/gin.glob..func1

/////===========
// 再加大測一次  -c 200 Number of concurrent requests
//test -m POST -u http://127.0.0.1:8888/v1/send -c 200 -t 10
// // =========================== Results ========================================
// // Total requests:                                164992
// // Time taken to complete requests:        10.019925294s
// // Requests per second:                            16466
// // ===========================================================================
// 206	goroutine
// 78	heap
// 結束： 看起來沒問題！！！！！！runtime.malg 沒增長
//6	goroutine
//81	heap
// Showing top 10 nodes out of 34
//       flat  flat%   sum%        cum   cum%
//  4107.01kB 44.47% 44.47%  4107.01kB 44.47%  bufio.NewWriterSize (inline)
//  2050.25kB 22.20% 66.67%  2050.25kB 22.20%  runtime.allocm
//  1542.01kB 16.70% 83.36%  1542.01kB 16.70%  bufio.NewReaderSize (inline)
//   512.31kB  5.55% 88.91%   512.31kB  5.55%  regexp/syntax.(*compiler).inst (inline)
//   512.20kB  5.55% 94.46%   512.20kB  5.55%  runtime.malg
//   512.04kB  5.54%   100%   512.04kB  5.54%  runtime.bgscavenge
//          0     0%   100%  1542.01kB 16.70%  bufio.NewReader (inline)
//          0     0%   100%   512.31kB  5.55%  github.com/go-playground/validator/v10.init
//          0     0%   100%     2565kB 27.77%  net/http.(*conn).readRequest
//          0     0%   100%  5649.02kB 61.16%  net/http.(*conn).serve

//=========================================
// 再加大 -c 20000
// test -m POST -u http://127.0.0.1:8888/v1/send -c 20000 -t 10
// =========================== Results ========================================
// Total requests:                              34123782
// Time taken to complete requests:        10.667939895s
// Requests per second:                          3198723
// 517 	goroutine
// (pprof) top
// Showing nodes accounting for 10257.11kB, 100% of 10257.11kB total
// Showing top 10 nodes out of 49
//       flat  flat%   sum%        cum   cum%
//  3594.04kB 35.04% 35.04%  3594.04kB 35.04%  bufio.NewWriterSize
//  2050.25kB 19.99% 55.03%  2050.25kB 19.99%  runtime.allocm
//     1028kB 10.02% 65.05%     1028kB 10.02%  bufio.NewReaderSize (inline)
//  1024.05kB  9.98% 75.03%  1024.05kB  9.98%  net/textproto.(*Reader).ReadMIMEHeader
//   512.31kB  4.99% 80.03%   512.31kB  4.99%  regexp/syntax.(*compiler).inst
//   512.20kB  4.99% 85.02%   512.20kB  4.99%  runtime.malg
//   512.17kB  4.99% 90.02%   512.17kB  4.99%  github.com/gin-gonic/gin/render.writeContentType
//   512.04kB  4.99% 95.01%   512.04kB  4.99%  runtime.bgscavenge

// 結論：雖然go routine有加大 但是runtime.malg沒有增加
// 但只要有印log就可能增加＠＠？
// 2. 不管 -c 怎麼增加 goroutine 最多只到518
// Showing top 10 nodes out of 31
//       flat  flat%   sum%        cum   cum%
//  3597.01kB 36.89% 36.89%  3597.01kB 36.89%  bufio.NewWriterSize
//  1542.01kB 15.82% 52.71%  1542.01kB 15.82%  bufio.NewReaderSize (inline)
//  1537.69kB 15.77% 68.48%  1537.69kB 15.77%  runtime.allocm
//  1536.61kB 15.76% 84.24%  1536.61kB 15.76%  runtime.malg