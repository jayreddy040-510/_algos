package search

// find the first or last occurence of an element in a sorted array

func lastBsearch(arr []int, target int) int {
	for i, j := 0, len(arr)-1; i <= j; {
		pivot := i + (j-i)/2
		if target == arr[pivot] {
			if pivot == len(arr)-1 || arr[pivot+1] != target {
				return pivot
			} else {
				i = pivot + 1
			}
		} else if target > arr[pivot] {
			i = pivot + 1
		} else {
			j = pivot - 1
		}
	}
	return -1
}

func firstBsearch(arr []int, target int) int {
	for i, j := 0, len(arr)-1; i <= j; {
		pivot := i + (j-i)/2
		if target == arr[pivot] {
			if pivot == 0 || arr[pivot-1] != target {
				return pivot
			} else {
				j = pivot - 1
			}
		} else if target > arr[pivot] {
			i = pivot + 1
		} else {
			j = pivot - 1
		}
	}
	return -1
}
