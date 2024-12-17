package sorting

func SelectionSort(arr []int) []int {
	newArr := []int{}

	for len(arr) > 0 {

		minIndex := minValueIndex(arr)
		newArr = append(newArr, arr[minIndex])
		arr = append(arr[:minIndex], arr[minIndex+1:]...)
	}

	return newArr
}

func minValueIndex(arr []int) int {
	minValue := arr[0]
	minIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
			minIndex = i
		}
	}

	return minIndex
}
