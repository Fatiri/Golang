package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"ilham fatiri": Vertex{
		40.927383, -83.8932683,
	},
	"Google": Vertex{
		50.838743, -70.89800,
	},
}

func main() {
	fmt.Println(m["Google"])
}
