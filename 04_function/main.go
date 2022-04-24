package main

import "fmt"

func main() {
	fmt.Println(printMyResult("Saya sedang"))
	fmt.Println("10 + 10 = ", add(10,10))

	sum, avrg := getResult(10,15)
	_, newAvrg := getResult(10,7)
	fmt.Println(sum, avrg, newAvrg)
}

func printMyResult(pesan string) string{
	return pesan + " belajar Golang"
}

func add(number1, number2 int) (sum int) {
	sum = number1 + number2
	return
}

// Multiple Value Return
func getResult(number1, number2 int) (sum int, avrg float32) {
	sum = number1 + number2
	avrg = (float32(number1) + float32(number2)) / 2
	return
}