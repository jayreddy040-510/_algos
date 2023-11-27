package search

// find an element in an array that is larger than both ele adjacent to it

func findPeak(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}
	for i, j := 0, len(arr)-1; i <= j; {
		pivot := i + (j-i)/2
		if (pivot == 0 || arr[pivot-1] <= arr[pivot]) && (pivot == len(arr)-1 || arr[pivot+1] <= arr[pivot]) {
			return pivot
		}
		if pivot > 0 && arr[pivot-1] > arr[pivot] {
			j = pivot - 1
		} else {
			i = pivot + 1
		}
	}
	return -1
}
