package main

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	UnLock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct {
}

func (b bike) Move() {
	fmt.Println("Moving the bike")
}

func (b bike) Lock() {
	fmt.Println("Locking the bike")
}

func (b bike) UnLock() {
	fmt.Println("Unlocking the bike")
}

func main() {
	var ml MoveLocker
	var m Mover
	ml = bike{}
	m = ml
	//ml = m

	// 使用类型断言，判断m是否有自行车
	b := m.(bike)
	ml = b
}
