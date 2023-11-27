package search

// bsearch in an array where array is roughly sorted but can be exchanged with
// any adjacent index e.g. [1, 3, 5, 2, 10, 7]
func bsearchNearlySortedArray(arr []int, target int) int {
	for i, j := 0, len(arr)-1; i <= j; {
		pivot := i + (j-i)/2
		if target == arr[pivot] {
			return pivot
		} else if pivot+1 < len(arr) && target == arr[pivot+1] {
			return pivot + 1
		} else if pivot-1 >= 0 && target == arr[pivot-1] {
			return pivot - 1
		} else if target < arr[pivot] {
			j = pivot - 1
		} else {
			i = pivot + 1
		}
	}
	return -1
}
