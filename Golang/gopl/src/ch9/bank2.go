package main

import (
	"fmt"
	"sync"
)

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		deposit2(200)
		fmt.Println("=", balance2())
		wg.Done()
	}()
	go func() {
		deposit2(100)
		wg.Done()
	}()

	wg.Wait()
}

func deposit2(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func balance2() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
