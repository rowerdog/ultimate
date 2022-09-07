package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < 6; i++ {
		fmt.Println(arr[i])
	}
}
