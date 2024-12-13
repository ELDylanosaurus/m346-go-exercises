package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float32) float32 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

type ShortSides struct {
	a, b float32
}

func (s ShortSides) Hypotenuse() float32 {
	return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}

func main() {
	fmt.Println("Using computeHypotenuse:")
	fmt.Printf("Hypotenuse for sides 3 and 4: %.2f\n", computeHypotenuse(3, 4))  // 5.00
	fmt.Printf("Hypotenuse for sides 5 and 12: %.2f\n", computeHypotenuse(5, 12)) // 13.00
	fmt.Printf("Hypotenuse for sides 8 and 15: %.2f\n", computeHypotenuse(8, 15)) // 17.00

	fmt.Println("\nShortSides.Hypotenuse:")
	side1 := ShortSides{3, 4}
	side2 := ShortSides{5, 12}
	side3 := ShortSides{8, 15}

	fmt.Printf("Hypotenuse for sides %v: %.2f\n", side1, side1.Hypotenuse()) // 5.00
	fmt.Printf("Hypotenuse for sides %v: %.2f\n", side2, side2.Hypotenuse()) // 13.00
	fmt.Printf("Hypotenuse for sides %v: %.2f\n", side3, side3.Hypotenuse()) // 17.00
}
