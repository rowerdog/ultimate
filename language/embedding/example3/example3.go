package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	*user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	a := admin{
		user: &user{
			"alice",
			"alice@gmail.com",
		},
		level: "admin",
	}

	a.user.notify()
	a.notify()

	sendNotification(&a)

}
