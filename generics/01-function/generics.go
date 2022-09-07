package main

import (
	"fmt"
	"reflect"
)

// 只能针对特定类型进行输出
func printNumbers(numbers []int) {
	fmt.Println("print Numbers")
	for _, number := range numbers {
		fmt.Println(number, " ")
	}
}

func printStrings(strings []string) {
	fmt.Println("print Strings")
	for _, str := range strings {
		fmt.Println(str, " ")
	}
}

// 使用空接口进行断言
func printAssert(v interface{}) {
	fmt.Println("assert")
	switch list := v.(type) {
	case []int:
		for _, number := range list {
			fmt.Println(number, " ")
		}
	case []string:
		for _, str := range list {
			fmt.Println(str, " ")
		}
	}
}

// 使用反射

func printReflect(v interface{}) {
	fmt.Println("Reflect")
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < val.Len(); i++ {
		fmt.Println(val.Index(i).Interface(), " ")
	}

}

func print[T any](slice []T) {
	fmt.Println("Generics")
	for _, v := range slice {
		fmt.Println(v, " ")
	}
}

func main() {
	numbers := []int{1, 2, 3}
	print[int](numbers)
	printAssert(numbers)
	printReflect(numbers)
}
