package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fib0 := closureFib()
	for i := 0; i < 49; i++ {
		fib0()
	}
	f := fib0()
	end := time.Now()
	d := end.Sub(start)
	fmt.Printf("Took %d nanoseconds to compute\n", d.Nanoseconds())
	fmt.Printf("50th Fibbanaci number=%v by iteration sequence\n", f)
}

func closureFib() func() uint {
	var a, b uint = 0, 1
	fib := func() uint {
		result := a
		a, b = b, a+b
		return result
	}
	return fib
}
