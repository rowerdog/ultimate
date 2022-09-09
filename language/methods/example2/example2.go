package main

import "fmt"

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Println("My name is", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "is age", d.age)
}

func main() {
	d := data{
		"bill",
		13,
	}
	d.displayName()
	d.setAge(18)

	f1 := d.displayName
	f1()
	d.name = "mask"
	// d.name没有改变
	f1()

	// d.name改变了
	f2 := d.setAge
	f2(25)
	d.name = "elon"
	f2(30)
}
