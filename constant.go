package main

import "fmt"

func main(){
	// penulisan versi 1
	// const firstName string = "irvan"
	// const lastName = "hilmi"
	// const value = 1000

	// penulisan versi 2
	const( 
		firstName string 	= "IRVAN"
		lastName 			= "HILMI"
		value 				= 1
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
	// error
	// firstName = "ini irvan"
}
