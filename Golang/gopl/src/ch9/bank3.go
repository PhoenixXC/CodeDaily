package main

import (
	"fmt"
	"sync"
)

var (
	mu         sync.Mutex
	balanceVar int
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		deposit3(200)
		fmt.Println("=", balance2())
		wg.Done()
	}()
	go func() {
		deposit3(100)
		wg.Done()
	}()

	wg.Wait()
}

func deposit3(amount int) {
	mu.Lock()
	balanceVar = balanceVar + amount
	mu.Unlock()
}

func balance3() int {
	defer mu.Unlock()
	mu.Lock()
	return balanceVar
}
