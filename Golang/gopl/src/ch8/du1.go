package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 设置一个channel，通过关闭channel来进行事件的广播
// 设置一个无缓存的channel，不向其发送数据，select会因为一直读取不到，如果channel已经关闭，对于一个已经关闭的channel发送数据
// 会导致 panic，读取数据会返回零值，这时 select 的这个 case 就可以执行下去
var done = make(chan struct{})

// 设置一个信号量，防止打开太多的文件，限制并发数量在一个合适的范围内
var sema = make(chan struct{}, 20)
var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	// 参数处理
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 定时打印统计信息
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 创建一个 goroutine 来读取输入，中止操作
	go func() {
		//goland:noinspection GoUnhandledErrorResult
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var nFiles, nBytes int64
	// break label，需要放到循环前，break 后不再执行这个循环
loop:
	for {
		select {
		case <-done:
			// 停止操作，释放资源
			fmt.Println("Stop operation!")
			for size := range fileSizes {
				nFiles++
				nBytes += size
			}
			printDiskUsage(nFiles, nBytes)
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nFiles++
			nBytes += size
		case <-tick:
			printDiskUsage(nFiles, nBytes)
		}
	}
	printDiskUsage(nFiles, nBytes)
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.2f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range directs(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func directs(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		//goland:noinspection GoUnhandledErrorResult
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
