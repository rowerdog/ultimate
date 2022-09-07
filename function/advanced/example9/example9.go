// 处理err和panic的函数带有堆栈打印信息
package main

import (
	"fmt"
	"runtime"
)

func main() {
	if err := test(); err != nil {
		fmt.Println("main error", err)
	}
}

func catchPanic(err *error) {
	if r := recover(); r != nil {
		fmt.Println("Deferred panic : ", r)
		// 将当前堆栈信息存储在buf中，然后打印堆栈信息
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, false)
		fmt.Println("Stack trace: ", string(buf))
		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}

}
func test() (err error) {
	defer catchPanic(&err)
	fmt.Println("start test")
	err = mimicError("1")
	var p *int
	*p = 10
	fmt.Println("end test")
	return err
}

func mimicError(key string) error {
	return fmt.Errorf("mimic Error : %s", key)
}
