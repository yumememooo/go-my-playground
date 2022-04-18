package main

import (
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

//go tool pprof http://localhost:8888/debug/pprof/heap?debug
func demoFunc(i int) {
	fmt.Println("Hello World!,I am done.", i)
}

//這個範例用來試著創建大量的go，然後透過gin的pprof接口監看效能
func main() {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	pprof.Register(engine) // pprof monitor
	endpoint := fmt.Sprintf(":%d", 8888)
	go func() { //分析用
		err := engine.Run(endpoint)
		if err != nil {
			fmt.Println(err)
		}
	}()

	runTimes := 1048575 //給他一個大數目創建go

	for i := 0; i < runTimes; i++ {
		go demoFunc(i)
	}
	fmt.Printf("end call")
	select {} //避免退出
}
