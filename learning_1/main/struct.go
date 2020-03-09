package main

import "fmt"

type Name struct {
	Name string
	Age  int
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{2, 3}
	v2 = Vertex{X: 4}
	v3 = Vertex{}
	z  = &Vertex{2, 3}
)

func main() {
	fmt.Println(Name{"Ilham Fatiri", 24})
	v := Name{"Jaka Tarub", 30}
	v.Name = "Joko"
	v.Age = 40

	fmt.Println(v.Name, v.Age)

	q := Name{"Maman Abdruahman", 50}
	p := &q
	p.Age = 1e9
	fmt.Println(q)

	fmt.Println(v1, v2, v3, *z)

}
