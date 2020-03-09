package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Print("Go berjalan pada ")
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux")
	// case "OSBD":
	// 	fmt.Println("OS BD")
	// default:
	// 	fmt.Printf("%s. \n", os)
	// }

	fmt.Println("Kapan hari sabtu ?")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Sekarang")
	case today + 1:
		fmt.Println("Besok")
	case today + 3:
		fmt.Println("Dua hari lagi")
	default:
		fmt.Println("Masih jauh")
	}
}
