package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Sebelum scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("Setelah scaling: %+v, Abs: %v\n", v, v.Abs())
}

// Memilih receiver sebagai nilai atau pointer

// Ada dua alasan kenapa menggunakan pointer-receiver.

// Pertama adalah supaya method dapat mengubah nilai yang ditunjuk oleh receiver.

// Kedua adalah untuk menghindari menyalin nilai setiap kali method dipanggil. Hal ini bisa lebih efisien jika receiver adalah sebuah struct yang besar, sebagai contohnya.

// Pada contoh ini, Scale dan Abs memiliki tipe receiver *Vertex, walaupun method Abs tidak perlu mengubah receiver-nya.

// Secara umum, semua method dari sebuah tipe harus memiliki receiver sebagai nilai atau sebagai pointer, tapi tidak gabungan dari keduanya. (Kita akan lihat alasannya pada beberapa laman berikutnya.)
