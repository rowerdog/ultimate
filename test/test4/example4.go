package main

import (
	"fmt"
	"os"
)

func main() {
	test()
}

func test() {
	file, err := os.Open("1.txt")

	fmt.Println(file)
	defer func() {
		err2 := file.Close()
		fmt.Println(err2)
	}()

	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
}
