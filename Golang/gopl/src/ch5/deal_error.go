package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//goland:noinspection GoUnhandledErrorResult
func main() {

	url := "http://gopl.io"

	// 第三种策略：若错误发生后无法继续，输出错误信息直接结束程序
	// 只应该在 main 中执行，库函数应仅向上传播错误，除非是BUG
	if err := retryFailedOp(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
		// 或者使用：
		// log.Fatalf("Site is down: %v\n", err)
	}

	// 第四种：只输出错误信息

	// 第五种：忽略错误

}

// spreadError 错误处理策略一：传播错误
//
// 函数中某个子程序的失败，变成该函数的失败
func spreadError(url string) (string, error) {
	// ...

	var data string
	_, err := http.Get(url)
	if err != nil {
		return "", err
	}

	//  ...

	return data, nil
}

// retryFailedOp 错误处理策略二：重试
//
// 若错误偶然发生，可以重新尝试失败的操作，当然需要限制重试的时间间隔和次数
func retryFailedOp(url string) error {
	const timeout = 1 * time.Minute
	// 终止重试的时间
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		// 执行操作
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		// 随着失败次数的增多，等待下一次重试的时间越长
		sleepTime := time.Second << uint(tries)
		if time.Now().Add(sleepTime).After(deadline) {
			// 如果发现继续等下去也不会执行这个操作，就可以直接结束了
			break
		}
		time.Sleep(sleepTime)
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
