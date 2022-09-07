// 测试函数的变长参数
package main

import "fmt"

func main() {
	f("hello", "hello", "dollar", "quality")
}

func f(prefix string, str ...string) {
	if len(str) == 0 {
		return
	}
	for _, s := range str {
		fmt.Println(s)
	}
}
