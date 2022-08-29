package main

import "fmt"

// 展示基本变量的声明和初始化、零值
func main() {

	//var 声明变量，变量都设置为他们的零值
	var a int
	var b string
	var c float64
	var d bool

	fmt.Println(c)
	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b int \t %T [%v]\n", b, b)
	fmt.Printf("var c int \t %T [%v]\n", c, c)
	fmt.Printf("var d int \t %T [%v]\n", d, d)

	//声明变量并且初始化

	aa := 10
	bb := "hello"
	cc := 3.1415926
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.1415926 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n", dd, dd)

	//指定类型并且进行转换
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) \t %T [%v]\n", aaa, aaa)

}
