// 测试函数的变长参数是空接口
package main

import "fmt"

func main() {
	typeCheck("hello", 1, struct {
		name string
	}{"hello"})
}

func typeCheck(str ...interface{}) {
	for _, val := range str {
		switch val.(type) {
		case string:
			fmt.Println("string类型")
		case int:
			fmt.Println("int类型")
		case struct{ name string }:
			fmt.Println("struct类型")
		}
	}
}
