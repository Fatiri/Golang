package main

import "fmt"

func main()  {
   var name = []int{1,2,3,4,5}
   number := "Number"
   pritnMessage(number, name);
   
}

func pritnMessage(number string, arr []int){
	fmt.Println(number, arr)
}