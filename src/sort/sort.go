package sort

func bubbleSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if arr[j-1] < arr[j] {
				tmp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for j = i; j > 0 && arr[j-1] < tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
	return arr
}

func hillSort(arr []int) []int {
	for d := len(arr); d > 0; d = d/2 - 1 {
		for i := d; i < len(arr); i++ {
			tmp := arr[i]
			j := i
			for j = i; j > d && arr[j-d] < tmp; j-- {
				arr[j] = arr[j-1]
			}
			arr[j] = tmp
		}
	}
	return arr
}

func stackSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minPosition := scanformin(arr, i, len(arr)-1)
		swap(&arr[i], &arr[minPosition])
	}
	return arr
}

func swap(a *int, b *int) {
	tmp := a
	a = b
	b = tmp
}
func scanformin(arr []int, i int, N int) int {

	return 0
}
