package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func demoFunc(i int) {
	//	fmt.Println("Hello World!,I am done.")  //如果直接印 竟然會造成堆積runtime.malg＠＠
	//log.Println("Hello World!,I am done.")  //換個方法印 949990	goroutine 一樣堆積 201.58MB 73.29%   runtime.malg
	//sugar.Debugf("debug message") //658392	goroutine // 152.56MB 23.67% 80.75%   152.56MB 23.67%  runtime.malg
	// defer sugar.Sync() 不要這樣做 會變超級超級超級慢
	// 如果sugur log level變成info 不印的話，一樣會造成塞住
	// 但是會多了bufferpool
	//265.25MB 56.07% 56.07%   265.25MB 56.07%  go.uber.org/zap/internal/bufferpool.init.func1
  //102.54MB 21.67% 77.74%   102.54MB 21.67%  runtime.malg
  // 如果main程式有做sugar.Sync() 還是卡在這邊
	// 493.96MB 56.29% 56.29%   493.96MB 56.29%  go.uber.org/zap/internal/bufferpool.init.func1
  // 194.58MB 22.17% 78.46%   194.58MB 22.17%  runtime.malg

	//temp.Store(i, "") //不要印 只存到其他map 從頭到尾只有 5 goroutine 沒有堆積
	//不做任何事 從頭到尾只有 5 goroutine 沒有堆積
	//time.Sleep(1 * time.Second) //如果只處理一秒的情況？  //77710	goroutine 會堆積 180.57MB 77.51% 77.51%   180.57MB 77.51%  runtime.malg
	time.Sleep(50 * time.Microsecond) //如果只處理50us的情況？ 856 goroutine 很快結束 4097.62kB 61.49% 61.49%  4097.62kB 61.49%  runtime.malg 有時沒堆積demoFunc
}

// 結論：所以不要在大量創建go的情境下去做印log的行為嗎？ 另外很快就結束的例子一樣不會同時花上太多的goＮum，也不會造成記憶體堆積

var temp sync.Map
var runGoNun = 0

//這個範例用來試著創建大量的go，然後透過gin的pprof接口監看效能
func main() {
	gin.SetMode(gin.ReleaseMode)
	InitLogger()
	defer sugar.Sync()

	engine := gin.New()
	pprof.Register(engine) // pprof monitor
	endpoint := fmt.Sprintf(":%d", 8888)
	go func() { //分析用
		err := engine.Run(endpoint)
		if err != nil {
			fmt.Println(err)
		}
	}()
	go runBigGo()

	fmt.Printf("end call")

	time.Sleep(120 * time.Second)
	temp.Range(p)
	log.Printf("runGoNun:%d", runGoNun) //runGoNun:1048575
	select {}                           //避免退出
}
func p(key interface{}, value interface{}) bool {
	runGoNun = runGoNun + 1
	return true
}

func runBigGo() {
	runTimes := 1048575 //給他一個大數目創建go

	for i := 0; i < runTimes; i++ {
		go demoFunc(i)
	}
}

// 418981	goroutine
// 27	heap
//結束：
// 5	goroutine
// 38	heap
//go tool pprof http://localhost:8888/debug/pprof/heap
// Showing top 10 nodes out of 32
//       flat  flat%   sum%        cum   cum%
//   177.07MB 52.06% 52.06%   177.07MB 52.06%  runtime.malg
//    96.52MB 28.38% 80.43%    96.52MB 28.38%  fmt.glob..func1
//    40.50MB 11.91% 92.34%    40.50MB 11.91%  internal/poll.runtime_Semacquire
//    17.50MB  5.15% 97.49%    17.50MB  5.15%  fmt.(*buffer).write (inline)
//     3.55MB  1.04% 98.53%     3.55MB  1.04%  runtime.allgadd
//     2.50MB  0.73% 99.26%   157.02MB 46.16%  main.demoFunc
//        2MB  0.59% 99.85%        2MB  0.59%  runtime.allocm
//          0     0% 99.85%    17.50MB  5.15%  fmt.(*fmt).fmtInteger
//          0     0% 99.85%    17.50MB  5.15%  fmt.(*fmt).pad
//          0     0% 99.85%    17.50MB  5.15%  fmt.(*pp).doPrintln
var sugar *zap.SugaredLogger

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugar = logger.Sugar()
}
