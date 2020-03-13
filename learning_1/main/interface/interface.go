package interface

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat mengimplementasikan Abser
	a = &v // a *Vertex mengimplementasikan Abser

	// Pada baris berikut, v adalah sebuah Vertex (bukan *Vertex)
	// dan TIDAK mengimplementasikan Abser.
	a = &v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Interface

// Sebuah type interface didefinisikan sebagai sebuah kumpulan method penanda.

// Nilai dari tipe interface dapat menyimpan nilai apapun yang mengimplementasikan method tersebut.

// Catatan: Terdapat kesalahan di contoh kode pada baris 22. Vertex (tipe nilai) tidak memenuhi Abser karena method Abs hanya didefinisikan pada *Vertex (tipe pointer).
