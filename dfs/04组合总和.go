package main

import (
	"fmt"
	"sort"
)

func DfsCountTarget(arr []int, target int, res *[][]int, tempArr []int) {
	if target <= 0 {
		if target == 0 {
			sort.Ints(tempArr)
			a := make([]int, len(tempArr))
			copy(a, tempArr)
			*res = append(*res, a)
		}
		return
	}

	for _, value := range arr {
		tempArr = append(tempArr, value)
		DfsCountTarget(arr, target - value, res, tempArr)
		tempArr = tempArr[:len(tempArr)-1]
	}
}

func main() {
	arr := []int{2, 3, 6, 7}
	target := 7
	var res [][]int
	DfsCountTarget(arr, target, &res, []int{})
	fmt.Println(res)
}

func SumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func IsStringsArrayContain(array []string, item string) bool {
	for _, str := range array {
		if str == item {
			return true
		}
	}
	return false
}
