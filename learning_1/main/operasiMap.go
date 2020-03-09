package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Age"] = 24
	fmt.Println("The Value is = ", m["Ilham Fatiri"])

	m["Age"] = 24
	fmt.Println("The value is = ", m["Age"])

	delete(m, "Age")
	fmt.Println("The value is = ", m["Age"])

	m["Age"] = 24
	v, ok := m["Age"]
	fmt.Println("The value is = ", v, "Present ? ", ok)
}
