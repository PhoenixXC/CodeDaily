package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			done <- i + 1
			time.Sleep(1 * time.Second)
		}
		// done关闭后可以读取零值
		close(done)
	}()

	for {
		select {
		case x := <-done:
			if x != 0 {
				fmt.Println("Wakeup...")
			} else {
				fmt.Println("done!")
				return
			}
		}
	}

}
