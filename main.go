package main

import (
	"algo/bubblesort"
	"algo/selectionsort"
	"fmt"
)

func main() {
	arr := []int{44, 50, 46, 1, 24, 26, 9, 6, 29, 31}

	fmt.Println("Bubble Sort")
	res := bubblesort.BubbleSort(arr)
	fmt.Println(res)

	fmt.Println("Selection Sort")
	res = selectionsort.SelectionSort(arr)
	fmt.Println(res)
}
