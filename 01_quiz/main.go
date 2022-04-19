package main

import (
	"bwa-go/01_quiz/calculation"
	"fmt"
)

func main(){
	fmt.Println("Program Perkalian")
	
	var variable1 int
	var variable2 int
	var hasil int

	fmt.Print("Input Angka Pertama : ")
	fmt.Scanln(&variable1)

	fmt.Print("Input Angka Kedua : ")
	fmt.Scanln(&variable2)

	hasil = calculation.Multiplication(variable1, variable2)

	fmt.Printf("%d * %d = %d",variable1, variable2, hasil)
}