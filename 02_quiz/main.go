package main

import (
	"bwa-go/02_quiz/Common"
	"fmt"
)

func main() {
	title := "Golang is the best language"

	fmt.Println(title)

	for index,karakter := range title {
		if Common.EvenChecker(index) {
			if Common.LetterChecker(string(karakter)) {
				fmt.Print("*")
			}else{
				fmt.Print(string(karakter))
			}
		}else{
			if Common.LetterChecker(string(karakter)) {
				fmt.Print("_")
			}else{
				fmt.Print(string(karakter))
			}
		}
	}
}