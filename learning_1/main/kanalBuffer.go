package main

import "fmt"

func main() {
	ch := make(chan int, 100)
	for i := 0; i < 5; i++ {
		ch <- 1 + 1
		fmt.Println(<-ch)
	}

}

// Kanal dengan buffer

// Kanal memiliki buffer. Cukup dengan menambahkan panjang buffer ke argumen kedua pada make untuk menginisialisasi buffer kanal:

// ch := make(chan int, 100)

// Pengiriman ke kanal buffer akan ditahan bila buffer telah penuh. Penerimaan akan ditahan saat buffer kosong.

// Ubahlah contoh untuk memenuhi buffer dan lihat apa yang terjadi.
