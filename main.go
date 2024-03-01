package main

import (
	"algo/bubblesort"
	"algo/selectionsort"
	"fmt"
)

func main() {
	fmt.Println("Bubble Sort")
	arr := []int{5, 6, 8, 9, 7, 10, 3, 1, 2, 4}
	fmt.Println(bubblesort.BubbleSort(arr))

	fmt.Println("Selection Sort")
	arr = []int{5, 6, 8, 9, 7, 10, 3, 1, 2, 4}
	fmt.Println(selectionsort.SelectionSort(arr))
}
