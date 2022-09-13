// 这是一个错误的模板展示继承

package main

import "fmt"

// Animal 作为所有动物的基类
type Animal struct {
	Name     string
	IsMammal bool
}

// Speak 提供了一个所有动物说话的模板
func (a *Animal) Speak() {
	fmt.Printf("UGH My name is %s it is %t I am a mammal\n", a.Name, a.IsMammal)
}

// Dog Dog有Animal有的一切，并且有且只有Dog有的属性
type Dog struct {
	Animal
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Printf("Woof! My name is %s, it is %t I am a mammal with a pack factor of %d.\n",
		d.Name,
		d.IsMammal,
		d.PackFactor)
}

type Cat struct {
	Animal
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Printf(
		"Meow! My name is %s, it is %t I am a mammal with a climb factor of %d.\n",
		c.Name,
		c.IsMammal,
		c.ClimbFactor,
	)
}

func main() {
	animals := []Animal{
		Dog{
			Animal{
				"旺财",
				true,
			},
			1,
		},
		Cat{
			Animal{
				"加菲",
				true,
			},
			2,
		},
	}

	for _, animal := range animals {
		animal.Speak()
	}
}
