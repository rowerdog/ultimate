package main

import "fmt"

// 声明一个notify行为的接口
type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending email to %s<%s>\n")
}

func main() {
	u := user{
		"bill",
		"bill@gmail.com",
	}
	sendNotification(&u)

}

func sendNotification(n notifier) {
	n.notify()
}
