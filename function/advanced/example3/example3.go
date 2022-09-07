package main

import "fmt"

func mimicError(key string) error {
	return fmt.Errorf("mimic error : %s", key)
}

func main() {
	test()
}

func test() {
	fmt.Println("start test")
	err := mimicError("1")
	fmt.Printf("err addr : %v\n", &err)
	defer func() {
		if err != nil {
			fmt.Printf("Defer err: %s, addr: %v\n", err, &err)
		}
	}()
	err = mimicError("2")
	fmt.Println("end test")
}
