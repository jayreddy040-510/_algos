package search

// bsearch a sorted array rotated around one pivot point e.g. [5, 7, 1, 3]

func rotatedBsearch(arr []int, target int) int {
    for i, j := 0, len(arr) - 1; i <= j; {
        pivot := i + (j-i)/2
        if target == arr[pivot] {
            return pivot
        } else if target > arr[pivot] {
            if arr[j] > arr[pivot] {
                i = pivot + 1
            } else {
                j = pivot - 1
            }
        } else {
            if arr[i] < arr[pivot] {
                j = pivot - 1
            } else {
                i = pivot + 1
            }
        }
    }
	return -1
}
