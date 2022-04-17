package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	b := &Bank{}
	n := 1000
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			b.Deposit(100) //模擬多個併發同時存款100
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(b.ReadBalance()) //印出存錢的總額
}

type Bank struct {
	balance int
}

func (b *Bank) Deposit(amount int) {
	b.balance += amount
}

func (b *Bank) ReadBalance() int {
	return b.balance
}
