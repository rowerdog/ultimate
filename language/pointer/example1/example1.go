// 演示下go的值传递
package main

import "fmt"

func main() {

	count := 10
	fmt.Printf("变量的值是:%d 变量的地址是: %v \n", count, &count)
	inc(count)
	fmt.Printf("变量的值是:%d 变量的地址是: %v \n", count, &count)

}

func inc(inc int) {
	inc++
	fmt.Printf("变量的值是:%d 变量的地址是: %v \n", inc, &inc)
}
