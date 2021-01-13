package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // 像素大小
	cells         = 100                 // 网格单元数
	xyrange       = 100                 // xy 范围
	xyscale       = width / 2 / xyrange // xy单元的单位像素
	zscale        = height * 0.4        // z单元的单位像素
	angle         = math.Pi / 6         // xy角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
	}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
