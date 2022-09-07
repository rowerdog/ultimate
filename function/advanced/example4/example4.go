package main

import "fmt"

func mimicError(key string) error {
	return fmt.Errorf("mimic Error : %s", key)
}

func main() {
	if err := test(); err != nil {
		fmt.Println("Main error:", err)
	}
}

func test() (err error) {
	fmt.Println("start test")

	err = mimicError("1")

	defer func() {
		if err != nil {
			fmt.Println("Defer error:", err)
		}
	}()

	// 重新调用mimicError()，err重新赋值
	err = mimicError("2")
	fmt.Println("end test")
	return
}
