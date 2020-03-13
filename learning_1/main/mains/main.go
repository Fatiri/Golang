package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	a := 10
	fmt.Println("nilai awal : ", a)
	tambah1(a)
	fmt.Println("pass by value ", a)

	tambah(&a)
	fmt.Println("pass by refrence ", a)

	convertInt("a")

	// Basic
	looping()
}

func tambah(a *int) {
	*a += 10
}
func tambah1(a int) {
	a += 10
}

// IF TANPA RETURN maka ketika error akan tetap melanjutkan ke code selanjutnya

func convertInt(huruf string) {
	angkaString, err := strconv.Atoi(huruf)
	if err != nil {
		println("INI ERROR")
	}
	println(angkaString)
	println("NEXT")
}

type User struct {
	id   int
	name string
}

func (u User) findUser() {

}

func findUserByID() {
	//var user User
}

func looping() {
	for i := 0; i < 10; i++ {
		fmt.Println("perulangan ", i)
	}

	array := [2]string{"susi", "cahyati"}
	fmt.Println("Array biasa ", array)

	a := []int{}
	fmt.Println(a)
	a = append(a, 10, 12, 11)
	fmt.Println("slice : ", a)

	for i := 0; i < 5; i++ {
		a = append(a, i)
	}
	fmt.Println("input slice menggunakan for : ", a)

	var b = a
	fmt.Println("mengisi variable baru dengan array : ", b)

	b = append(b, 150)
	fmt.Println("INI A", a)
	fmt.Println("INI B", b)

	c := [2]int{0, 1}
	fmt.Println("INI C", c)

	d := c
	fmt.Println("INI D", d)
	d[0] = 10
	fmt.Println(c)

	var str string = "INI STRING"
	fmt.Println(str)
	var iniByte []byte

	iniByte = []byte(str)
	fmt.Println(iniByte)

	var aFloat float32 = 1.0
	var bInt int = 10
	var cString string = "STRING ---- string"

	fmt.Println(aFloat + float32(bInt))
	fmt.Println(strings.Contains(cString, "s"))

	test := strings.Replace(cString, "-", "+", 3)
	fmt.Println(test)

	var statement string = "INI STRING"
	arrStatement := strings.Split(statement, "")
	fmt.Println(arrStatement)

	arrStatement1 := strings.Split(statement, " ")
	fmt.Println(arrStatement1)

	arrStatementJoin := strings.Join(arrStatement, "+")
	fmt.Println(arrStatementJoin)

}
