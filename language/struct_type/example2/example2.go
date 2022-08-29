package main

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	var e1 example

	fmt.Printf("%+v\n", e1)

	//声明一个匿名的结构体
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		true,
		10,
		3.1415,
	}

	fmt.Printf("%+v\n", e2)

}
