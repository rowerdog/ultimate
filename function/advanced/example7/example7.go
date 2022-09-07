package main

import "fmt"

func mimicError(key string) error {
	return fmt.Errorf("mimic Error : %s", key)
}

func main() {
	if err := test(); err != nil {
		fmt.Println("Main err", err)
	}
}

func test() (err error) {
	fmt.Println("start test")

	// 处理panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("defer panic", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	// 处理err
	defer func() {
		if err != nil {
			fmt.Println("defer err", err)
		}
	}()
	err = mimicError("1")
	panic("Mimic panic")
}
