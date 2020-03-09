package main

import "fmt"

func main() {
	number := [6]int{1, 2, 3, 4, 5, 6}

	var number2 []int = number[1:4]

	fmt.Println(number2)

	names := [4]string{
		"jaka",
		"joko",
		"jiki",
		"juku",
	}

	fmt.Println(names)

	a := names[0:4]
	b := names[1:2]
	c := names[:2]
	d := names[0:]
	fmt.Println(a, b, c, d)

	b[0] = "ngok"
	fmt.Println(a, b)
	fmt.Println(names)

}
