package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var eyes = rand.Intn(6) + 1
	var when = time.Now()

	fmt.Fprintln(os.Stdout, "The dice shows", eyes, "eyes")

	fmt.Fprintln(os.Stderr, "The dice was rolled at", when)

}
