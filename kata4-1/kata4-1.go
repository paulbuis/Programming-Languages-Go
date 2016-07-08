package main

import (
	"fmt"
	"time"
)

func main() {
	fib := memoFib()

	start := time.Now()
	fib50 := fib(49)
	end := time.Now()
	d := end.Sub(start)
	fmt.Printf("Took %d nanoseconds to compute\n", d.Nanoseconds())
	fmt.Printf("50th Fibbanaci number=%d by recursion (first try)\n", fib50)

	// faster on second try because just one table lookup
	start = time.Now()
	fib50 = fib(49)
	end = time.Now()
	d = end.Sub(start)
	fmt.Printf("\nTook %d nanoseconds to compute\n", d.Nanoseconds())
	fmt.Printf("50th Fibbanaci number=%d by recursion (second try)\n", fib50)
}

func memoFib() func(uint) uint {
	var table []uint
	table = append(table, 0)
	table = append(table, 1)

	var fib func(uint) uint
	fib = func(n uint) uint {
		if n >= uint(len(table)) {
			result := fib(n-2) + fib(n-1)
			table = append(table, result)
		}
		return table[n]
	}
	return fib
}
