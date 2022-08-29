package main

import (
	"fmt"
	"unsafe"
)

type foo struct {
	f1 int8
	f2 int64
}

func main() {

	fmt.Println(unsafe.Alignof(bool(true)))

	foo1 := &foo{
		1,
		2,
	}

	fmt.Println(unsafe.Sizeof(foo1.f1), unsafe.Offsetof(foo1.f1))
	fmt.Println(unsafe.Sizeof(foo1.f2), unsafe.Offsetof(foo1.f2))
	
}
