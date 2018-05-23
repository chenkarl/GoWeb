package sort

import (
	"testing"
)

var arr1 = []int{4, 981, 10, -17, 0, -20, 29, 50, 8, 43, -5}

func TestBubbleSort(t *testing.T) {
	result := bubbleSort(arr1)
	t.Log(result)
}

func TestInsertSort(t *testing.T) {
	result := insertSort(arr1)
	t.Log(result)
}

func TestHillSort(t *testing.T) {
	result := hillSort(arr1)
	t.Log(result)
}
