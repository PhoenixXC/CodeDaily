package main

import (
	"fmt"
	"os"
)

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

	vals := []int{0, 1, 2, 3, 4, 5}
	// 需要加上...
	fmt.Println(sum(vals...))

}

func squares() func() int {
	var x int
	// 函数值不仅仅是一串代码，还记录了状态。
	// 在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，
	// 这意味着匿名函数和squares中，存在变量引用。
	// 这就是函数值属于引用类型和函数值不可比较的原因。
	// Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。
	return func() int {
		x++
		return x * x
	}
}

func tempDirs() []string {
	return []string{
		"/tmp/a.txt",
		"/tmp/b.txt",
		"/tmp/c.txt",
		"/tmp/d.txt",
	}
}

//goland:noinspection GoUnhandledErrorResult
func delDir() {
	var rmDirs []func()
	for _, d := range tempDirs() {
		// 匿名函数中记录的不是变量的值，而是变量的地址
		// 若使用原来的局部变量 d，那么最后每个函数中记录的是最后一次
		// 循环时 d 的地址，是一个值，那么每次删除的也是相同的目录；
		// dir := d
		// 也可以声明一个相同名称的局部变量，这样是合法的
		d := d
		os.MkdirAll(d, 0755)
		rmDirs = append(rmDirs, func() {
			os.RemoveAll(d)
		})
	}
	for _, rmdir := range rmDirs {
		rmdir()
	}
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
