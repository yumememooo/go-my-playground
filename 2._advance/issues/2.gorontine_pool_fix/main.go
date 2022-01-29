package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/panjf2000/ants/v2"
)

func demoFunc(i interface{}) {
	fmt.Println("Hello World!,I am done.", i)
}
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
	defer ants.Release()

	runTimes := 1048575 //百萬次

	var wg sync.WaitGroup
	options := ants.Options{}
	options.ExpiryDuration = time.Duration(10) * time.Second
	options.Nonblocking = true
	options.PreAlloc = true
	// poolOpts, _ := NewPool(1, WithOptions(options))
	// set 100000 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(100000, func(i interface{}) {
		demoFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()

	fmt.Printf("running goroutines: %d\n", p.Running())
	p.Release() //不須要跑了 強制先做release

	select {} //避免退出
}

//>go tool pprof http://localhost:8888/debug/pprof/heap?debug
//(pprof) top
//      flat  flat%   sum%        cum   cum%
//  41.02MB 43.19% 43.19%    41.02MB 43.19%  runtime.malg
