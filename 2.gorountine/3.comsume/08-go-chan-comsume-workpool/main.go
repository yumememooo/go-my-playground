package main

import (
	"fmt"
	"os"
	"os/signal"

	"sync"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	// "github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"go.uber.org/zap"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	b := &Event{}
	eventChan := make(chan int, 10) //buffer size 100
	go b.ListenChanHandler(eventChan)
	go func() {
		i := 0
		for {
			time.Sleep(100 * time.Millisecond)
			go func() { //每個100ms不斷接收event進來
				i = i + 1
				// fmt.Println("add eventChan:", i)
				b.SendData(i, eventChan) //塞一個event data
			}()
		}
	}()

	//------------------------------------
	errs := make(chan error, 3)
	listenForInterrupt(errs)
	StartHttpServer(errs, 6060)
	c := <-errs
	zap.S().Warnf("terminating: %v", c)
}

func (b *Event) SendData(amount int, eventChan chan int) {
	// tt := time.Now()
	eventChan <- amount
	fmt.Println("add eventChan:", amount, "in")
	// fmt.Println("time", time.Since(tt))
}

func (b *Event) ListenChanHandler(eventChan chan int) { //消費方
	for w := 1; w <= 3; w++ {
		go worker(w, eventChan)
	}
	// go func() {
	// 	for {
	// 		select {
	// 		case e := <-eventChan: //接收channel
	// 			fmt.Println("out:eventChan Size:", len(eventChan), " value", e, " process done")
	// 			//dosometing 5s
	// 			time.Sleep(5 * time.Second)
	// 		}
	// 	}
	// }()
}
func worker(id int, jobs <-chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(3 * time.Second)
		fmt.Println("worker", id, "finished job", j)
		fmt.Println("out:eventChan Size:", len(jobs), " value", j, " process done")

	}
}

type Event struct {
	// balance int useless
	mux sync.Mutex
}

func StartHttpServer(errChan chan error, port int) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	// loadRoutes(engine)
	pprof.Register(engine) // pprof monitor
	// profiler.Start(profiler.Config{
	// 	ApplicationName: "simple.golang.app",
	// 	ServerAddress:   "http://localhost:4040",
	// })
	endpoint := fmt.Sprintf(":%d", port)
	go func() { errChan <- engine.Run(endpoint) }()
}

func listenForInterrupt(errChan chan error) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGALRM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
}

// func prometheusHandler() gin.HandlerFunc {
// 	h := promhttp.Handler()

// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }
