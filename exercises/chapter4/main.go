package main

import (
	"fmt"

	"github.com/Ma-Leal/algorithms-go/sorting"
)

func main() {

	original := []int{
		68, 85, 24, 50, 34, 41, 93, 13, 82, 56,
	}

	sorted := sorting.QuickSort(original)

	fmt.Printf("Original array: %v\nSorted   array: %v\n", original, sorted)

}
