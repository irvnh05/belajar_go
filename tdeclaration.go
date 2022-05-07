package main

import "fmt"

func main(){
	type NoKTP string
	type Married bool

	var ktpIrvan NoKTP = "1321312412312"
	var marriedStatus Married = true
	fmt.Println(ktpIrvan)
	fmt.Println(marriedStatus)
	
}