package search

import (
	"testing"
)

func TestBinarysearch(t *testing.T) {
	arr := []int{12, 31, 55, 89, 101}
	num := 31
	loc := binarysearch(arr, num)
	if loc == 1 {
		t.Log("查找成功")
	} else {
		t.Log("查找失败")
	}
}
