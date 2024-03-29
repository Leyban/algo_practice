package insertionsort_test

import (
	"algo/insertionsort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	arr10    = []int{73, 91, 62, 67, 16, 80, 55, 60, 93, 31}
	arr100   = []int{2, 2, 58, 4, 66, 96, 68, 22, 55, 1, 93, 94, 31, 90, 39, 28, 2, 68, 93, 10, 71, 70, 64, 64, 14, 75, 1, 3, 74, 93, 25, 38, 47, 97, 37, 58, 0, 90, 73, 4, 29, 84, 16, 41, 65, 18, 90, 44, 34, 73, 25, 79, 65, 62, 88, 23, 24, 87, 83, 95, 86, 0, 73, 16, 9, 40, 48, 11, 28, 44, 67, 57, 12, 89, 49, 99, 37, 84, 53, 25, 4, 69, 62, 69, 23, 36, 91, 63, 47, 12, 33, 93, 84, 96, 71, 9, 33, 91, 97, 86}
	arr1000  = []int{55, 59, 98, 38, 72, 17, 43, 2, 79, 63, 85, 85, 38, 83, 68, 90, 7, 20, 94, 92, 21, 0, 48, 34, 97, 65, 43, 5, 99, 96, 5, 100, 16, 84, 62, 80, 63, 100, 69, 34, 98, 4, 15, 63, 46, 67, 57, 83, 29, 65, 36, 29, 3, 83, 58, 100, 65, 80, 52, 30, 5, 5, 98, 92, 30, 20, 73, 51, 84, 33, 38, 54, 97, 46, 47, 31, 78, 93, 35, 39, 59, 79, 98, 45, 65, 0, 73, 16, 34, 32, 14, 78, 51, 54, 27, 93, 95, 95, 73, 23, 44, 29, 12, 94, 98, 63, 88, 85, 0, 9, 89, 35, 65, 12, 28, 37, 90, 78, 58, 67, 64, 10, 69, 90, 1, 32, 2, 47, 57, 97, 5, 94, 1, 19, 23, 96, 53, 61, 35, 97, 15, 53, 56, 42, 23, 88, 38, 54, 98, 94, 26, 2, 95, 23, 63, 24, 82, 78, 14, 57, 61, 85, 100, 31, 20, 36, 61, 31, 8, 41, 45, 72, 87, 66, 27, 0, 27, 72, 39, 47, 5, 71, 4, 72, 43, 26, 35, 52, 9, 47, 73, 6, 10, 52, 74, 3, 30, 42, 43, 10, 80, 52, 11, 33, 80, 86, 42, 17, 94, 10, 97, 97, 0, 69, 43, 34, 17, 38, 77, 76, 65, 89, 15, 26, 62, 2, 22, 4, 64, 89, 48, 19, 43, 62, 91, 94, 40, 99, 77, 95, 50, 49, 4, 11, 32, 1, 38, 76, 76, 55, 53, 62, 13, 100, 66, 48, 79, 88, 92, 51, 67, 100, 98, 27, 78, 43, 77, 69, 12, 36, 30, 30, 61, 80, 18, 74, 77, 31, 19, 68, 88, 88, 64, 23, 38, 97, 51, 25, 20, 72, 79, 23, 19, 89, 47, 25, 50, 53, 95, 5, 92, 56, 55, 37, 19, 0, 69, 4, 41, 44, 58, 1, 50, 14, 42, 38, 45, 79, 17, 87, 56, 59, 23, 71, 86, 33, 91, 54, 18, 48, 42, 36, 38, 33, 78, 90, 74, 88, 16, 8, 27, 60, 27, 92, 16, 84, 70, 45, 56, 88, 44, 23, 100, 12, 30, 25, 49, 49, 38, 15, 54, 39, 48, 44, 18, 70, 50, 43, 6, 5, 22, 31, 52, 92, 41, 30, 65, 52, 1, 1, 12, 51, 38, 10, 90, 29, 97, 47, 100, 55, 60, 58, 19, 18, 26, 42, 16, 26, 45, 24, 83, 14, 38, 64, 70, 88, 66, 55, 35, 27, 22, 50, 51, 87, 48, 7, 92, 7, 27, 31, 66, 92, 8, 20, 95, 68, 80, 8, 46, 78, 0, 94, 76, 58, 59, 50, 81, 75, 46, 98, 14, 61, 9, 49, 81, 8, 79, 87, 11, 57, 96, 95, 48, 7, 89, 39, 97, 62, 71, 8, 75, 62, 32, 22, 73, 40, 63, 98, 78, 95, 9, 6, 77, 22, 93, 56, 93, 16, 22, 32, 10, 15, 44, 35, 27, 60, 95, 37, 67, 49, 81, 3, 84, 45, 29, 45, 68, 40, 28, 92, 15, 29, 76, 40, 39, 40, 100, 40, 47, 28, 85, 80, 56, 57, 26, 51, 39, 84, 63, 27, 10, 60, 37, 56, 10, 46, 0, 7, 59, 57, 70, 94, 67, 97, 67, 27, 74, 46, 41, 17, 29, 34, 64, 57, 29, 81, 41, 47, 25, 31, 33, 42, 46, 57, 21, 84, 31, 85, 83, 69, 15, 93, 42, 66, 58, 47, 81, 46, 79, 86, 20, 15, 94, 19, 13, 59, 57, 94, 99, 52, 45, 26, 95, 44, 35, 68, 95, 83, 20, 16, 45, 36, 75, 9, 53, 22, 30, 33, 47, 74, 91, 24, 60, 2, 28, 41, 96, 91, 25, 76, 88, 47, 19, 23, 98, 63, 79, 32, 25, 33, 74, 57, 6, 50, 39, 8, 27, 65, 11, 8, 4, 5, 77, 92, 74, 86, 32, 83, 71, 68, 70, 27, 67, 30, 81, 98, 48, 22, 59, 70, 52, 87, 41, 32, 94, 8, 78, 97, 13, 93, 77, 10, 86, 13, 68, 1, 69, 74, 2, 10, 53, 0, 53, 65, 85, 2, 52, 21, 79, 29, 62, 76, 41, 59, 86, 81, 88, 55, 67, 13, 43, 92, 93, 46, 36, 75, 24, 93, 43, 43, 83, 61, 54, 11, 60, 47, 11, 41, 13, 100, 85, 83, 0, 46, 46, 89, 49, 62, 61, 56, 9, 14, 31, 38, 12, 66, 66, 79, 0, 6, 15, 34, 49, 75, 31, 70, 92, 5, 75, 64, 57, 24, 22, 80, 60, 34, 38, 38, 0, 70, 72, 86, 39, 28, 94, 82, 9, 0, 32, 2, 39, 69, 43, 16, 68, 84, 84, 17, 26, 78, 4, 51, 71, 57, 77, 66, 89, 70, 46, 0, 27, 58, 84, 14, 74, 68, 15, 42, 10, 9, 33, 47, 59, 70, 57, 69, 92, 44, 58, 62, 32, 62, 57, 66, 6, 50, 25, 44, 31, 54, 98, 11, 37, 2, 36, 65, 59, 52, 81, 90, 65, 40, 33, 4, 88, 21, 17, 70, 96, 2, 39, 14, 87, 2, 9, 62, 18, 27, 35, 63, 21, 22, 19, 36, 13, 32, 2, 54, 99, 98, 81, 88, 72, 17, 39, 6, 26, 17, 45, 64, 74, 6, 43, 79, 10, 70, 16, 75, 16, 25, 52, 52, 19, 11, 62, 0, 6, 64, 2, 81, 92, 96, 12, 80, 70, 54, 56, 14, 85, 82, 22, 97, 81, 84, 36, 64, 14, 98, 5, 41, 91, 25, 33, 62, 11, 85, 81, 82, 40, 56, 3, 72, 32, 90, 14, 50, 92, 34, 85, 29, 41, 66, 97, 0, 7, 15, 52, 85, 3, 36, 32, 100, 14, 74, 46, 15, 65, 100, 60, 71, 46, 19, 35, 7, 30, 15, 75, 16, 72, 27, 97, 71, 2, 68, 96, 84, 21, 66, 35, 10, 39, 49, 65, 34, 49, 46, 91, 16, 79, 21, 26, 32, 53, 64, 23, 59, 43, 95, 99, 86, 40, 57, 97, 37, 82, 67, 36, 15, 40, 7, 38, 64, 82, 2, 24, 39, 36, 21, 99, 17}
	sortFunc = insertionsort.InsertionSort
)

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Should return empty arr",
			args: []int{},
			want: []int{},
		},
		{
			name: "Should return same arr",
			args: []int{1},
			want: []int{1},
		},
		{
			name: "Should return sorted arr",
			args: []int{5, 6, 8, 9, 7, 10, 3, 1, 2, 4},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "Should return sorted arr with duplicates",
			args: []int{5, 5, 8, 9, 9, 10, 3, 1, 2, 4},
			want: []int{1, 2, 3, 4, 5, 5, 8, 9, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortFunc(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}

func benchmarkSort(arr []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sortFunc(arr)
	}
}

func BenchmarkSort10(b *testing.B)   { benchmarkSort(arr10, b) }
func BenchmarkSort100(b *testing.B)  { benchmarkSort(arr100, b) }
func BenchmarkSort1000(b *testing.B) { benchmarkSort(arr1000, b) }
