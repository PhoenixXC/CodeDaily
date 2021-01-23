package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	_, _ = c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	_, _ = fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
