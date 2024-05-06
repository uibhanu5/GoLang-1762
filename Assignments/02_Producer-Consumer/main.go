package main

import "fmt"

func main() {
	numbers := make(chan int)

	go producer(numbers)
	go consumer(numbers)

	// This is to Prevent main from exiting until consumer is done processing
	var input string
	fmt.Scanln(&input)
}

func producer(numbers chan<- int) {
	for i := 1; i <= 100; i++ {
		numbers <- i
	}
	close(numbers)
}

func consumer(numbers <-chan int) {
	for num := range numbers {
		fmt.Println("Consumed:", num)
	}
}
