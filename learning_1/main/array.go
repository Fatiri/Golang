package main

import "fmt"

func main() {
	var a [2]string

	a[0] = "Jaka"
	a[1] = "Tarub"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [6]int{1, 2, 3, 4, 5}
	b[5] = 2

	fmt.Println(b)
}
