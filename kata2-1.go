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
