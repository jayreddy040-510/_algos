package search

import "math"

func minRotatedBsearch(arr []int) int {
	min := int(math.Inf(-1))
	for i, j := 0, len(arr)-1; i <= j; {
		pivot := i + (j-i)/2
		if arr[pivot] < min {
			min = arr[pivot]
		}
		if arr[i] < arr[pivot] {
			j = pivot - 1
		} else {
			i = pivot + 1
		}
	}
	return min
}
