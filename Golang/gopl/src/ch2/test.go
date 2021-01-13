package main

import "fmt"

func main() {
	fmt.Printf("%p\n", f())
	fmt.Printf("%p\n", f())
	fmt.Printf("%p\n", f())
	fmt.Printf("%p\n", f())
	fmt.Printf("%p\n", f())
	fmt.Printf("%p\n", f())
}

func f() *int {
	// v := 1
	// return &v
	return new(int)
}
