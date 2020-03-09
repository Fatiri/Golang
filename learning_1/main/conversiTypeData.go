package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x, y = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	//Convert to string from int
	var convertToString string = strconv.Itoa(x)

	fmt.Println(x, y, z, convertToString)
}
