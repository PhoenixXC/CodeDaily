package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown. Press ENTER to abort.")

	// create abort channel
	abort := make(chan struct{})
	go func() {
		//goland:noinspection GoUnhandledErrorResult
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	// 这样会导致goroutine泄露
	// tick := time.Tick(1 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	for countdown := 3; countdown > 0; countdown-- {
		select {
		case <-ticker.C:
			fmt.Printf("\rLauch countdown [%02d]", countdown)
		case <-abort:
			fmt.Println("\rLaunch aborted!")
			return
		}
	}
	ticker.Stop()
	launch()
}

func launch() {
	fmt.Println("\rLaunch!")
}
