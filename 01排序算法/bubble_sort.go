// 选择排序 时间复杂度O(N^2)
package main

import (
	"fmt"
	"leet_code_go/helper"
)

//func BubbleSortAsc(array []int) []int {
//	for i := 0; i < len(array)-1; i++ {
//		for j := i + 1; j <= len(array)-1; j++ {
//			if array[i] > array[j] {
//				array[i], array[j] = array[j], array[i]
//			}
//		}
//	}
//	return array
//}

func BubbleSortAsc(array []int) []int {
	for i := len(array) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func BubbleSortDesc(array []int) []int {
	for i := len(array) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}

func main() {
	var array []int
	array = helper.RandomIntArray(10, 1000)
	fmt.Println(array)
	ascResult := BubbleSortAsc(array)
	fmt.Println(ascResult)

	array = helper.RandomIntArray(10, 100)
	fmt.Println(array)
	descResult := BubbleSortDesc(array)
	fmt.Println(descResult)
}
