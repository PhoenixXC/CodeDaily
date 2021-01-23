package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	filenames := []string{
		"a", "b", "c", "d", "e", "f",
	}
	makeThumbnails(filenames)
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	// 计算器，goroutine启动时+1，退出时-1
	// 这个计数器可以在多个goroutine操作时做到安全并且提供在其减为零之前一直等待的方法
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			_, err := imageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info := rand.Int63()
			sizes <- info
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

// 可以给一个函数的返回值指定名字。如果指定了一个返回值的名字，则可以视为在该函数的第一行中定义了该名字的变量。
func makeThumbnailsDealErrPlus(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbFile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbFile, it.err = imageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbFile)
	}

	return thumbfiles, nil
}

// BUG: 当它遇到第一个非nil的error时会直接将error返回到调用方，使得没有一个goroutine去排空errors channel。这样剩下的worker goroutine
// 在向这个channel中发送值时，都会永远地阻塞下去，并且永远都不会退出。这种情况叫做goroutine泄露（§8.4.4），
// 可能会导致整个程序卡住或者跑出out of memory的错误。
func makeThumbnailsDealErr(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := imageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

func makeThumbnails(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			//goland:noinspection GoUnhandledErrorResult
			imageFile(f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

func imageFile(infile string) (string, error) {
	fmt.Printf("deal file [%s]\n", infile)
	time.Sleep(1 * time.Second)
	return infile, nil
}
