package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

type Circle struct {
	radius float64
}

func Distance(p, q Point) float64 {
	fmt.Println("Distance defined as function")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	fmt.Println("Distance defined as method")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	fmt.Println("Distance defined as Path([]Point) method")
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (c Circle) changeRadius(newRadius float64) {
	c.radius = newRadius
}

func (c *Circle) realChangeRadius(newRadius float64) {
	c.radius = newRadius
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	path := Path{
		p, q,
	}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Println(path.Distance())
	circle := Circle{10.0}
	// 不会实际修改值，修改的只是拷贝的副本
	circle.changeRadius(12.0)
	fmt.Println(circle.radius)
	circle.realChangeRadius(20.0)
	fmt.Println(circle.radius)
	// 方法值和表达式
	// 返回一个函数，这个函数已经绑定了接收器变量
	distanceFromP := p.Distance
	fmt.Printf("%T\n", distanceFromP)
	fmt.Println(distanceFromP(q))

}
