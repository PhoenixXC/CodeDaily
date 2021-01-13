package main

import (
	"fmt"
	"os"
)

type Celsius float64
type Fahrenheit float64

// const boilingF = 212.0

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func main() {
	// var f = boilingF
	// var c = (f - 32) * 5 / 9
	// fmt.Printf("boiling point = %g'F or %g'C\n", f, c)

	c := FToC(212.0)
	fmt.Println(c.String())
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func CToF(c Celsius) Fahrenheit {
	// type 的类型转换
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 为类型定义行为，即类型的方法集
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func test() {
	f, err := os.Open("test")
	fmt.Println(f)
	fmt.Println(err)
	f, out := os.Create("test")
	// f, err = os.Create("test")
	fmt.Println(out)
}
