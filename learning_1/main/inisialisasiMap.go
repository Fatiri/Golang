package main

import "fmt"

type Vertex struct {
	string, int string
}

var m = map[string]Vertex{
	"ilham fatiri": Vertex{
		"jaka", "joko",
	},
	"Google": Vertex{
		"juku", "jiki",
	},
}

var n = map[string]int{"jaka": 2, "joko": 3}

func main() {
	fmt.Println(m)
	fmt.Println(n)
}
