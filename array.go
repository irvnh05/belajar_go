package main

import "fmt"

func main(){
// membuat array secara manual
	var names[3]string
	names[0] = "irvan"
	names[1] = "hilmi"
	names[2] = "dermawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

// membuat array + datanya lagnsung
	var values = [3]int{ //panjang array 3
		90,
		20,
		23,
	}
	fmt.Println(values)

// function array 
// len(array) - untuk mendapatkan panjang array
// array[index] - mendapatkan data di posisi index
// array[index]= value - mengubah data di posisi index

	fmt.Println(len(names)) //mengambil panjang data bukan jumlah data
	fmt.Println(len(values))
	
}