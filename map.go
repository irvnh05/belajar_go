// tipe data lain yang berisikan kumpulan data yang sama,namun kita bisa memnentukan jenis tipe data index yang akan digunakan

package main

import "fmt"

func main(){

	person := map[string]string{
		"name":"irvan",
		"address":"bandung",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"]) //mengambil data di map dengan key
	fmt.Println(person["address"])

	var book map[string]string =make(map[string]string) 
	book["title"] = "belajar go-lag"
	book["author"] = "irvan"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book,"ups")

	fmt.Println(book)
	fmt.Println(len(book))

}