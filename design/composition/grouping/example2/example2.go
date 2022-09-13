// 这是一个使用组合和接口的例子。 这是
// 我们想在 Go 中做的事情。 我们将按以下方式对常见类型进行分组
// 他们的行为而不是他们的状态。 这种模式确实
// 在 Go 程序中提供良好的设计原则。
package main

import "fmt"

// 单个方法的接口，后缀是er
// 提供了通用的行为接口，如果有人想加入进来
// 需要遵守规则
type Speaker interface {
	Speak()
}

type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

// 实现了接口
func (d *Dog) Speak() {
	fmt.Printf(
		"Woof! My name is %s, it is %t I am a mammal with a pack factor of %d.\n",
		d.Name,
		d.IsMammal,
		d.PackFactor,
	)
}

type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

// 实现Speaker接口
func (c *Cat) Speak() {
	fmt.Printf(
		"Meow! My name is %s, it is %t I am a mammal with a climb factor of %d.\n",
		c.Name,
		c.IsMammal,
		c.ClimbFactor,
	)
}

func main() {
	speakers := []Speaker{
		&Dog{
			"旺财",
			true,
			5,
		},
		&Cat{
			"加菲",
			true,
			6,
		},
	}

	for _, speaker := range speakers {
		speaker.Speak()
	}
}
