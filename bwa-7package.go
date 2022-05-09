package main

import (
	"fmt"
	"pertama/calculation"
)


func main() {
	fmt.Println("Halo, Belajar Golang")
	// sentence := TestAja()

	// fmt.Println(sentence)

	result := calculation.Add(8,120)
	fmt.Println(result)
}

// perihal package
// merupakan standar go untuk file utama, bisa mengguakan selain main tapi menjadi libry yang tetap di akses pada package main
// agar bisa di eksekusi dengan comand line

// perihal import
// standart libary menggunakan import - fmt
// ketika ingin akses package yang berbeda maka gunakan import - calculation
// akses package yg dibuat orang lain yg di publish di git - soon

// perihal func main
// bersifat wajib, dengan penamaan tidak sembarangan karena ini merupakan aturan default dari go kecuali pada package custom/library