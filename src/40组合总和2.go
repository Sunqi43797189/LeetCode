package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var cacheArr []string
	sort.Ints(candidates)
	boolArr := make([]bool, len(candidates))
	dfsCombinationSum2(candidates, &res, target, []int{}, &cacheArr, boolArr)
	return res
}

func dfsCombinationSum2(arr []int, res *[][]int, target int, tempArr []int, cacheArr *[]string, boolArr []bool) {
	if target <= 0 {
		if target == 0 {
			a := make([]int, len(tempArr))
			copy(a, tempArr)
			//sort.Ints(a)
			//str := helper.IntArrayJoinString(a, "")
			//if !helper.IsStringsArrayContain(*cacheArr, str) {
				*res = append(*res, a)
				//*cacheArr = append(*cacheArr, str)
			//}
		}
		return
	}

	for index, value := range arr {
		if index > 0  && arr[index] == arr[index -1]{
			continue
		}
		if !boolArr[index] {
			boolArr[index] = true
			tempArr = append(tempArr, value)
			dfsCombinationSum2(arr, res, target-value, tempArr, cacheArr, boolArr)
			boolArr[index] = false
			tempArr = tempArr[:len(tempArr)-1]
		}
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
