# Programming-Languages-Go
Simple examples of Go Programs

## Kata 0
_Print a simple greeting and current time_
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello!")
	fmt.Println("The time is", time.Now())
}
```

Go programs start in function `main` of package `main`.

Source code from other packages is parsed for declarations via the import statement at the beginning of each go source code file.


## Kata 1
_Read in 3 integers and print their average_
```go
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
```

## Kata 2
_Read in a sequence of integers and compute their average._

### Level 0
_Store sequence in a slice with separate functions for input and computation._

```go
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

func ReadInts() (result []int) {
	var item int
	nRead, err := fmt.Scanf("%d\n", &item)
	for nRead > 0 && err == nil {
		result = append(result, item)
		nRead, err = fmt.Scanf("%d\n", &item)
	}
	return
}
```

### Level 1
_Use producer-consumer model:  
*producer reads integers  
*consumer computes average_
  
```go
package main

import (
	"fmt"
)

func main() {
	var aveChannel = make(chan float64)
	var dChannel = make(chan int)
	fmt.Println("Enter integers, one per line, terminate with blank line")
	go produce(dChannel)
	go consume(dChannel, aveChannel)

	fmt.Println("producer/consumer launched")
	ave := <-aveChannel
	fmt.Println("average = ", ave)
}

func produce(dataChannel chan<- int) {
	defer close(dataChannel)
	var item int
	nRead, err := fmt.Scanf("%d\n", &item)
	for nRead > 0 && err == nil {
		dataChannel <- item
		nRead, err = fmt.Scanf("%d\n", &item)
	}
}

func consume(dataChannel <-chan int, aveChannel chan<- float64) {
	sum := 0
	count := 0

	for item := range dataChannel {
		sum += item
		fmt.Println("sum=", sum)
		count++
	}
	aveChannel <- float64(sum) / float64(count)
}
```
  
## Kata 3
_Recursively compute factorial_
```go
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
```

## Kata 4
_Compute 50th Fibanacci number using a closure
and attempt performance measurement._
### Level 0
_Use iteration_

```go
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
```


### Level 1
_Use recursion with memoization_
```go
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
```
