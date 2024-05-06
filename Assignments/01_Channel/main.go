package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Enter the value of N: ")
	fmt.Scanln(&N)

	ch := make(chan int)

	go calculateSum(1, N/2, ch)
	go calculateSum(N/2+1, N, ch)

	sum1 := <-ch
	sum2 := <-ch

	totalSum := sum1 + sum2

	fmt.Println("Total sum of numbers from 1 to", N, "is:", totalSum)
}

func calculateSum(start, end int, c chan int) {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	c <- sum
}
