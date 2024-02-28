package main

import (
	"algo/bubblesort"
	"fmt"
)

func main() {
	res := bubblesort.BubbleSort([]int{29, 10, 14, 37, 14})
	fmt.Println(res)
}
