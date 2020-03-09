package main

import "fmt"

func fibonacci() func() int {
	n, t1, t2, sum := 10, 0, 1, 0
	return func() int {
		for i := 1; i < n; i++ {
			sum = t1 + t2
			t1 = t2
			t2 = sum
			return t1
		}
		return 0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
