package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	Name             FullName
	DateOfBirth      BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Dylan",
			LastName:  "Bellomusto",
		},
		DateOfBirth: BirthDate{
			Day:   27,
			Month: 5,
			Year:  2008,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       '\u264A',
	}

	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
