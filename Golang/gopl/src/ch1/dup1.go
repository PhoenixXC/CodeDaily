// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 类似 Java 中的 HashMap
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// 读取输入
	for input.Scan() {
		// map 中不存在的会被自动计算为 0
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
