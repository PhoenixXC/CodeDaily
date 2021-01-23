package main

import (
	"fmt"
	"syscall"
)

var errors = []string{
	'k': "Hello",
}

func main() {
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}
