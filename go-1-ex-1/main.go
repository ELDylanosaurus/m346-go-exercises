package main

import "fmt"

func main() {

	var firstName string = "Dylan"
	var lastName string = "Bellomusto"
	var dayOfBirth int = 27
	var monthOfBirth int = 5
	var yearOfBirth int = 2008
	var numberOfSiblings int = 2
	var heightInMeters float64 = 1.70
	var zodiacSign rune = '\u264a'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
