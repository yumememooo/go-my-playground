package main

import (
	"fmt"
	"os"
	"os/signal"

	"syscall"
	"time"

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
	fmt.Printf("terminating: %v", c)
}

func (b *Event) SendData(amount int, eventChan chan int) {
	eventChan <- amount
	fmt.Println("add eventChan:", amount, "in")
}

func (b *Event) ListenChanHandler(eventChan chan int) {
	go func() {
		for {
			select {
			case e := <-eventChan:
				fmt.Println("out:eventChan Size:", len(eventChan), " value", e, " process done")

				//dosometing 5s
				time.Sleep(3 * time.Second)
			}
		}
	}()
}

type Event struct {
	// balance int useless
	// mux sync.Mutex
}

func StartHttpServer(errChan chan error, port int) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	// loadRoutes(engine)
	// pprof.Register(engine)                      // pprof monitor
	engine.GET("/metrics", prometheusHandler()) // prometheus monitor
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
