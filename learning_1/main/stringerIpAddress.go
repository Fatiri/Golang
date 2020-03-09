package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v.%v.%v.%v\n", name, ip[0], ip[1], ip[2], ip[3])
	}
}

// Latihan: Stringer

// Buat tipe IPAddr yang mengimplementasikan fmt.Stringer untuk mencetak alamat dengan empat tanda titik.

// Misalnya, IPAddr{1, 2, 3, 4} mengeluarkan "1.2.3.4".
