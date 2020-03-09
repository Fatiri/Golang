package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// Method lanjutan

// Anda bisa mendeklarasikan method pada tipe selain struct juga.

// Pada contoh berikut kita dapat melihat sebuah tipe numerik MyFloat dengan method Abs.

// Anda hanya bisa mendeklarasikan sebuah method dengan sebuah receiver yang tipenya didefinisikan di paket yang sama dengan method-nya. Anda tidak bisa mendeklarasikan sebuah method dengan receiver yang tipenya didefinisikan di paket yang lain (termasuk tipe dasar seperti int).
