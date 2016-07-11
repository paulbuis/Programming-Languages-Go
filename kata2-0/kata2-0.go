package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter integers, one per line, terminate with blank line")
	data := ReadInts()
	ave := AveInts(data)
	fmt.Printf("\nAverage: %6.3g\n", ave)
}

// ReadInts reads a sequence of integers from standard input,
// one integer per line
// until an error occurs (e.g., a blank line)
// and returns the sequence in a slice
func ReadInts() (result []int) {
	var item int
	nRead, err := fmt.Scanf("%d\n", &item)
	for nRead > 0 && err == nil {
		result = append(result, item)
		nRead, err = fmt.Scanf("%d\n", &item)
	}
	return
}

// AveInts takes a slice of integers and return their mean.
// Will return NaN if slice is empty
func AveInts(items []int) (result float64) {
	len := len(items)
	sum := 0
	for i := 0; i < len; i++ {
		sum += items[i]
	}
	result = float64(sum) / float64(len)
	return
}
