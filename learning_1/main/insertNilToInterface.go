package main

import "fmt"

type User interface {
	M()
}

func main() {
	var u User
	describe(u)
	u.M()
}

func describe(u User) {
	fmt.Printf("(%v, %T)\n", u, u)
}

// Isi interface dengan nil

// Sebuah isi interface yang nil tidak mengandung nilai dan tipe konkrit.

// Memanggil sebuah method pada sebuah interface yang nil akan mengakibatkan kesalahan run-time karena tidak ada tipe di dalam interface yang mengindikasikan method konkrit yang akan dipanggil.
