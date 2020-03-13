package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

// Method dan pointer tidak langsung

// Bandingkan dua program sebelumnya, anda mungkin memperhatikan bahwa fungsi dengan sebuah argumen pointer harus menerima sebuah pointer:

// var V Vertex
// ScaleFunc(v, 5)  // Compile error!
// ScaleFunc(&v, 5) // OK

// Sementara method dengan pointer-receiver dapat menerima sebuah nilai atau sebuah pointer sebagai receiver saat dipanggil:

// var v Vertex
// v.Scale(5)   // OK
// p := &v
// p.Scale(10)  // OK

// Untuk perintah v.Scale(5), walaupun v adalah sebuah nilai bukan sebuah pointer, method dengan pointer-receiver akan dipanggil secara otomatis. Maka, untuk mempermudah, Go menginterpretasikan perintah v.Scale(5) sebagai (&v).Scale(5) karena method Scale memiliki pointer-receiver.
