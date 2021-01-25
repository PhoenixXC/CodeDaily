package memo1

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestMemo_Get(t *testing.T) {
	var waitGroup sync.WaitGroup
	incomingURLs := []string{
		"http://cs.xiyoulinux.org",
		"http://xiyoulinux.org",
		"http://blog.xiyoulinux.org",
		"http://xiyoulinux.org",
		"http://cs.xiyoulinux.org",
		"http://join.xiyoulinux.org",
		"http://blog.xiyoulinux.org",
		"http://xiyoulinux.org",
		"http://cs.xiyoulinux.org",
	}
	m := New(httpGetBody)
	// 本来是个函数，为了简单用数组处理
	for _, url := range incomingURLs {
		waitGroup.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			defer waitGroup.Done()
		}(url)
	}
	waitGroup.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
