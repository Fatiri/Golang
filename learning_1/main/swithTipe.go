package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Dua kali %v adalah %v\n", v, v*2)
	case string:
		fmt.Printf("%q adalah %v bytes panjangnya\n", v, len(v))
	default:
		fmt.Printf("Saya tidak kenal dengan tipe %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)

}

// Penggunaan switch untuk tipe

// Sebuah tipe switch adalah bentukan yang membolehkan beberapa penegasan tipe secara serial.

// Tipe switch sama seperti perintah switch biasa, tapi dengan nilai case mengacu pada tipe (bukan nilai), dan nilai case tersebut dibandingkan dengan tipe yang dikandung oleh isi interface yang diberikan.

// switch v := i.(type) {
// case T:
//     // di sini v bertipe T
// case S:
//     // di sini v bertipe S
// default:
//     // tidak ada yang cocok; disini v bertipe sama dengan i
// }

// Deklarasi dalam sebuah tipe switch memiliki sintaks yang sama dengan penegasan tipe i.(T), tapi dengan T diganti dengan kata type.

// Perintah switch berikut menguji apakah isi interface i mengandung sebuah nilai dari tipe T atau S. Pada setiap case dari T dan S, variabel v akan memiliki tipe T atau S dan memiliki nilai yang dikandung oleh i. Pada `default case` (yang mana tidak ada tipe yang sama ditemukan), variabel v akan memiliki tipe dan nilai yang sama dengan i.
