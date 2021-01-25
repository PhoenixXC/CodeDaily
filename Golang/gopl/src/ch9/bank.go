package main

import (
	"fmt"
	"sync"
)

type Bank struct {
	balance int
}

func (b Bank) Deposit(amount int) {
	b.balance = b.balance + amount
}

func (b Bank) Balance() int {
	return b.balance
}

var bank = Bank{100}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		wg.Done()
	}()
	go func() {
		bank.Deposit(100)
		wg.Done()
	}()

	wg.Wait()
}
