package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float32) ([]float32, error) {
	if a == 0 {
		return nil, fmt.Errorf("a cant be zero")
	}

	discriminant := math.Pow(b, 2) - 4*a*c

	switch {
	case discriminant > 0:
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return []float32{x1, x2}, nil
	case discriminant == 0:
		x := -b / (2 * a)
		return []float32{x}, nil
	default:
		return nil, fmt.Errorf("no solutions")
	}
}

func main() {
	testCases := []struct {
		a, b, c float32
	}{
		{3, 4, 1}, // L = {-0,3333333; -1}
		{2, 4, 2}, // L = {-1}
		{3, 4, 2}, // L = {}
	}

	for _, tc := range testCases {
		fmt.Printf(" a=%.2f, b=%.2f, c=%.2f:\n", tc.a, tc.b, tc.c)
		solutions, err := computeQuadraticFormula(tc.a, tc.b, tc.c)
		if err != nil {
			fmt.Println("  Error:", err)
		} else {
			fmt.Println("  Solutions:", solutions)
		}
	}
}
