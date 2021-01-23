package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Printf("%T\n", f)
	// c := w.(*bytes.Buffer)
	// fmt.Printf("%v\n", c)
}

// 在最简单的形式中，一个类型分支像普通的switch语句一样，
// 它的运算对象是x.(type)——它使用了关键词字面量type
func sqlQuote(x interface{}) string {
	// 重用字面量
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return x // (not shown)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
