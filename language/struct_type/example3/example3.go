package main

import (
	"fmt"
	"unsafe"
)

type example struct {
	flag    bool    // 1
	counter int16   // 2
	pi      float32 // 4
}

func main() {
	e1 := &example{
		true,
		10,
		3.1415,
	}
	e1Next := &example{
		true,
		11,
		3.12,
	}

	fmt.Println(unsafe.Sizeof(e1))
	alignof := unsafe.Alignof(e1)
	fmt.Println(alignof)
	sizeBool := unsafe.Sizeof(e1.flag)
	offsetBool := unsafe.Offsetof(e1.flag)
	sizeInt := unsafe.Sizeof(e1.counter)
	offsetInt := unsafe.Offsetof(e1.counter)
	sizeFloat := unsafe.Sizeof(e1.pi)
	offsetFloat := unsafe.Offsetof(e1.pi)

	sizeBoolNext := unsafe.Sizeof(e1Next.flag)
	offsetBoolNext := unsafe.Offsetof(e1Next.flag)

	fmt.Println("对齐边界: ", alignof)
	fmt.Printf("flag: size: %d offset: %d addr: %v \n", sizeBool, offsetBool, &e1.flag)
	fmt.Printf("counter: size: %d offset: %d addr: %v \n", sizeInt, offsetInt, &e1.counter)
	fmt.Printf("pi: size: %d offset: %d addr: %v \n", sizeFloat, offsetFloat, &e1.pi)
	fmt.Printf("下一个: size: %d offset: %d addr: %v \n", sizeBoolNext, offsetBoolNext, &e1Next.flag)

	fmt.Println(unsafe.Sizeof(example{}))
}
