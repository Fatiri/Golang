package main

import (
	"fmt"
	"math"
)

func main() {

	var items, lang float64
	items = 464740
	lang = math.Round(0.01 * items)
	fmt.Println(lang)
}
