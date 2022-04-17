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
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
				// if i == 10 {
				// 	time.Sleep(10 * time.Second)
				// }
				if i <= 100 {
					// fmt.Println("add eventChan:", i)
					b.SendData(i, eventChan) //塞一個event data
				}

			}()
		}
	}()

	//------------------------------------
	errs := make(chan error, 3)
	listenForInterrupt(errs)
	StartHttpServer(errs, 6060)
	c := <-errs
	fmt.Println("terminating:", c)
}

func (b *Event) SendData(amount int, eventChan chan int) {
	//https: //segmentfault.com/a/1190000017537297 可以實現無阻塞
	select {
	case eventChan <- amount:
		fmt.Println("add eventChan:", amount, "in")
		// return nil
	default:
		fmt.Println("default no:", amount)
		go waitProcess(amount, eventChan)
		// return errors.New("channel blocked, can not write")
	}

}

//這一個範例使用select+default 實現無阻塞讀取，但需要注意default條件處理，避免丟失。
//假設default開了一個等一下再處理event的go
//跑一下這範例會發現out:eventChan Size: 0  value 100  process done
//竟然所有的範例都跑完了!!!  看起來好像可以利用這樣處理, 但到底要等多久呢???

func waitProcess(amount int, eventChan chan int) {
	time.Sleep(20 * time.Second)
	eventChan <- amount
	fmt.Println("Sleep add eventChan:", amount, "in")

}

func (b *Event) ListenChanHandler(eventChan chan int) { //消費方
	go func() {
		for {
			select {
			case e := <-eventChan:
				fmt.Println("out:eventChan Size:", len(eventChan), " value", e, " process done")
				//dosometing 5s
				time.Sleep(5 * time.Second)
			}
		}
	}()
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

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
