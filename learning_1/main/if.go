package main

import "fmt"

func main() {
	count := 10

	if count == 2 {
		fmt.Println("Nice")
	} else if count > 2 {
		fmt.Println("Very Nice")
	} else {
		fmt.Println("Try Again")
	}
}
