// 展示结构体的声明、初始化、零值
package main

import "fmt"

//
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	// 声明一个example的结构体，并且设为零值
	var e1 example

	//打印e1每个字段的值
	fmt.Printf("%+v\n", e1)

	//声明一个结构体，并且使用字面量初始化

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.145,
	}

	//输出每个字段的值
	fmt.Println("Flag: ", e2.flag)
	fmt.Println("counter: ", e2.counter)
	fmt.Println("pi: ", e2.pi)
	
}
