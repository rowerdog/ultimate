package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	// 声明一个admin变量，并且初始化

	a := admin{
		user: user{
			"aili",
			"aili@gmail.com",
		},
		level: "admin",
	}

	// 因为嵌入和类型提升，admin实现了notifier接口
	sendNotification(a)

}
