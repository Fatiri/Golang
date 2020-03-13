package main

import (
	"fmt"
	"math"
	"math/rand"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	//Math.SQRT2: Berisi hasil akar kuadrat dari 2, dengan nilai 1.4142135623730951
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	for i := 0; i < 20; i++ {
		//rand.intn untuk mngenerate nilai integer
		fmt.Println(rand.Intn(20))
	}
}

// Method lanjutan

// Anda bisa mendeklarasikan method pada tipe selain struct juga.

// Pada contoh berikut kita dapat melihat sebuah tipe numerik MyFloat dengan method Abs.

// Anda hanya bisa mendeklarasikan sebuah method dengan sebuah receiver yang tipenya didefinisikan di paket yang sama dengan method-nya. Anda tidak bisa mendeklarasikan sebuah method dengan receiver yang tipenya didefinisikan di paket yang lain (termasuk tipe dasar seperti int).
