package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	// TODO fmt 并发安全？
	// time.Sleep(99 * time.Millisecond)
}

func spinner(delay time.Duration) {
	// 主函数返回时，所有的goroutine都会被直接打断，程序退出。
	// 通过goroutine之间的通信来让一个goroutine请求其它的goroutine，
	// 并让被请求的goroutine自行结束执行
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
