package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	const name = "ilham fatiri"
	fmt.Println("Hello", name)
	fmt.Println("Happy", Pi, "Day")

	const truth = false
	fmt.Println("Go rules ? ", truth)
}

// Constants are declared like variables, but with the const keyword.

// Constants can be character, string, boolean, or numeric values.

// Constants cannot be declared using the := syntax.
