package leetcode

import "fmt"

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	total, sum := 0, 0
	for _, value := range nums {
		sum += value
	}
	for index, value := range nums {
		total += value
		fmt.Println(total, sum, index)
		if total == sum {
			return index
		}
		sum -= value
	}
	return -1
}
