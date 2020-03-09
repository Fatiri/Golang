package main

import "fmt"

func main() {
	// defer fmt.Println("Terakhir")
	// fmt.Println("Awal")
	//penundaan bertumpuk

	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
