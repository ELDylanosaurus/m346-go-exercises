package main

import "fmt"

func convertCelsiusToFahrenheit(celsius float32) float32 {
	return celsius*9/5 + 32
}

func convertFahrenheitToCelsius(fahrenheit float32) float32 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	c1 := 0.0
	c2 := 25.0
	c3 := -10.0
	fmt.Printf("%.2f°C = %.2f°F\n", c1, convertCelsiusToFahrenheit(c1)) // 32.00°F
	fmt.Printf("%.2f°C = %.2f°F\n", c2, convertCelsiusToFahrenheit(c2)) // 77.00°F
	fmt.Printf("%.2f°C = %.2f°F\n", c3, convertCelsiusToFahrenheit(c3)) // 14.00°F

	f1 := 32.0
	f2 := 77.0
	f3 := 14.0
	fmt.Printf("%.2f°F = %.2f°C\n", f1, convertFahrenheitToCelsius(f1)) // 0.00°C
	fmt.Printf("%.2f°F = %.2f°C\n", f2, convertFahrenheitToCelsius(f2)) // 25.00°C
	fmt.Printf("%.2f°F = %.2f°C\n", f3, convertFahrenheitToCelsius(f3)) // -10.00°C
}
