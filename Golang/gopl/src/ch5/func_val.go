package main

import "fmt"

func main() {

	f := square
	var f2 func(int) int
	fmt.Println(f(3))
	f2 = negative
	fmt.Println(f2(3))

}

func square(n int) int   { return n * n }
func negative(n int) int { return -n }
