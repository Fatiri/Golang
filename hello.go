package main

import "fmt"

func main(){

	// %s - %d - %t = String Formater


	//Pemberian type data sesuai valuenya
	name := "jaka"; 
	fmt.Println(name);
	fmt.Println("----------------------------------");
	
	//Type data numeric decimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	fmt.Println("----------------------------------");

	//Operator Aritmatika
	var value = (((2+6)%3)* 4-2) / 3

	fmt.Println(value);
	fmt.Println("----------------------------------");

	//Operator Aritmatica and operator Perbandingan
	var value1 = (((2+6)%3)* 4-2) / 3
	var isEqual =  (value1 == 2)

	fmt.Printf("nilai %d (%t) \n", value1, isEqual);
	fmt.Println("----------------------------------");
	
	//Operator Logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left & right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left or right \t(%t) \n", leftOrRight)
	var leftRight = !left
	fmt.Printf("left \t\t(%t) \n", leftRight)
	fmt.Println("----------------------------------");

	//Seleksi Kondisi
	var point = 5

	if point == 10{
		fmt.Println("Perfect")
	}else if point > 5 {
		fmt.Println("Lulus")
	}else if point == 4 {
		fmt.Println("Hampir Lulus")
	}else {
		fmt.Println("Tidak Lulus")
	}
	fmt.Println("----------------------------------");

	//Perulangan
	for i := 1; i < 5; i++ {
		if i == 4{
			break
		}
		fmt.Println("angka", i);
	}

	fmt.Println("----------------------------------");

	//Array
	var fruits =  [...]string{"jaka","joko","juku","jeke", "juku"}
	fmt.Println("Jumlah Element \t\t" , len(fruits))
	fmt.Println("Isi Semua Element \t", fruits)
	fmt.Println("----------------------------------");
	
	//Array Mutidimensi
	var fruits2 = [2][3]int{[3]int{1,2,3},[3]int{1,2,3}}
	fmt.Println("Array Multi dimensi", fruits2)
	fmt.Println("----------------------------------");

	//Perulangan Array menggunakan for
	var fruits3 = [...]string{"mangga","jeruk","melon","anggur"}
	for i :=0; i<len(fruits3); i++{
		fmt.Printf("Element index %d : %s\n", i, fruits3[i])
	}
	fmt.Println("----------------------------------");

	//Slice Pada Array
	var fruits4 = [...]string{"mangga","jeruk","melon","anggur"}
	slice1 := fruits4[0:2]

	fmt.Println("Setelah di Slace", slice1)
	fmt.Println("----------------------------------");

	//MAP - key and value
	var moli = map[string]int{"January":1, "February":2,"Maret": 3}
	fmt.Println("Array Map :", moli)
	fmt.Println("----------------------------------");


	//MAP - Perulangan 
	var moli2 = map[string]int{"January": 1, "February": 2, "Marte":3}
	for key, value := range moli2{
		fmt.Println(key, value)
	}
	fmt.Println("----------------------------------");
	
	//
}