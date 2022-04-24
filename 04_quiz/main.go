package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	fmt.Printf("Jumlah dari %v adalah %d\n", scores, sum(scores))

	result, err := calculate(10, 2, "+")
	fmt.Println(result, err)
	result, err = calculate(10, 2, "-")
	fmt.Println(result, err)
	result, err = calculate(10, 2, "*")
	fmt.Println(result, err)
	result, err = calculate(10, 2, "/")
	fmt.Println(result, err)
	result, err = calculate(10, 2, "=")
	fmt.Println(result, err)

	if err != nil {
		fmt.Println("Terdapat Error")
	}
}

func sum(scores []int) int {
	jumlah := 0
	for _,value := range scores {
		jumlah += value
	}
	return jumlah
}

func calculate(number1, number2 int, method string) (hasil float32, err error) {
	switch method {
	case "+":
		hasil = float32(number1) + float32(number2)
	case "-":
		hasil = float32(number1) - float32(number2)
	case "*":
		hasil = float32(number1) * float32(number2)
	case "/":
		hasil = float32(number1) / float32(number2)
	default:
		err = errors.New("undefined method")
	}
	return
}