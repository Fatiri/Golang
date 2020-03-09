package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter aman digunakan secara konkuren.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc meningkatkan nilai dari key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock sehingga hanya satu goroutine pada satu waktu yang dapat
	// mengakses map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value mengembalikan nilai dari key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock sehingga hanya satu gorouting pada satu waktu yang dapat
	// mengakses map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// sync.Mutex

// Kita telah melihat bagaimana channel sangat bagus untuk komunikasi antara beberapa goroutine.

// Namun bagaimana jika kita tidak ingin komunikasi? Bagaimana jika kita hanya ingin memastikan hanya satu goroutine yang dapat mengakses suatu variabel pada satu waktu untuk menghindari konflik?

// Konsep ini disebut juga dengan mutual exclusion, dan nama umum untuk struktur data yang menyediakan fungsi tersebut adalah mutex.

// Pustaka standar Go menyediakan mutual exclusion dengan sync.Mutex dan dua fungsinya:

//     Lock
//     Unlock

// Kita bisa mendefinisikan sebuah blok kode untuk dieksekusi secara mutual exclusion dengan mengungkungnya dengan pemanggilan fungsi Lock dan Unlock seperti yang terlihat pada method Inc.

// Kita juga bisa menggunakan defer untuk memastikan mutex akan dibuka kembali seperti pada method Value.
