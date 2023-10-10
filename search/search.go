package search

func linearSearch(arr []int, target int) int {
	for idx, num := range arr {
		if num == target {
			return idx
		}
	}
	return -1
}

func BSearch(arr []int, target int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			lo = mid + 1
		} else if target < arr[mid] {
			hi = mid
		}
	}
	return -1
}
