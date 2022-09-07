// sample program show defer、panic、recover
package main

import "fmt"

// 模拟发生错误的程序
func mimicError(key string) error {
	return fmt.Errorf("mimicError: %s", key)
}

func main() {

	fmt.Println("start test")

	err := mimicError("1")

	defer func() {
		fmt.Println("start defer")
		if err != nil {
			fmt.Println("Defer Error:", err)
		}
	}()

	fmt.Println("end test")
}
