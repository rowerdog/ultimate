// 展示下接口的切片是如何工作的
package main

import "fmt"

type printer interface {
	print()
}

type canon struct {
	name string
}

func (c canon) print() {
	fmt.Printf("Printer name is %s\n", c.name)
}

type epson struct {
	name string
}

func (e *epson) print() {
	fmt.Printf("Printer name is %s\n", e.name)
}

func main() {
	c := canon{
		"佳能",
	}
	e := epson{
		"惠普",
	}
	// 声明一个pirnter接口的切片
	printers := []printer{
		c,
		&e,
	}

	for _, p := range printers {
		p.print()
	}
	c.name = "PROGRAF PRO-1000"
	e.name = "Home XP-4100"
	for _, p := range printers {
		p.print()
	}
}
