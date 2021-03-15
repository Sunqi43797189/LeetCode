package main

import (
	"fmt"
	"leet_code_go/helper"
	"sort"
)

func QuickSort(array []int, left, right int) []int {
	if left >= right {
		return array
	}
	pivot := partition(array, left, right)
	QuickSort(array, left, pivot-1)
	QuickSort(array, pivot+1, right)
	return array
}

func partition(array []int, left, right int) int {
	pivot := array[right]
	var l, r = left, right - 1
	for l <= r {
		for l <= r && array[l] <= pivot {
			l++
		}
		for l <= r && array[r] > pivot {
			r--
		}

		if l < r {
			array[l], array[r] = array[r], array[l]
		}
	}
	array[l], array[right] = array[right], array[l]
	return l
}
func main() {
	array := helper.RandomIntArray(1000, 1000)
	newArray := make([]int, len(array))
	for index, value := range array {
		newArray[index] = value
	}
	result := QuickSort(newArray, 0, len(newArray)-1)
	sort.Ints(array)
	fmt.Println(result)
	for index, value := range array {
		if value != result[index] {
			fmt.Println(index)
			panic("结果不一样")
		}
	}
}
