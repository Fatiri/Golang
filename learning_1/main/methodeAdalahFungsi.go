package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

// Method adalah fungsi

// Ingat: sebuah method hanyalah fungsi dengan argumen sebuah receiver.

// Berikut ini Abs ditulis sebagai fungsi biasa tanpa ada perubahan pada fungsionalitas.
