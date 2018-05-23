package search

func binarysearch(arr []int, num int) int {
	right := len(arr) - 1
	left := 0
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < num {
			left = mid + 1
		} else if arr[mid] > num {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
