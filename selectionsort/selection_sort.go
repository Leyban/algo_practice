package selectionsort

// Start from the left. Find the smalles possible min value for each position.
// Iterate through all items to find the minimum value
// Swap the minimum with the current position in question
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minimum := i
		for j := i; j < len(arr)-1; j++ {
			if arr[minimum] > arr[j] {
				minimum = j
			}
		}
		if arr[minimum] < arr[i] {
			temp := arr[minimum]
			arr[minimum] = arr[i]
			arr[i] = temp
		}
	}

	return arr
}
