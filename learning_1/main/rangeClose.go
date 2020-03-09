package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// "range" dan "close"

// Pengirim dapat menutup (close) sebuah kanal untuk menandakan bahwa tidak ada lagi data yang dikirim. Penerima dapat memeriksa apakah kanal telah ditutup dengan menambahkan parameter kedua pada ekspresi penerimaan:

// v, ok := <-ch

// ok bernilai false jika tidak ada lagi nilai yang diterima dan kanal telah ditutup.

// Pengulangan for i := range c menerima nilai dari kanal berulang kali sampai ditutup.

// Catatan: Hanya pengirim yang sebaiknya menutup kanal, bukan si penerima. Mengirim ke kanal yang telah tertutup akan menyebabkan panic.

// Catatan lain: Kanal tidak seperti file; yang mana anda tidak selalu perlu menutupnya. Penutupan hanya diperlukan saat penerima harus diberitahu bahwa tidak ada lagi nilai yang akan diterima, misalnya untuk menghentikan pengulangan pada range.
