package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	b := &Bank{}
	n := 10
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			b.Deposit(100) //模擬多個併發同時存款
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(b.ReadBalance()) //印出存錢的總額
}

type Bank struct {
	balance int
	mux     sync.Mutex
}

func (b *Bank) Deposit(amount int) {
	b.mux.Lock()
	b.ReadBalance()     //錯誤 裡面也有lock
	b.balance += amount //保護
	b.mux.Unlock()
}

func (b *Bank) ReadBalance() (balnce int) {
	b.mux.Lock()
	balance := b.balance
	b.mux.Unlock()
	return balance
}
