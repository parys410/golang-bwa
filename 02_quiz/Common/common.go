package Common

import "strings"

func EvenChecker(index int) bool {
	if index%2 == 0 {
		return true
	} else {
		return false
	}
}

func LetterChecker(huruf string) bool {
	switch strings.ToLower(huruf) {
		case "a":
			return true
		case "i":
			return true
		case "u":
			return true
		case "e":
			return true
		case "o":
			return true
		default:
			return false
	}
}