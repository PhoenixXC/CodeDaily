package main

import (
	"fmt"
)

func main() {
	var uVar int8 = -2
	fmt.Println(uVar << 1)
	fmt.Println(uVar >> 1)

	var apple int32 = 1
	var oranges int16 = 2
	var compote int = int(apple) + int(oranges)
	fmt.Println(compote)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	unicode := '国'
	fmt.Printf("%d %[1]c %[1]q\n", unicode)

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
}
