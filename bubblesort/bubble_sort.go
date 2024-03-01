package bubblesort

// Start from left. Iterate through each item. Swap if left is larger than left
// No need to go through the sorted items at the end
func BubbleSort(arr []int) []int {
	for i := range arr {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
