package main 

import "fmt"

func main(){
	var months = [...]string{ //kalau belum tau jumlah panjang datanya maka kasih ... kalau sudah tau bisa langsung 
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"july",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "diubah" //merubah slice
	// fmt.Println(slice1)

	// slice1[0] = "mei lagi" //index 0 pada slice 4 yaitu mei dirubah jadi mei lagi
	// fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "irvan") //
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// membuat slice dari awal banget = gunakane fungsi make slice

	newSlice := make([]string, 2, 5)

	newSlice[0] = "irvan"
	newSlice[1] = "hilmi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1,2,3,5,6}
	iniArrayy := [...]int{3,2,3,1,3}
	inislice := []int{3,2,1,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniArrayy)
	fmt.Println(inislice)
}