package main

import "fmt"

func main() {
	count := 20
	sum := 0
	for i := 0; i < count; i++ {
		sum += 1
	}

	fmt.Println(sum)

	names := [...]int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(names); i++ {
		fmt.Printf("Array Loop  %d  = %d\n ", i, names[i])
	}

	sum2 := 1
	//Jika for hanya di beri kondisi saja maka bisa di bilang while
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//Menjadi stack overflow
	// for {

	// }
	fmt.Println("")
	length := 8

	for i := 0; i < 4; i++ {
		for i := 0; i < length; i++ {
			for j := length - i; j > 1; j-- {
				fmt.Print(" ")
			}

			for j := 0; j <= i; j++ {
				fmt.Print("* ")
			}

			fmt.Println("")
		}
	}

}
