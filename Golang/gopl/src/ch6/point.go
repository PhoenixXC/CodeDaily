package main

import (
	"fmt"
	"image/color"
)

type xPoint struct {
	X, Y float64
}

func (x xPoint) echo() {
	fmt.Println("xPoint method")
}

func (x xPoint) echoWith(p xPoint) {
	fmt.Println("xPoint method with param")
}

type xColoredPoint struct {
	xPoint
	Color color.RGBA
}

func main() {
	var cp xColoredPoint
	cp.X = 1
	fmt.Println(cp.xPoint.X)
	cp.xPoint.Y = 2
	fmt.Println(cp.Y)
	cp.echo()
	// Go struct嵌套是 has a 不是 is a
	cp.echoWith(xColoredPoint{
		xPoint: xPoint{1, 2},
		Color:  nil,
	}.xPoint)
}
