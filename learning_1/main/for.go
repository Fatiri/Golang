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

	length := 4

	for i := 1; i < length-1; i-- {
		for j := i; j < length; j++ {
			fmt.Print("*")
		}

		fmt.Println("")
	}

}
