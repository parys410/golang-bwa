package main

import (
	"bwa-go/01_intro/calculation" // "namaModule/package"
	"fmt"
)

func main(){
	fmt.Println("Halo, belajar Golang")

	result := calculation.Add(8, 10)

	fmt.Println("Result", result);

	// sentence := TestAja()

	// fmt.Println(sentence)
}