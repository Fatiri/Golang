package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//methode
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// Go tidak memiliki class. Namun, anda bisa mendefinisikan method pada tipe.

// Sebuah method adalah sebuah fungsi dengan argumen khusus receiver.

// receiver muncul pada bagian antara kata kunci func and nama method.

// Pada contoh berikut, method Abs memiliki receiver dengan tipe Vertex bernama v.
