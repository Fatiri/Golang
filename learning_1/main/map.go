package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main()  {
	m = make(map[string]Vertex)
	m["Ilham Fatiri"] = Vertex{
		40.783467, -63.329332,
	}

	fmt.Println(m["Ilham Fatiri"])
}