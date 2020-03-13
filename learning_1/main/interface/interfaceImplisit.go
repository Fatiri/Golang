package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Method berikut berarti type T mengimplementasikan interface I,
// tapi kita tidak perlu secara eksplisit mendeklarasikannya.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

// Interface dipenuhi secara implisit

// Sebuah tipe mengimplementasikan sebuah interface dengan mengimplementasikan method-methodnya. Tidak ada deklarasi eksplisit; tidak ada perintah "implements".

// Interface implisit memisahkan definisi dari sebuah interface dari implementasinya, yang bisa saja muncul dalam paket manapun tanpa adanya penataan sebelumnya.
