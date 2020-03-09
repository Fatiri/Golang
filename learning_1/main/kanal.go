package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

// Kanal

// Kanal adalah sebuah penghubung tipe, yang mana anda bisa mengirim dan menerima nilai dengan operator kanal, <-.

// ch <- v    // Kirim v ke kanal ch.
// v := <-ch  // Terima dari ch, dan simpan nilainya ke v.

// (Aliran data sesuai dengan arah panah.)

// Seperti map dan slice, kanal harus dibuat sebelum digunakan:

// ch := make(chan int)

// Secara bawaan, pengiriman dan penerimaan ditahan sampai sisi yang lain siap. Hal ini membolehkan goroutine untuk melakukan sinkronisasi tanpa melakukan penguncian secara eksplisit atau menggunakan variabel kondisi.

// Contoh kode menjumlahkan angka yang ada di slice, dengan mendistribusikan kerja antara dua goroutine. Saat kedua goroutine selesai, hasil akhirnya kemudian akan dihitung.
