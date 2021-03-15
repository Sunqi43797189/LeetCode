package main

import (
	"fmt"
	"leet_code_go/helper"
)

func QuickSort(array []int) []int {
	partition(array, 0, len(array)-1)
	return array
}

func partition(array []int, left, right int) {
	if left >= right {
		return
	}
	pivot := right
	var l, r = left, right - 1
	for l < r {
		for l < pivot && array[l] <= array[pivot] {
			l++
		}
		for r > left && array[r] >= array[pivot] {
			r--
		}

		if l < r {
			array[l], array[r] = array[r], array[l]
		}
	}
	array[l], array[pivot] = array[pivot], array[l]
	partition(array, left, l-1)
	partition(array, l+1, right)
}
func main() {
	array := helper.RandomIntArray(10, 1000)
	fmt.Println(array)
	result := QuickSort(array)
	fmt.Println(result)
}
