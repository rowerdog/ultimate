package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("有错误")
		}
	}()
	gop(2)
	fmt.Println("程序结束")

}

func gop(i int) {
	if i != 1 {
		panic("int 的值不是 1")
	}
}

func goa(i int) error {
	if i != 2 {
		return errors.New("i 的值不是2")
	}
	return fmt.Errorf("i的值 %s 是2")
}
