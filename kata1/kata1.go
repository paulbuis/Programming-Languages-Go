package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	fmt.Println("Input 3 integers on one line & hit [return]")
	fmt.Scanf("%d %d %d", &x, &y, &z)
	sum := x + y + z
	ave := float64(sum) / 3.0
	fmt.Printf("\nAverage: %6.3g\n", ave)
}
