package main

import (
	"fmt"
	"time"

	"github.com/Ma-Leal/algorithms-go/searching"
)

func main() {
	sortedList := generateNumbers(10000)
	item := 9999

	func() {
		start := time.Now()
		exists := searching.BinarySearch(sortedList, item)
		duration := time.Since(start)
		fmt.Printf("BinarySearch => Target number: %v | Found: %v | Execution time: %v\n", item, exists, duration)
	}()

	func() {
		start := time.Now()
		exists := searching.LinearSearch(sortedList, item)
		duration := time.Since(start)
		fmt.Printf("LinearSearch => Target number: %v | Found: %v | Execution time: %v\n", item, exists, duration)
	}()
}

func generateNumbers(count int) []int {
	numbers := []int{}
	for i := 0; i <= count; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}
