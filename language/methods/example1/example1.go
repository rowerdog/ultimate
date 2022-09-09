package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	u := user{
		"bill",
		"bill@gmail.com",
	}
	u.notify()
	u.changeEmail("dialing@gmail.com")
	u.notify()

	(*user).changeEmail(&u, "mask@tesla.com")
	user.notify(u)
}
