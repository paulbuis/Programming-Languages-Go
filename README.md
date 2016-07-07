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

`x`, `y`, and `z` have their type explicitly declared using the `var` keyword.
`sum` and `ave` have their types inferred from the expression on the right hand side of their initializing expressions by using the `:=` operator which can only be used in such initializations.

As in C, the `&` is the address-of operator, allowing `Scanf` to modify their values.

No mixed-mode arithmetic is allowed, so `sum` must be explicitly converted to match type of `3.0`.

C-like formatting is done with `Printf`. 

Note identifiers from other packages are available for use only if they are capitalized. Access control in Go is
managed without keywords like "public" or "private."

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

func AveInts(items []int) (result float64) {
	len := len(items)
	sum := 0
	for i := 0; i < len; i++ {
		sum += items[i]
	}
	result = float64(sum) / float64(len)
	return
}
```

Here we see an example of defining and calling functions.

One way to describe the value returned by a function is shown here. Both a name and a type are specified
in a list. The `return` uses the value stored in the specified name.

Note also that the `for` statement does not require parentheses. It is used in the C/Java-like 3-part form or
like a C/Java `while` with just a condition. Go does not have a `while` keyword.

Pay careful attention to the form of the return value of `Scanf`. It is returning a 
[tuple](https://en.wikipedia.org/wiki/Tuple). The first value in
the tuple is the number of values successfully read. The second value in the tuple is of type `Error` with
the value of `nil` indicating no error. Returning a tuple with the last value being of type `Error` is a Go idiom.
This allows the first part of the tuple to be unrelated to success or failure, with no special values needing
to be realated to failure indicators. Go also has no "Exception" returns or `try`/`catch` constructs which aids
in efficient implementation of function call/return and simpler semantics.

### Level 1
_Use producer-consumer model:_  
+producer reads integers
+consumer computes average
  
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

This example illustrates the primary constructs for concurrency in Go. A channel is essentially a producer-consumer queue.

The built-in `make` function is used to construct built-in, non-primitive types in Go.
Here it is being used to construct two unbuffered channels, one to transmit integers and the other for floating point numbers.
The keyword `var` is used in conjunction with initializations involving `make` rather than the `:=` operator.

The `go` keyword launches a function in a separate thread. Go refers to these as _*goroutines*_.

A function parameter with a `<-chan` indicates the function can only read from the channel. A value is extracted from
a channel an assigned to a variable with the `<-` operator being used instead of a `=` operator, with the channel on the
right hand side.

A function parameter with a `chan<-` indicates the function can only write to the channel. A value is written to a channel
with the channel on the left hand side and the expression to be written to the channel on the right hand side, again with
syntax similar to a simple assignment statement.

The keyword `range` used in the `for` loop in `consume()` indicates the channel should be read until the writer closes it.
The `defer` statement in `produce()` arranges for the channel to be closed when `produce()` exits.
`defer` is used to ensure a function cleans up after itself when it exits, even if it exits abnormally.
`defer` also allows such cleanup to be specified in a line of code right next to the line that creates a
resource that will need cleanup, rather than farther away where the function is dealing with returning a value.
Since a function may have multiple `return` statements, being able to use `defer` makes sure a cleanup
activity only needs to be specified once.

This particular style of using a goroutine to send values to a channel to another function consuming those values with
a `for`/`range` construct is how Go does [_*generators*_](https://en.wikipedia.org/wiki/Generator_(computer_programming))
(Python and C# use the `yield` keyword to construct generators). Generators are not part of the Java 8 language or standard
library.

In this example `produce()`, `consume()`, and `main()` are all running concurrently.
The Go runtime manages mapping "goroutines" to threads and may use a thread pool to use
fewer threads than simultaneously active goroutines.
Each goroutine will have its own stack.
These stacks are "segmented" meaning they do not require large contiguous blocks of memory,
instead consisting of a linked structure of chunks of memory.
This allows many simultaneous goroutines to be active without consuming lots of memory to support them.


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

Like `for`, `if` does not require parentheses around the conditional expression. Curly braces, however,
are required to group a sequence of statements, even if the sequence is only a single statement.


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

Here is an example of an anonymous function (a.k.a. a "lambda"),
nested inside another function, which returns the inner function.

We also see here an example of tuple assignment in the body of the inner function.
Without tuple assignment, an extra temporary variable would need to be introduced.

The subtle thing is that the inner function uses variables that are local to the
outer function. This makes the inner function
a [_*closure*_](https://en.wikipedia.org/wiki/Closure_(computer_programming)).
It uses these
variables even after the outer function has returned, indicating that the
[stack frame / activation record](https://en.wikipedia.org/wiki/Call_stack#ACTIVATION-RECORD)
in which they reside has persisted. These values, however,
are not global to the whole program, package, or file and another instance of
them would be created if `closureFib()` were invoked again.

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
