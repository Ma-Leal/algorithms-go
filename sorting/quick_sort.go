package sorting

import (
	"math/rand/v2"
)

func QuickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	} else {

		pivotIndex := rand.IntN(len(arr))
		pivot := arr[pivotIndex]
		var min, max []int

		for _, v := range arr {
			if pivot == v {
				continue
			}
			if v > pivot {
				max = append(max, v)
			} else {
				min = append(min, v)
			}
		}

		return append(append(QuickSort(min), pivot), QuickSort(max)...)

	}

}
