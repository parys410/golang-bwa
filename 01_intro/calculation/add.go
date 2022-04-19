package calculation

func Add(number int, numberTwo int) int {
	/*
		Add huruf A besar artinya bersifat public untuk bisa dipanggil di package lainnya
		apabila menggunakan huruf kecil, maka hanya bisa dipanggil di local package tersebut
	*/
	return add(number, numberTwo)
}

func add(number int, numberTwo int) int {
	return number + numberTwo
}