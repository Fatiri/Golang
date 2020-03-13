package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func main() {
	for i:=0; i<10     ;i++  {
		fmt.Print(rand.Intn(30))
	}
	fmt.Println("nilai menggunakan format math untuk mencetak  : ", rand.Intn(20))
	fmt.Println("math menggunkaan Sqrt / nilai akar kuadrat dari nilai 10 : ", math.Sqrt(10))
	fmt.Println("math menghasilkan nilai PI : ", math.Pi)
	fmt.Println("hasil dari penjumlahan 10 + 8 = ", count(10, 8))
	fmt.Println("hasil dari penjumlahan dengan menggabungkan tipe data pada parameter 10 + 8 = ", count1(10, 8))
	hello, world := helloWorld("hello", "world")
	fmt.Println("contoh 2 return value string : ", hello, world)

	fmt.Print("menghasilkan 2 return dari 1 parameter : ")
	fmt.Println((split(30)))

	var i, j = 1, 2
	var c, phyton, java = true, false, "no!"
	fmt.Println("variable with initializer : ", i, j, c, phyton, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//tipe data menggunakan var
	var in int
	var f float64
	var f32 float32
	var b bool
	var s string
	fmt.Println("i : ", in, ", float32 : ", f32, ", float64 : ", f, ", boolean : ", b, ", string : ", s )

	// type data conversion
	var x, y int = 3, 4
	var f64 float64 = math.Sqrt(float64(x + y))
	var z uint = uint(f64)
	fmt.Println(x, y, z)

	// type inference
	inference := 0.56 + 6i
	int_i := 27
	float := 27.15
	fmt.Printf("nilai 0.56 + 6i merupakan typr data : %T\n", inference)
	fmt.Printf("nilai 27 merupakan typr data : %T\n", int_i)
	fmt.Printf("nilai 27.0 merupakan typr data : %T\n", float)

	// constanta
	const name  = "SUSI"
	fmt.Println("nama saya : ",name)

	//perulangan
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))
}

//fungsi penjumlahan menggunakan format penulisan function bisasa
func count(x int, y int) int {
	return x + y
}

// fungsi penjumlahan dengan menggabungan tipe data yang sama pada paramater
func count1(x, y int) int {
	return x + y
}

//fungsi untuk menampilkan string dengan 2 return
func helloWorld(x, y string) (string, string) {
	return x, y
}

// fungsi yang menghasilkan return 2 dengan 1 parameter
func split(sum int) (x, y int) {
	x = sum * 5
	y = sum - 10
	return
}

//function tipe data basic
var (
	ToBe   bool = false
	MaxInt uint64 = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}





