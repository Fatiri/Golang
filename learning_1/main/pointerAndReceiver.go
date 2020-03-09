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

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// Pointer-receiver

// Anda bisa mendeklarasikan method dengan receiver berupa pointer.

// Hal ini berarti tipe receiver memiliki sintaks *T untuk tipe T. (Dan juga, T itu sendiri tidak bisa berupa pointer ke tipe dasar seperti *int.)

// Sebagai contohnya, method Scale didefinisikan pada *Vertex.

// Method dengan pointer-receiver dapat mengubah nilai yang ditunjuk (seperti yang dilakukan oleh Scale). Karena method terkadang perlu mengubah receiver nya, pointer-receiver lebih umum digunakan daripada receiver dengan value.

// Coba hapus * dari deklarasi fungsi Scale pada baris 16 dan perhatikan bagaimana perilaku program berubah.

// Dengan receiver-sebagai-value, method Scale beroperasi pada salinan dari nilai asli Vertex. (Perilaku ini juga berlaku pada argumen pada fungsi.) Method Scale harus memiliki sebuah pointer-receiver untuk dapat mengubah nilai Vertex yang dideklarasikan pada fungsi main.
