package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	b := &Bank{}
	n := 10000
	wg.Add(n)
	valueCh := make(chan int)
	go b.ListenChanHandler(valueCh)
	for i := 1; i <= n; i++ {
		go func() {
			b.Deposit(100, valueCh) //模擬多個併發同時存款
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(b.ReadBalance()) //印出
	fmt.Println("done")
}

type Bank struct {
	balance int
	mux     sync.Mutex
}

func (b *Bank) Deposit(amount int, valueCh chan int) {

	valueCh <- amount
}

func (b *Bank) ListenChanHandler(valueCh chan int) {
	go func() {
		for {
			select {
			case e := <-valueCh:
				b.mux.Lock()
				b.balance = b.balance + e
				b.mux.Unlock()
			}
		}
	}()
}

func (b *Bank) ReadBalance() (balnce int) {
	b.mux.Lock()
	balance := b.balance
	b.mux.Unlock()
	// fmt.Println("balance:", balance)
	return balance
}
