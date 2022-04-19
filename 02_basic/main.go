package main

import "fmt"

func main() {
	// Variable
	// var name string = "Golang"
	// var name2 = "Ruby on Rails"
	// age := 20
	// age = 29
	// fmt.Println(name, name2, age)

	
	// Conditional Variable
	age := 14

	if (age > 20){
		fmt.Println("Bekerja dulu")
	} else if (age > 10){
		fmt.Println("Boleh bermain game")
	} else {
		fmt.Println("Belajar dulu sana")
	}

	score := 80
	var grade string

	// if (score <= 50){
	// 	grade = "E"
	// }else if (score <= 60){
	// 	grade = "D"
	// }else if (score <= 70){
	// 	grade = "C"
	// }else if (score <= 90){
	// 	grade = "B"
	// }else {
	// 	grade = "A"
	// }

	switch (score) {
		case 90:
			grade = "A"
		case 80:
			grade = "B"
		case 70:
			grade = "C"
		case 60:
			grade = "D"
		default :
			grade = "E"
	}

	fmt.Println(grade)

	// Foor Loop
	for i:=1; i <= 10; i++{
		fmt.Println(i)
	}

	// While Loops in Go
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	// For each in Go
	title := "Golang is the best language"

	for _, letter := range title {
		// fmt.Println(index, string(letter))
		fmt.Println(string(letter))
	}

}