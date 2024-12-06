package main

import "fmt"

const (
	Aries       = '\u2648' 
	Taurus      = '\u2649' 
	Gemini      = '\u264a' 
	Cancer      = '\u264b' 
	Leo         = '\u264c' 
	Virgo       = '\u264d' 
	Libra       = '\u264e' 
	Scorpius    = '\u264f' 
	Sagittarius = '\u2650' 
	Capricornus = '\u2651' 
	Aquarius    = '\u2652' 
	Pisces      = '\u2653' 
)

func outputWithZodiacSign(p Person) {
	var zodiacSign rune

	if (p.Month == 3 && p.Day >= 21) || (p.Month == 4 && p.Day <= 20) {
		zodiacSign = Aries
	} else if (p.Month == 4 && p.Day >= 21) || (p.Month == 5 && p.Day <= 20) {
		zodiacSign = Taurus
	} else if (p.Month == 5 && p.Day >= 21) || (p.Month == 6 && p.Day <= 21) {
		zodiacSign = Gemini
	} else if (p.Month == 6 && p.Day >= 22) || (p.Month == 7 && p.Day <= 22) {
		zodiacSign = Cancer
	} else if (p.Month == 7 && p.Day >= 23) || (p.Month == 8 && p.Day <= 23) {
		zodiacSign = Leo
	} else if (p.Month == 8 && p.Day >= 24) || (p.Month == 9 && p.Day <= 23) {
		zodiacSign = Virgo
	} else if (p.Month == 9 && p.Day >= 24) || (p.Month == 10 && p.Day <= 23) {
		zodiacSign = Libra
	} else if (p.Month == 10 && p.Day >= 24) || (p.Month == 11 && p.Day <= 22) {
		zodiacSign = Scorpius
	} else if (p.Month == 11 && p.Day >= 23) || (p.Month == 12 && p.Day <= 21) {
		zodiacSign = Sagittarius
	} else if (p.Month == 12 && p.Day >= 22) || (p.Month == 1 && p.Day <= 20) {
		zodiacSign = Capricornus
	} else if (p.Month == 1 && p.Day >= 21) || (p.Month == 2 && p.Day <= 19) {
		zodiacSign = Aquarius
	} else if (p.Month == 2 && p.Day >= 20) || (p.Month == 3 && p.Day <= 20) {
		zodiacSign = Pisces
	}

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

type FullName struct {
	FirstName string
	LastName  string
}
type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Person struct {
	FullName
	BirthDate
}

func main() {
	grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
	dennis := Person{FullName{"Dennis", "Ritchie"}, BirthDate{9, 9, 1941}}
	rick := Person{FullName{"Rick", "Astley"}, BirthDate{6, 2, 1966}}
	edsger := Person{FullName{"Edsger", "Dijkstra"}, BirthDate{11, 5, 1930}}
	alan := Person{FullName{"Alan", "Turing"}, BirthDate{23, 6, 1912}}

	outputWithZodiacSign(grace)
	outputWithZodiacSign(dennis)
	outputWithZodiacSign(rick)
	outputWithZodiacSign(edsger)
	outputWithZodiacSign(alan)

	Nando := Person{FullName{"Nando", "Brun"}, BirthDate{01, 01, 2005}}
	albert := Person{FullName{"Albert", "Einstein"}, BirthDate{14, 3, 1879}}
	Leandro := Person{FullName{"Leandro", "Schweggler"}, BirthDate{12, 01, 2008}}

	outputWithZodiacSign(Nando)
	outputWithZodiacSign(albert)
	outputWithZodiacSign(Leandro)
}
