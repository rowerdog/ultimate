// panic、defer、recover的联合用法
package main

import "fmt"

func main() {
	test()
}

func test() error {
	defer func() {
		fmt.Println("Start Panic Defer")

		if r := recover(); r != nil {
			fmt.Println("Defer Panic:", r)
		}
	}()

	fmt.Println("Start Test")
	panic("Mimic Panic")
}
