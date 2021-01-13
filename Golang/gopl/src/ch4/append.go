package main

func main() {

}

func appendInt(x []int, y int) []int {
	var z []int
	// 新的长度
	zlen := len(x) + 1

	if zlen <= cap(x) {
		// 如果容量够存放新的数据
		z = x[:zlen]
	} else {
		// 容量不够，分配新的底层数组
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
