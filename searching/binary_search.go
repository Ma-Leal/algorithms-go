package searching

func BinarySearch(sortedList []int, item int) bool {
	// O(log n)
	min, max := 0, len(sortedList)-1

	for min <= max {
		mid := (min + max) / 2
		if sortedList[mid] == item {
			return true
		}
		if sortedList[mid] > item {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	return false
}
