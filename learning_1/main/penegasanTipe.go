package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

// Penegasan tipe

// Suatu penegasan tipe menyediakan akses ke isi interface di balik nilai konkritnya.

// t := i.(T)

// Perintah di atas menegaskan bahwa isi interface i menyimpan tipe konkrit T dan memberikan nilai T ke variabel t.

// Jika i tidak mengandung tipe T, perintah tersebut akan memicu panic.

// Untuk memeriksa apakah sebuah isi interface benar mengandung tipe tertentu, penegasan tipe bisa mengembalikan dua nilai: nilai yang dikandung dan sebuah nilai boolean yang memberitahu apakah penegasan sukses atau tidak.

// t, ok := i.(T)

// Jika i mengandung T, maka t akan menyimpan nilai dan ok akan bernilai true.

// Jika tidak, ok akan bernilai false dan t akan bernilai kosong sesuai dengan tipe T, dan panic tidak akan terjadi.

// Ingatlah kesamaan antara sintaks ini dengan membaca dari sebuah map.
