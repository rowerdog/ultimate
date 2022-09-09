// 演示下类型嵌入到另一个类型如何做
// 演示下内部类型的类型提升
package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // embedding，是嵌入类型
	level string
}

func main() {
	a := admin{
		user: user{
			"bill",
			"bill@gmail.com",
		},
		level: "high",
	}

	a.user.notify()
	// 通过这种方法，内部类型可以提升到外部
	a.notify()
}
