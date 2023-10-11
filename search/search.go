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

func RotatedBSearch(arr []int, target int) int {
	// given a sorted arr rotated at random pivot, implement bsearch
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if target == arr[mid] {
			return mid
		} else if arr[lo] <= arr[mid] {
			if arr[lo] <= target && target < arr[mid] {
				hi = mid
			} else {
				lo = mid + 1
			}
		} else {
			if arr[mid] < target && target <= arr[hi] {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	}
	return -1 // sentinel value
}
