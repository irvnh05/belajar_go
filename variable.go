package main

import "fmt"

func main(){

	// deklrasi tipe data dengan var
    var name string // tipe data hanya bisa satu
    
	name = "irvan hilmi"
	fmt.Println(name)

	name = "dermawan"
	fmt.Println(name)	

	// tanpa deklrasi tipe data
	var friendName = "hilmi" // contoh tidak perlu deklrasi tipe data (string)
	fmt.Println(friendName)

	var age = 23 // contoh tidak perlu deklrasi tipe data (int)
	// age = 30 contoh keliru karena ini perlu di deklrasikan
	fmt.Println(age)

	// tidak wajib penggunaan kata kunci var
	names := "IRVAN HILMI" //tidak wajib menggunakan kata kunci var, diganti jadi := dan nameselanjutnya = 
	fmt.Println(names)

	names = "HILMI"
	fmt.Println(names)

	// deklrasi multiple variable
	var (
		firstName = "irvan"
		lastName = "hilmi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}