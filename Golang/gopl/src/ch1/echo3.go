package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sepChar := "-----------------"
	fmt.Println("ELF name: " + os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(sepChar)
	fmt.Println(os.Args[1:])

	fmt.Println(sepChar)
	for index, arg := range os.Args {
		fmt.Println("index: " + strconv.Itoa(index) + ", arg: " + arg)
	}
}
