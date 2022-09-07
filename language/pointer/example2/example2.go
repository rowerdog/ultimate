// 数据跨程序边界共享时，每个函数或者goroutine都有自己的数据副本
package main

import "fmt"

func main() {

	count := 10

	ptr1 := &count
	ptr2 := &count

	inc(ptr1)
	fmt.Println(&ptr1)
	fmt.Println(&ptr2)

}

func inc(count *int) {
	*count++
	fmt.Printf("count指针变量的值是: %v\t count指针变量的地址是: %v\t count指针指向的值的大小: %d\t\n", count, &count, *count)
}
