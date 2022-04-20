package main

import "fmt"

func main() {
	// ********************* //
	// **      Array		  ** //
	// ********************* //
	var arrString [5]string
	var arrInt [5]int
	var arrBool [5]bool

	fmt.Println(arrString, arrInt, arrBool)

	var languages [5]string
	languages[0] = "Go"
	languages[1] = "Ruby"
	languages[2] = "JavaScript"
	languages[3] = "C#"
	languages[4] = "Python"

	// Shorten Method
	var newLanguages = [5]string { "Go", "Ruby", "JavaScript", "C#", "Python" }
	
	// Mau buat array tapi males hitung jumlahnya gunakan [...]
	carType := [...]string {
		"Toyota",
		"Daihatsu",
		"Suzuki",
		"Honda",
	}

	fmt.Println(languages, newLanguages, carType)

	// Length of Array
	fmt.Println(len(languages), len(carType))

	// For range of arrays
	for _,language := range languages {
		fmt.Printf("%s ", language)
	}

}