// 测试函数的命名返回值
package main

import "fmt"

func main() {
	x, y := f()
	fmt.Println(x, y)
}

func f() (x int, y int) {
	x = 1
	y = 2
	return
}
