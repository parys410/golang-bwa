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
		fmt.Printf("%s", language)
	}
	fmt.Printf("\n\n")


	// ********************* //
	// **      Slice		  ** //
	// ********************* //
	// Slice adalah Array Dinamis, ketika kita tidak tahu berapa jumlah data dalam Array tersebut.

	var gamingConsoles []string
	// gamingConsoles[0] = "PS4" // Error, karena pada slice jumlah array nya dinamis dan saat ini slice belum terdapat data
	gamingConsoles = append(gamingConsoles, "Test Console")
	gamingConsoles = append(gamingConsoles, "Nintento Switch")
	gamingConsoles = append(gamingConsoles, "XBOX Series X")
	gamingConsoles = append(gamingConsoles, "Steamdeck")
	gamingConsoles[0] = "Playstation 5" // Tidak Error, karena saat ini slice sudah berisi data

	gameList := []string{ "Dynasty Warior 4", "Pokemon", "Halo", "Dota 2" }

	fmt.Println(gamingConsoles, len(gamingConsoles))
	fmt.Println(gameList, len(gameList))

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}


	// ********************* //
	// **       MAP 		  ** //
	// ********************* //
	// Map adalah Key - Value variable
	var myMap = make(map[string]int)
	mySecondMap := map[string]string{"Ruby":"1", "Go":"2", "JavaScript":"3"}

	myMap["Playstation"] = 1
	myMap["Nintendo"] = 2
	myMap["XBOX"] = 3

	fmt.Println(myMap)
	fmt.Println(myMap["Playstation"])
	fmt.Println(mySecondMap)

	for key,val := range myMap {
		fmt.Printf("%s = %d \n", key, val)
	}

	// Delete Map based on the key
	delete(myMap, "XBOX")
	for key,val := range myMap {
		fmt.Printf("%s = %d \n", key, val)
	}

	value, isAvailable := myMap["Steamdeck"] // Assign map seperti ini, mengirimkan 2 value yaitu 1 = Main Value; 2 = Boolean data which is available/not
	fmt.Println(value, isAvailable)

	// Slice of Maps
	students := []map[string]string{
		{"name": "Ary","score": "A"},
		{"name": "Dwija","score": "A"},
		{"name": "Bangkit","score": "B"},
	}

	fmt.Println(students)
	for _,student := range students {
		fmt.Printf("The score of %s is %s \n", student["name"], student["score"])
	}
}