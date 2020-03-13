package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

// Goroutines

// Sebuah goroutine yaitu suatu thread ringan yang diatur oleh Go runtime. Perintah

// go f(x, y, z)

// memulai sebuah goroutine baru f dengan menjalankan

// f(x, y, z)

// Evaluasi dari nilai f, x, y, dan z terjadi di goroutine yang memanggil dan eksekusi dari f terjadi di goroutine yang baru.

// Goroutine berjalan di ruang alamat yang sama, sehingga akses ke shared memory harus disinkronisasi. Paket sync menyediakan fungsi primitif untuk sinkronisasi, walaupun anda tidak terlalu membutuhkannya karena ada fungsi primitif lainnya. (Lihat slide selanjutnya.)
