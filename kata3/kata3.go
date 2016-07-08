package main

import (
	"fmt"
)

func main() {
	fmt.Println("5!=", factorial(5))
	fmt.Printf("50!=%v\n", factorial(50))
}

func factorial(n uint16) float64 {
	if n == 0 {
		return 1.0
	}
	return float64(n) * factorial(n-1)
}
