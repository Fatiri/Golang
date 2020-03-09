package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

// Method dan pointer tidak langsung (2)

// Hal yang sama berlaku sebaliknya.

// Fungsi yang menerima argumen nilai harus mendapatkan sebuah nilai dari tipe yang spesifik:

// var v Vertex
// fmt.Println(AbsFunc(v))  // OK
// fmt.Println(AbsFunc(&v)) // Compile error!

// Sementara method dengan value-receiver bisa menerima sebuah nilai atau sebuah pointer sebagai receiver saat dipanggil:

// var v Vertex
// fmt.Println(v.Abs()) // OK
// p := &v
// fmt.Println(p.Abs()) // OK

// Pada kasus ini, pemanggilan method p.Abs() diinterpretasikan sebagai (*p).Abs().
