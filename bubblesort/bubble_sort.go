package bubblesort

func BubbleSort(arr []int) []int {
	for i := range arr {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}

	return arr
}
