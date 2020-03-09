package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Seleksi "default"

// Kondisi default pada select dijalankan jika tidak ada kondisi yang siap.

// Gunakan kondisi default untuk mencoba mengirim dan menerima tanpa diblok:

// select {
// case i := <-c:
//     // use i
// default:
//     // receiving from c would block
// }
