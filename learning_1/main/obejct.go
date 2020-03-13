package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	v := User{"Ilham", 30}
	v.Name = "Jaka Fatiri"
	v.Age = 24

	fmt.Println(v)
}
