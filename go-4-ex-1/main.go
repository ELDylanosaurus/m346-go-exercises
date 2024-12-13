package main

import (
	"errors"
	"fmt"
)

func computeGrade(gotPoints, maxPoints float64) (float64, error) {
	if gotPoints < 0 || maxPoints <= 0 || gotPoints > maxPoints {
		return 0, errors.New("invalid points: gotPoints must be between 0 and maxPoints, and maxPoints must be positive")
	}

	grade := 1.0 + 5.0*(gotPoints/maxPoints)
	return grade, nil
}

func main() {
	grades := []struct {
		gotPoints float64
		maxPoints float64
	}{
		{17.5, 28.0}, // 4.1
		{25.0, 30.0}, // 5.2
		{0.0, 20.0},  // 1.0 
		{-5.0, 30.0}, // Error
		{35.0, 30.0}, // Error
	}

	for _, g := range grades {
		grade, err := computeGrade(g.gotPoints, g.maxPoints)
		if err != nil {
			fmt.Printf("Error for %v/%v: %v\n", g.gotPoints, g.maxPoints, err)
		} else {
			fmt.Printf("Grade for %v/%v: %.3f\n", g.gotPoints, g.maxPoints, grade)
		}
	}
}
