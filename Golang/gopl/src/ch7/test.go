package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	// test nil

	var wr io.Writer = nil
	var wb *bytes.Buffer
	// 接口值分为：类型描述符 + 类型的值
	// 一个接口变量默认不赋值的情况下，类型描述符和类型的值都是 nil，也可以说是 nil 接口
	// 这里 wr2 ... = wb 类型描述符的值为 *byte.Buffer 类型的值为nil
	// 这种情况下 "包含nil指针的接口不是nil接口"
	var wr2 io.Writer = wb
	fmt.Println(wr == nil)  // true
	fmt.Println(wr2 == nil) // false

	var t0 io.Writer
	var t1 interface{}
	fmt.Println(t0 == t1) // true

	// var buf *bytes.Buffer
	// if debug {
	//     buf = new(bytes.Buffer)
	// }
	// f(nil)
	// f(buf)
}

//goland:noinspection GoUnhandledErrorResult
func f(out io.Writer) {
	// 这里如果声明了 var buf *bytes.Buffer，那么即便 buf 的值是 nil，但是他的类型已经被限定了，这样的一个包含了nil指针的接口，不等于nil
	// 对于 var buf io.Writer 而言，限定类型为接口，而接口的零值的类型和值都是nil

	// 当main函数调用函数f时，它给f函数的out参数赋了一个*bytes.Buffer的空指针，所以out的动态值是nil。
	// 然而，它的动态类型是*bytes.Buffer，意思就是out变量是一个包含空指针值的非空接口，所以out!=nil的结果依然是true。
	if out != nil {
		fmt.Println("call f func")
		out.Write([]byte("don!\n"))
	}
}
