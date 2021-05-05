package leetcode

import (
	"leet_code_go/helper"
	"sort"
)

// https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	cacheArray := []string{}
	DfsCountTarget(candidates, target, &res, []int{}, &cacheArray)
	return res
}

func DfsCountTarget(arr []int, target int, res *[][]int, tempArr []int, cacheArray *[]string) {
	if target <= 0 {
		if target == 0 {
			a := make([]int, len(tempArr))
			copy(a, tempArr)
			sort.Ints(a)
			str := helper.IntArrayJoinString(a, "")
			if !helper.IsStringsArrayContain(*cacheArray, str) {
				*res = append(*res, a)
				*cacheArray = append(*cacheArray, str)
			}
		}
		return
	}

	for _, value := range arr {
		tempArr = append(tempArr, value)
		DfsCountTarget(arr, target-value, res, tempArr, cacheArray)
		tempArr = tempArr[:len(tempArr)-1]
	}
}
