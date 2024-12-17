package searching

func LinearSearch(sortedList []int, item int) bool {
	// O(n)
	for _, v := range sortedList {
		if v == item {
			return true
		}
	}
	return false
}
