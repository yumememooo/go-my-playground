package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Bank struct {
	balance int
	mux     sync.RWMutex
}

func (b *Bank) Deposit(amount int) {
	b.mux.Lock()
	time.Sleep(time.Second) // spend 1 second
	b.balance += amount
	b.mux.Unlock()

}

func (b *Bank) Balance() (balance int) {
	b.mux.RLock()
	// b.mux.Lock()

	time.Sleep(time.Second) // spend 1 second
	balance = b.balance
	b.mux.RUnlock()
	// b.mux.Unlock()
	return
}

func main() {
	var wg sync.WaitGroup
	b := &Bank{}
	start := time.Now()
	n := 5
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			b.Deposit(1000)
			log.Printf("Write: deposit amonut: %v", 1000)
			wg.Done()
		}()
	}
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			log.Printf("Read: balance: %v", b.Balance())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

//
// D:\go\src\demo-go\02-go-race-fix-readlock>go run main.go
// 2021/08/04 11:58:41 Write: deposit amonut: 1000
// 2021/08/04 11:58:42 Read: balance: 1000
// 2021/08/04 11:58:42 Read: balance: 1000
// 2021/08/04 11:58:42 Read: balance: 1000
// 2021/08/04 11:58:42 Read: balance: 1000
// 2021/08/04 11:58:42 Read: balance: 1000
// 2021/08/04 11:58:43 Write: deposit amonut: 1000
// 2021/08/04 11:58:44 Write: deposit amonut: 1000
// 2021/08/04 11:58:45 Write: deposit amonut: 1000
// 2021/08/04 11:58:46 Write: deposit amonut: 1000
