package calculation

// package bersifat publik atau bisa di akses diluar
func Add(number int, numberTwo int) int {
	return number + numberTwo
}

// package bersifat private atau hanya bisa di panggil pada package yang sama
func add(number int, numberTwo int) int{
	return number + numberTwo
}