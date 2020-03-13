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

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(v.Abs())
}

// Di sini kita lihat method Abs dan Scale dibuat ulang sebagai fungsi.

// Sekali lagi, coba hilangkan * pada baris 16. Bisakah anda melihat perubahan perilakunya? Apa yang harus anda ubah selanjutnya supaya contoh tersebut dapat di kompilasi?

// (Jika anda tidak yakin, lanjutkan ke halaman berikutnya.)
