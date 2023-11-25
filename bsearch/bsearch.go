package search

import "archive/tar"

func linearSearch(arr []int, target int) int {
	for idx, num := range arr {
		if num == target {
			return idx
		}
	}
	return -1
}

func BSearch(arr []int, target int) int {
	for i, j := 0, len(arr)-1; i <= j; {
		pivot := i + (j-i)/2
		if arr[pivot] == target {
			return pivot
		} else if arr[pivot] < target {
			i = pivot + 1
		} else {
			j = pivot - 1
		}
	}
	return -1
}
