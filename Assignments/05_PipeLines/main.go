package main

import "fmt"

func main() {
	fmt.Println("Pipeline started: Generating numbers...")
	numbers := generator(1, 5)
	fmt.Println("Numbers generated: Squaring the numbers...")
	squared := square(numbers)
	fmt.Println("Numbers squared: Printing the squared numbers...")
	printer(squared)
	fmt.Println("Pipeline completed")
}

func generator(start, end int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func printer(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}
